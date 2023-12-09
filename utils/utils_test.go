package utils

import "testing"

var links = []map[string]interface{}{
	{
		"link":     "https://www.youtube.com/watch?v=example",
		"can_pass": true,
	},
	{
		"link":     "https://www.youtue.com/video?video=example",
		"can_pass": false,
	},
}

func TestIsYouTubeLink(t *testing.T) {

	for i := 0; i < len(links); i++ {
		link, exists := links[i]["link"]
		if exists != true {
			t.Error("no links")
			return
		}

		canPass, exists := links[i]["can_pass"]
		if exists != true {
			t.Error("no pass status")
			return
		}

		if IsYouTubeLink(link.(string)) != canPass.(bool) {
			t.Fatal("Link passed as not expected")
		}
	}
}
