# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

from polyaxon_deploy.operators.cmd_operator import CmdOperator
from polyaxon_deploy.operators.exceptions import OperatorException


class DockerOperator(CmdOperator):
    CMD = 'docker'

    @classmethod
    def params(cls, args):
        params = [cls.CMD] + args
        return params

    @classmethod
    def execute(cls, args):
        params = cls.params(args)
        return cls._execute(params=params, env=None)

    @classmethod
    def set_volume(cls, volume):
        args = ['volume', 'create', '--name={}'.format(volume)]
        if not volume:
            raise OperatorException('docker',
                                    args,
                                    None,
                                    None,
                                    'Volume name was not provided')
        return cls.execute(args)
