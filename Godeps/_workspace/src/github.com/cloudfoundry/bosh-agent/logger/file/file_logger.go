package file

import (
	"os"

	bosherr "github.com/cloudfoundry/bosh-agent/errors"
	boshlog "github.com/cloudfoundry/bosh-agent/logger"
	boshsys "github.com/cloudfoundry/bosh-agent/system"
)

const DefaultLogFileMode = os.FileMode(0666)

// New returns a new logger (that writes to the specified file) & the open file handle
// All log levels >= the specified level are written to the specified file.
// User is responsible for closing the returned file handle, unless an error is returned.
func New(level boshlog.LogLevel, filePath string, fileMode os.FileMode, fs boshsys.FileSystem) (boshlog.Logger, boshsys.File, error) {
	file, err := fs.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, fileMode)
	if err != nil {
		return boshlog.Logger{}, file, bosherr.WrapErrorf(err, "Failed to open log file '%s'", filePath)
	}

	return boshlog.NewWriterLogger(level, file, file), file, nil
}
