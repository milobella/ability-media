package plex

import (
	"github.com/jrudio/go-plex-client"
)

func Search(conf Config, query string) (url string, title string, err error) {
	//TODO: instantiate the plex connection asynchronously
	var plexConn *plex.Plex
	if plexConn, err = plex.New(conf.Url, conf.Token); err != nil {
		return
	}

	var results plex.SearchResults
	if results, err = plexConn.Search(query); err != nil {
		return
	}
	for _, metadata := range results.MediaContainer.Metadata {
		if len(metadata.GUID) > 0 {
			url = metadata.GUID

			if len(metadata.TitleSort) > 0 {
				// titleSort field seems to be shorter and thus more suitable for vocal answer.
				// Is it a typo from plex ?
				title = metadata.TitleSort
			} else {
				title = metadata.Title
			}

			return
		}
	}
	return "", "", nil
}
