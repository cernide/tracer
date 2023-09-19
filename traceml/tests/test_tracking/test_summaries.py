import os
import pytest
import uuid

from mock import patch

from polyaxon import settings
from polyaxon.env_vars.keys import (
    ENV_KEYS_COLLECT_ARTIFACTS,
    ENV_KEYS_COLLECT_RESOURCES,
)
from polyaxon.utils.test_utils import BaseTestCase
from traceml.artifacts import V1RunArtifact
from traceml.events import V1Events
from traceml.tracking.run import Run


@pytest.mark.tracking_mark
class TestEventsSummaries(BaseTestCase):
    def setUp(self):
        super().setUp()
        settings.CLIENT_CONFIG.is_managed = False
        settings.CLIENT_CONFIG.is_offline = True
        os.environ[ENV_KEYS_COLLECT_ARTIFACTS] = "false"
        os.environ[ENV_KEYS_COLLECT_RESOURCES] = "false"
        uid = uuid.uuid4().hex
        with patch("traceml.tracking.run.Run._set_exit_handler") as exit_mock:
            self.run = Run(project="test.test", run_uuid=uid)
        assert exit_mock.call_count == 1

    def test_metrics_summaries(self):
        summaries, last_values = self.run._collect_events_summaries(
            events_path="tests/fixtures/events",
            events_kind="metric",
            last_check=None,
        )
        abspath = os.path.abspath("tests/fixtures/events/metric/metric_events.plx")
        events = V1Events.read(
            name="metric_events",
            kind="metric",
            data=abspath,
        )
        assert events.name == "metric_events"
        assert summaries == [
            V1RunArtifact(
                name="metric_events",
                kind="metric",
                connection=None,
                summary=events.get_summary(),
                path=abspath,
                is_input=False,
            )
        ]
        assert last_values == {"metric_events": 0.3}

    def test_images_summaries(self):
        summaries, last_values = self.run._collect_events_summaries(
            events_path="tests/fixtures/events",
            events_kind="image",
            last_check=None,
        )
        abspath = os.path.abspath("tests/fixtures/events/image/image_events.plx")
        events = V1Events.read(
            name="image_events",
            kind="image",
            data=abspath,
        )
        assert events.name == "image_events"
        assert summaries == [
            V1RunArtifact(
                name="image_events",
                kind="image",
                connection=None,
                summary=events.get_summary(),
                path=abspath,
                is_input=False,
            )
        ]
        assert last_values == {}

    def test_histograms_summaries(self):
        summaries, last_values = self.run._collect_events_summaries(
            events_path="tests/fixtures/events",
            events_kind="histogram",
            last_check=None,
        )
        abspath = os.path.abspath(
            "tests/fixtures/events/histogram/histogram_events.plx"
        )
        events = V1Events.read(
            name="histogram_events",
            kind="histogram",
            data=abspath,
        )
        assert events.name == "histogram_events"
        assert summaries == [
            V1RunArtifact(
                name="histogram_events",
                kind="histogram",
                connection=None,
                summary=events.get_summary(),
                path=abspath,
                is_input=False,
            )
        ]
        assert last_values == {}

    def test_videos_summaries(self):
        summaries, last_values = self.run._collect_events_summaries(
            events_path="tests/fixtures/events",
            events_kind="video",
            last_check=None,
        )
        abspath = os.path.abspath("tests/fixtures/events/video/video_events.plx")
        events = V1Events.read(
            name="video_events",
            kind="video",
            data=abspath,
        )
        assert events.name == "video_events"
        assert summaries == [
            V1RunArtifact(
                name="video_events",
                kind="video",
                connection=None,
                summary=events.get_summary(),
                path=abspath,
                is_input=False,
            )
        ]
        assert last_values == {}

    def test_audios_summaries(self):
        summaries, last_values = self.run._collect_events_summaries(
            events_path="tests/fixtures/events",
            events_kind="audio",
            last_check=None,
        )
        abspath = os.path.abspath("tests/fixtures/events/audio/audio_events.plx")
        events = V1Events.read(
            name="audio_events",
            kind="audio",
            data=abspath,
        )
        assert events.name == "audio_events"
        assert summaries == [
            V1RunArtifact(
                name="audio_events",
                kind="audio",
                connection=None,
                summary=events.get_summary(),
                path=abspath,
                is_input=False,
            )
        ]
        assert last_values == {}

    def test_htmls_summaries(self):
        summaries, last_values = self.run._collect_events_summaries(
            events_path="tests/fixtures/events",
            events_kind="html",
            last_check=None,
        )
        abspath = os.path.abspath("tests/fixtures/events/html/html_events.plx")
        events = V1Events.read(
            name="html_events",
            kind="html",
            data=abspath,
        )
        assert events.name == "html_events"
        assert summaries == [
            V1RunArtifact(
                name="html_events",
                kind="html",
                connection=None,
                summary=events.get_summary(),
                path=abspath,
                is_input=False,
            )
        ]
        assert last_values == {}

    def test_charts_summaries(self):
        summaries, last_values = self.run._collect_events_summaries(
            events_path="tests/fixtures/events",
            events_kind="chart",
            last_check=None,
        )
        abspath = os.path.abspath("tests/fixtures/events/chart/chart_events.plx")
        events = V1Events.read(
            name="chart_events",
            kind="chart",
            data=abspath,
        )
        assert events.name == "chart_events"
        assert summaries == [
            V1RunArtifact(
                name="chart_events",
                kind="chart",
                connection=None,
                summary=events.get_summary(),
                path=abspath,
                is_input=False,
            )
        ]
        assert last_values == {}

    def test_curves_summaries(self):
        summaries, last_values = self.run._collect_events_summaries(
            events_path="tests/fixtures/events",
            events_kind="curve",
            last_check=None,
        )
        abspath = os.path.abspath("tests/fixtures/events/curve/curve_events.plx")
        events = V1Events.read(
            name="curve_events",
            kind="curve",
            data=abspath,
        )
        assert events.name == "curve_events"
        assert summaries == [
            V1RunArtifact(
                name="curve_events",
                kind="curve",
                connection=None,
                summary=events.get_summary(),
                path=abspath,
                is_input=False,
            )
        ]
        assert last_values == {}

    def test_artifacts_summaries(self):
        summaries, last_values = self.run._collect_events_summaries(
            events_path="tests/fixtures/events",
            events_kind="artifact",
            last_check=None,
        )
        abspath = os.path.abspath("tests/fixtures/events/artifact/artifact_events.plx")
        events = V1Events.read(
            name="artifact_events",
            kind="artifact",
            data=abspath,
        )
        assert events.name == "artifact_events"
        assert summaries == [
            V1RunArtifact(
                name="artifact_events",
                kind="artifact",
                connection=None,
                summary=events.get_summary(),
                path=abspath,
                is_input=False,
            )
        ]
        assert last_values == {}

    def test_models_summaries(self):
        summaries, last_values = self.run._collect_events_summaries(
            events_path="tests/fixtures/events",
            events_kind="model",
            last_check=None,
        )
        summaries = {s.name: s for s in summaries}
        abspath = os.path.abspath("tests/fixtures/events/model/model_events.plx")
        events = V1Events.read(
            name="model_events",
            kind="model",
            data=abspath,
        )
        assert events.name == "model_events"
        assert summaries["model_events"] == V1RunArtifact(
            name="model_events",
            kind="model",
            connection=None,
            summary=events.get_summary(),
            path=abspath,
            is_input=False,
        )

        abspath = os.path.abspath(
            "tests/fixtures/events/model/model_events_without_step.plx"
        )
        events_without_step = V1Events.read(
            name="model_events_without_step",
            kind="model",
            data=abspath,
        )
        assert events_without_step.name == "model_events_without_step"
        assert summaries["model_events_without_step"] == V1RunArtifact(
            name="model_events_without_step",
            kind="model",
            connection=None,
            summary=events_without_step.get_summary(),
            path=abspath,
            is_input=False,
        )
        assert last_values == {}

    def test_dataframes_summaries(self):
        summaries, last_values = self.run._collect_events_summaries(
            events_path="tests/fixtures/events",
            events_kind="dataframe",
            last_check=None,
        )
        abspath = os.path.abspath(
            "tests/fixtures/events/dataframe/dataframe_events.plx"
        )
        events = V1Events.read(
            name="dataframe_events",
            kind="dataframe",
            data=abspath,
        )
        assert events.name == "dataframe_events"
        assert summaries == [
            V1RunArtifact(
                name="dataframe_events",
                kind="dataframe",
                connection=None,
                summary=events.get_summary(),
                path=abspath,
                is_input=False,
            )
        ]
        assert last_values == {}
