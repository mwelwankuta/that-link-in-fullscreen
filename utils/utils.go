package utils

import (
	"os/exec"
	"regexp"
	"runtime"
)

func IsYouTubeLink(link string) bool {
	youtubeRegex := `^(https?://)?(www\.)?(youtube\.com|youtu\.be)/.*$`

	re := regexp.MustCompile(youtubeRegex)

	return re.MatchString(link)
}

func OpenBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
