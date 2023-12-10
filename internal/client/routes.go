package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	response "github.com/mwelwankuta/that-link-in-fullscreen/api/response"
	utils "github.com/mwelwankuta/that-link-in-fullscreen/utils"
)

func HandlePlayRoute(w http.ResponseWriter, r *http.Request) {
	videoLink := mux.Vars(r)["video_link"]
	fmt.Println(videoLink)
	if videoLink == "" {
		response.MessageResponse(w, "a video link was not passed", http.StatusBadRequest)
		return
	}

	if !utils.IsYouTubeLink(videoLink) {
		response.MessageResponse(w, "invalid youtube link passed", http.StatusBadRequest)
		return
	}

	videoId, err := ExtractVideoId(videoLink)

	if err != nil {
		response.MessageResponse(w, "invalid youtube link passed", http.StatusBadRequest)
		return
	}

	app := player(videoId)
	app.Render(context.Background(), w)
}

func HandleHealthRoute(w http.ResponseWriter, r *http.Request) {
	response.MessageResponse(w, "healthy", http.StatusOK)
}
