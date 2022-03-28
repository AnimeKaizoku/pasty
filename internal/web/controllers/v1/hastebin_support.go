package v1

import (
	"encoding/json"
	"time"

	"github.com/AnimeKaizoku/pasty/internal/config"
	"github.com/AnimeKaizoku/pasty/internal/shared"
	"github.com/AnimeKaizoku/pasty/internal/storage"
	"github.com/AnimeKaizoku/pasty/internal/utils"
	"github.com/valyala/fasthttp"
)

// HastebinSupportHandler handles the legacy hastebin requests
func HastebinSupportHandler(ctx *fasthttp.RequestCtx) {
	// Check content length before reading body into memory
	if config.Current.LengthCap > 0 &&
		ctx.Request.Header.ContentLength() > config.Current.LengthCap {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("request body length overflow")
		return
	}

	// Define the paste content
	var content string
	if string(ctx.Request.Header.ContentType()) == "multipart/form-data" {
		content = string(ctx.FormValue("data"))
	} else {
		content = string(ctx.PostBody())
	}

	// Acquire the paste ID
	id, err := storage.AcquireID()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(err.Error())
		return
	}

	// Create the paste object
	paste := &shared.Paste{
		ID:      id,
		Content: content,
		Created: time.Now().Unix(),
	}

	// Set a modification token
	if config.Current.ModificationTokens {
		paste.ModificationToken = utils.RandomString(config.Current.ModificationTokenCharacters, config.Current.ModificationTokenLength)

		err = paste.HashModificationToken()
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetBodyString(err.Error())
			return
		}
	}

	// Save the paste
	err = storage.Current.Save(paste)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(err.Error())
		return
	}

	// Respond with the paste key
	jsonData, _ := json.Marshal(map[string]string{
		"key": paste.ID,
	})
	ctx.SetBody(jsonData)
}
