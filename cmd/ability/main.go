package main

import (
	"github.com/milobella/ability-media/pkg/plex"
	"github.com/milobella/ability-sdk-go/pkg/config"
	"github.com/milobella/ability-sdk-go/pkg/model"
	"github.com/milobella/ability-sdk-go/pkg/server"
	"github.com/milobella/ability-sdk-go/pkg/server/conditions"
	"github.com/milobella/ability-sdk-go/pkg/server/interpreters"
)

const titleEntity = "title"
const playMediaAction = "play_media"

// ConfigExtension of the SDK configuration
type ConfigExtension struct {
	Plex plex.Config `mapstructure:"plex"`
}

func main() {
	// Read configuration
	var confExt ConfigExtension
	conf := config.Read(&confExt)
	// Initialize server
	srv := server.New("Media", conf.Server.Port)

	// Register first the conditions on actions because they have priority on intents.
	// The condition returns true if an action is pending.
	srv.Register(conditions.IfInSlotFilling(playMediaAction), handlePlayMedia(confExt.Plex))

	// Then we register intents routing rules.
	// It means that if no pending action has been found in the context, we'll use intent to decide the handler.
	srv.Register(conditions.IfIntents("PLAY_MOVIE", "PLAY_SERIES"), handlePlayMedia(confExt.Plex))

	srv.Serve()
}

func handlePlayMedia(conf plex.Config) func(*model.Request, *model.Response) {
	return func(request *model.Request, response *model.Response) {
		var stopper func(*model.Response)
		var instrument *string
		var title *string

		instrument, stopper = interpreters.
			FromInstrument(model.InstrumentKindChromeCast, playMediaAction).
			Interpret(request)
		if stopper != nil {
			stopper(response)
			return
		}

		title, stopper = interpreters.
			FromNLU(titleEntity, playMediaAction).
			OverridingNotFoundMsg(model.NLG{Sentence: "Which title do you want to play ?"}).
			InterpretFirst(request)
		if stopper != nil {
			stopper(response)
			return
		}

		url, titleFound, err := plex.Search(conf, *title)
		if err != nil {
			response.Nlg.Sentence = "An error occurred while trying to search a media."
			return
		}

		if len(url) == 0 {
			response.Nlg.Sentence = "Didn't find any media corresponding to your search."
			return
		}

		if len(url) > 0 {
			response.Nlg = model.NLG{
				Sentence: "Playing {{ title }} on the chrome cast {{ instrument }}.",
				Params: []model.NLGParam{{
					Name:  "title",
					Value: titleFound,
					Type:  "string",
				}, {
					Name:  "instrument",
					Value: instrument,
					Type:  "string",
				}},
			}
			response.Actions = []model.Action{{
				Identifier: playMediaAction,
				Params: []model.ActionParameter{{
					Key:   "instrument",
					Value: *instrument,
				}},
			}}
			return
		}
	}
}
