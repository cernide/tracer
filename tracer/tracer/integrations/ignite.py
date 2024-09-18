from tracer.exceptions import TracemlException
from tracer.run import Run

try:
    from ignite.contrib.handlers.polyaxon_logger import PolyaxonLogger
except ImportError:
    raise TracemlException("ignite is required to use the tracking Logger")


class Logger(PolyaxonLogger):
    def __init__(self, *args, **kwargs):
        self.experiment = kwargs.get("run")
        if not self.experiment:
            self.experiment = Run(*args, **kwargs)
