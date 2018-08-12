package camera

import (
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

const (
	still     = "raspistill"
	tlapse    = "-tl"
	timeout   = "-t"
	output    = "-o"
	filetype  = ".jpg"
	timestamp = "2006-01-02_15:04:05"

	defaultTimeout = 1
	noTimelapse    = -1
)

// Camera with params
type Camera struct {
	params params
}

type params struct {
	timeout   int
	timelapse int
	path      string
	filename  string
}

// New Camera with path
func New(path string) *Camera {
	return &Camera{params{defaultTimeout, noTimelapse, path, getDefaultFilename()}}
}

// NewTimelapsed Camera with path and timelapse interval
func NewTimelapsed(path string, timelapse int) *Camera {
	return &Camera{params{defaultTimeout, timelapse, path, getTimeLapseFilename("data")}}
}

func makeArgs(p params) []string {
	args := make([]string, 0)
	args = append(args, timeout)
	args = append(args, strconv.Itoa(p.timeout))
	if p.timelapse != noTimelapse {
		args = append(args, tlapse)
		args = append(args, strconv.Itoa(p.timelapse))
	}
	args = append(args, output)
	args = append(args, filepath.Join(p.path, p.filename))
	return args
}

func getDefaultFilename() string {
	return time.Now().Format(timestamp) + filetype
}

func getTimeLapseFilename(prefix string) string {
	return prefix + "_%04d" + filetype
}

// Capture an image or a timelapse
func (c *Camera) Capture() (string, error) {
	args := makeArgs(c.params)
	fullPath := filepath.Join(c.params.path, c.params.filename)

	cmd := exec.Command(still, args...)
	_, err := cmd.StdoutPipe()
	if err != nil {
		return fullPath, err
	}
	err = cmd.Start()
	if err != nil {
		return fullPath, err
	}
	cmd.Wait()
	return fullPath, nil
}

