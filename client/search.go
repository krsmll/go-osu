package client

import (
	"osu-client/client/opts"
	"osu-client/enum"
	"osu-client/model"
)

// Search searches users and wiki pages.
func (c *OsuClient) Search(opts opts.SearchOpts) (*model.SearchResults, error) {
	return handleResponse[model.SearchResults](c.getReq(enum.EndpointSearch, opts))
}
