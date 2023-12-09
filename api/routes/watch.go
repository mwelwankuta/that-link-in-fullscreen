package routes

import (
	"encoding/json"
	"io"
	"net/http"

	response "github.com/mwelwankuta/that-link-in-fullscreen/api/response-handlers"
	"github.com/mwelwankuta/that-link-in-fullscreen/utils"
)

type OpenVideoLinkInBrowserBody struct {
	Link string `json:"link"`
}

func OpenVideoLinkInBrowser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		response.MessageResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var parsedBody OpenVideoLinkInBrowserBody
	if err := json.Unmarshal(body, &parsedBody); err != nil {
		response.MessageResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !utils.IsYouTubeLink(parsedBody.Link) {
		response.MessageResponse(w, "link passed is not a valid youtube link", http.StatusBadRequest)
		return
	}

	if err := utils.PlayVideo(parsedBody.Link); err != nil {
		response.MessageResponse(w, err.Error(), http.StatusInternalServerError)
	}

	response.MessageResponse(w, "playing...", http.StatusOK)
}
