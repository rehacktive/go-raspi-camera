package camera

import (
	"os/exec"
	"path/filepath"
	"time"
	"strconv"
)

const (
	STILL      = "raspistill"
	TLAPSE     = "-tl"
	TIMEOUT    = "-t"
	OUTFLAG    = "-o"
	FILE_TYPE  = ".jpg"
	TIME_STAMP = "2006-01-02_15:04:05"

	DEFAULT_TIMEOUT   = 1
	NO_TIMELAPSE = -1
)

type Camera struct {
	params		Params	
}

type Params struct {
	timeout		int
	timelapse	int
	path		string
	filename	string
}

func New(path string) *Camera {
	return &Camera{Params{DEFAULT_TIMEOUT,NO_TIMELAPSE, path, getDefaultFilename()}}
}

func NewTimelapsed(path string, timelapse int) *Camera {
	return &Camera{Params{DEFAULT_TIMEOUT, timelapse, path, getTimeLapseFilename("data")}}
}

func makeArgs(p Params) ([]string) {
	args := make([]string, 0)
	args = append(args, TIMEOUT)
	args = append(args, strconv.Itoa(p.timeout))
	if(p.timelapse != NO_TIMELAPSE) {
		args = append(args, TLAPSE)
		args = append(args, strconv.Itoa(p.timelapse))
	}
	args = append(args, OUTFLAG)
	args = append(args, filepath.Join(p.path, p.filename))		
	return args
}

func getDefaultFilename() (string) {
	return time.Now().Format(TIME_STAMP) + FILE_TYPE
}

func getTimeLapseFilename(prefix string) (string) {
	return prefix + "_%04d" + FILE_TYPE
}

func (c *Camera) Capture() (string, error) {
	args := makeArgs(c.params)	
        fullPath := filepath.Join(c.params.path,c.params.filename)

	cmd := exec.Command(STILL, args...)
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

