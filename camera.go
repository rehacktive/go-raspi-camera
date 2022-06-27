package camera

import (
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

const (
	still     = "libcamera-still"
	vlip      = "--vflip"
	timeout   = "-t"
	width     = "--width"
	height    = "--height"
	output    = "-o"
	filetype  = ".jpg"
	timestamp = "2006-01-02_15:04:05"

	defaultTimeout = 1
)

type resolution struct {
	width  int
	height int
}

// Camera with params
type Camera struct {
	timeout    int
	resolution resolution
	path       string
}

// New Camera with path
func New(path string, width int, height int) *Camera {
	return &Camera{defaultTimeout, resolution{width, height}, path}
}

func makeArgs(c *Camera) []string {
	args := make([]string, 0)
	args = append(args, timeout)
	args = append(args, strconv.Itoa(c.timeout))
	args = append(args, vlip)
	args = append(args, strconv.Itoa(1))

	args = append(args, width)
	args = append(args, strconv.Itoa(c.resolution.width))
	args = append(args, height)
	args = append(args, strconv.Itoa(c.resolution.height))

	args = append(args, output)
	args = append(args, filepath.Join(c.path, getFilename()))
	return args
}

func getFilename() string {
	return time.Now().Format(timestamp) + filetype
}

// Capture an image or a timelapse
func (c *Camera) Capture() (string, error) {
	args := makeArgs(c)
	fullPath := args[len(args)-1]

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
