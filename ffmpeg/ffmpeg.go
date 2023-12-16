package ffmpeg

import (
	"fmt"
	"strings"
)

func GetDuration(ffmpegOutput string) (string, error) {
	durationIndex := strings.Index(string(ffmpegOutput), "Duration: ")
	if durationIndex == -1 {
		return "", fmt.Errorf("No duration information on the ffmpeg video output")
	}

	durationStr := string(ffmpegOutput)[durationIndex+len("Duration: "):]
	duration := strings.Split(durationStr, ",")[0]

	return duration, nil
}
