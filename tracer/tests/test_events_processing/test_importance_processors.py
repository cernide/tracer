import numpy as np
import pytest

from unittest import TestCase

from tracer.processors.importance_processors import calculate_importance_correlation


@pytest.mark.processors_mark
class TestFeatureImportance(TestCase):
    def test_empty_value(self):
        assert calculate_importance_correlation(None, None) is None

    def test_single_value(self):
        res = calculate_importance_correlation([{"param1": 3}], [4])
        exp = {"param1": {"correlation": None, "importance": 0.0}}
        assert res == exp

    def test_correct_values(self):
        res = calculate_importance_correlation(
            [{"param1": 3}, {"param1": 4}, {"param1": 5}],
            [3, 4, 5],
        )
        exp = {"param1": {"correlation": 1.0, "importance": 1.0}}
        assert res == exp

    def test_multiple_params(self):
        res = calculate_importance_correlation(
            [
                {
                    "param1": 1,
                    "param2": 3,
                },
                {
                    "param1": 2,
                    "param2": 2,
                },
                {
                    "param1": 3,
                    "param2": 1,
                },
            ],
            [1, 2, 3],
        )
        exp = {
            "param1": {"correlation": 1.0, "importance": 0.464},
            "param2": {"correlation": -1.0, "importance": 0.536},
        }
        assert res == exp

    def test_wrong_string_params(self):
        assert calculate_importance_correlation(["foo", "bar"], []) is None

    def test_complex_params(self):
        res = calculate_importance_correlation(
            [{"param1": "str1", "param2": 1}, {"param1": 2, "param2": 2}], [1, 2]
        )
        exp = {
            "param1_2": {"correlation": 1.0, "importance": 0.348},
            "param1_str1": {"correlation": -1.0, "importance": 0.308},
            "param2": {"correlation": 1.0, "importance": 0.344},
        }
        assert res == exp

    def test_nan_value(self):
        assert (
            calculate_importance_correlation(
                [{"param1": 3, "param2": 1}, {"param1": 2, "param2": 2}], [np.nan, 2]
            )
            is None
        )

    def test_empty_metrics(self):
        assert calculate_importance_correlation(
            [{"foo": 2, "bar": 4}], []) is None
