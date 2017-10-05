# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

import itertools
import six

from polyaxon_schemas.polyaxonfile import validator
from polyaxon_schemas.polyaxonfile import reader
from polyaxon_schemas.polyaxonfile.parser import Parser
from polyaxon_schemas.polyaxonfile.specification import Specification
from polyaxon_schemas.polyaxonfile.utils import cached_property
from polyaxon_schemas.settings import ClusterConfig, RunTypes


class PolyaxonFile(object):
    """Parses Polyaxonfiles, and validate that it respects the current file specification"""

    def __init__(self, filepath):
        self._filepath = filepath

        self._data = reader.read(self._filepath)
        Parser.check_data(data=self._data)
        headers = Parser.get_headers(self._data)
        matrix = Parser.get_matrix(self._data)
        self._matrix = validator.validate_matrix(matrix)
        self._headers = validator.validate_headers(headers)
        self._parsed_data = []
        self._validated_data = []

        matrix_declarations = self.matrix_declarations if self.matrix_declarations else [{}]
        for matrix_declaration in matrix_declarations:
            parsed_data = Parser.parse(self._data, matrix_declaration)
            self._parsed_data.append(parsed_data)
            self._validated_data.append(validator.validate(parsed_data))

    @classmethod
    def read(cls, filepath):
        if isinstance(filepath, cls):
            return filepath
        return cls(filepath)

    @cached_property
    def data(self):
        return self._data

    @cached_property
    def matrix(self):
        return self._matrix

    @cached_property
    def matrix_space(self):
        if not self.matrix:
            return 1

        space_size = 1
        for value in six.itervalues(self.matrix):
            space_size *= len(value.to_numpy())
        return space_size

    @cached_property
    def experiments_def(self):
        concurrent_experiments = self.settings.concurrent_experiments if self.settings else 1
        return self.matrix_space, concurrent_experiments

    @cached_property
    def matrix_declarations(self):
        if not self.matrix:
            return []

        declarations = []
        keys = list(six.iterkeys(self.matrix))
        values = [v.to_numpy() for v in six.itervalues(self.matrix)]
        for v in itertools.product(*values):
            declarations.append(dict(zip(keys, v)))

        if len(declarations) != self.matrix_space:
            raise PolyaxonFile('The matrix declaration is not valid.')
        return declarations

    @cached_property
    def headers(self):
        return self._headers

    @cached_property
    def parsed_data(self):
        return self._parsed_data

    @cached_property
    def version(self):
        return self.headers[Specification.VERSION]

    @cached_property
    def project(self):
        return self.headers[Specification.PROJECT]

    @cached_property
    def settings(self):
        return self.headers.get(Specification.SETTINGS, None)

    @cached_property
    def run_type(self):
        return self.settings.run_type if self.settings else RunTypes.LOCAL

    @cached_property
    def project_path(self):
        project_path = None
        if self.settings:
            project_path = self.settings.logging.path

        return project_path or '/tmp/plx_logs/' + self.project.name

    @cached_property
    def validated_data(self):
        if self.matrix_space == 1:
            return self.get_validated_data_at(0)
        raise AttributeError("""Current polyaxonfile has multiple experiments ({}),
        please use `get_validated_data_at(experiment)` instead.""".format(self.matrix_space))

    def get_validated_data_at(self, experiment):
        if experiment > self.matrix_space:
            raise ValueError("""Could not find an experiment at index {},
            this file has {} experiments""".format(experiment, self.matrix_space))

        return self._validated_data[experiment]

    @cached_property
    def model(self):
        if self.matrix_space == 1:
            return self.get_model_at(0)
        raise AttributeError("""Current polyaxonfile has multiple experiments ({}),
        please use `get_model_at(experiment)` instead.""".format(self.matrix_space))

    def get_model_at(self, experiment):
        return self.get_validated_data_at(experiment)[Specification.MODEL]

    @cached_property
    def environment(self):
        if self.matrix_space == 1:
            return self.get_environment_at(0)
        raise AttributeError("""Current polyaxonfile has multiple experiments ({}),
        please use `get_environment_at(experiment)` instead.""".format(self.matrix_space))

    def get_environment_at(self, experiment):
        return self.get_validated_data_at(experiment).get(Specification.ENVIRONMENT, None)

    @cached_property
    def train(self):
        if self.matrix_space == 1:
            return self.get_train_at(0)
        raise AttributeError("""Current polyaxonfile has multiple experiments ({}),
                please use `get_train_at(experiment)` instead.""".format(self.matrix_space))

    def get_train_at(self, experiment):
        return self.get_validated_data_at(experiment).get(Specification.TRAIN, None)

    @cached_property
    def eval(self):
        if self.matrix_space == 1:
            return self.get_eval_at(0)
        raise AttributeError("""Current polyaxonfile has multiple experiments ({}),
        please use `get_eval_at(experiment)` instead.""".format(self.matrix_space))

    def get_eval_at(self, experiment):
        return self.get_validated_data_at(experiment).get(Specification.EVAL, None)

    @cached_property
    def cluster_def(self):
        if self.matrix_space == 1:
            return self.get_cluster_def_at(0)
        raise AttributeError("""Current polyaxonfile has multiple experiments ({}),
        please use `get_train_at(experiment)` instead.""".format(self.matrix_space))

    def get_cluster_def_at(self, experiment):
        cluster = {
            'master': 1,
        }
        is_distributed = False
        environment = self.get_environment_at(experiment)

        if environment:
            cluster['worker'] = environment.n_workers
            cluster['ps'] = environment.n_ps
            is_distributed = True

        return cluster, is_distributed

    def get_cluster(self,
                    experiment=0,
                    host='127.0.0.1',
                    master_port=10000,
                    worker_port=11000,
                    ps_port=12000):
        def get_address(port):
            return '{}:{}'.format(host, port)

        cluster_def, is_distributed = self.get_cluster_def_at(experiment)

        cluster_config = {
            'master': [get_address(master_port)]
        }

        workers = []
        for i in range(cluster_def.get('worker', 0)):
            workers.append(get_address(worker_port))
            worker_port += 1

        cluster_config['worker'] = workers

        ps = []
        for i in range(cluster_def.get('ps', 0)):
            ps.append(get_address(ps_port))
            ps_port += 1

        cluster_config['ps'] = ps

        return ClusterConfig.from_dict(cluster_config)
