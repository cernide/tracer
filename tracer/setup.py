import os
import sys

from setuptools import find_packages, setup
from setuptools.command.test import test as TestCommand


def read_readme():
    if not os.path.exists("./README.md"):
        return ""
    with open("./README.md") as f:
        return f.read()


class PyTest(TestCommand):
    def finalize_options(self):
        TestCommand.finalize_options(self)
        self.test_args = []
        self.test_suite = True

    def run_tests(self):
        import pytest

        errcode = pytest.main(self.test_args)
        sys.exit(errcode)


with open(os.path.join("./tracer/pkg.py"), encoding="utf8") as f:
    pkg = {}
    exec(f.read(), pkg)


with open("requirements/dev.txt") as requirements_file:
    dev_requirements = requirements_file.read().splitlines()

extra = {
    "polyaxon": ["polyaxon"],
    "dev": dev_requirements,
    "all": [
        "scikit-learn",
        "Pillow",
        "matplotlib",
        "moviepy",
        "plotly",
        "bokeh",
        "pandas",
        "altair",
    ],
}

setup(
    name=pkg["NAME"],
    version=pkg["VERSION"],
    description=pkg["DESC"],
    long_description=read_readme(),
    long_description_content_type="text/markdown",
    maintainer=pkg["AUTHOR"],
    maintainer_email=pkg["EMAIL"],
    author=pkg["AUTHOR"],
    author_email=pkg["EMAIL"],
    url=pkg["URL"],
    license=pkg["LICENSE"],
    platforms="any",
    packages=find_packages(
        exclude=["*.tests", "*.tests.*", "tests.*", "tests"]),
    keywords=[
        "polyaxon",
        "aws",
        "s3",
        "microsoft",
        "azure",
        "google cloud storage",
        "gcs",
        "deep-learning",
        "machine-learning",
        "data-science",
        "neural-networks",
        "artificial-intelligence",
        "ai",
        "reinforcement-learning",
        "kubernetes",
        "aws",
        "microsoft",
        "azure",
        "google cloud",
        "tensorFlow",
        "pytorch",
        "matplotlib",
        "plotly",
        "visualization",
        "analytics",
    ],
    install_requires=[],
    extras_require=extra,
    python_requires=">=3.8",
    classifiers=[
        "Intended Audience :: Information Technology",
        "Intended Audience :: System Administrators",
        "Intended Audience :: Developers",
        "Intended Audience :: Science/Research",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python",
        "Topic :: Software Development :: Libraries :: Application Frameworks",
        "Topic :: Software Development :: Libraries :: Python Modules",
        "Topic :: Software Development :: Libraries",
        "Topic :: Software Development",
        "Topic :: Scientific/Engineering :: Artificial Intelligence",
        "Typing :: Typed",
        "Programming Language :: Python :: 3 :: Only",
        "Programming Language :: Python :: 3.6",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
    ],
    tests_require=["pytest"],
    cmdclass={"test": PyTest},
)
