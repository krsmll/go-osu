package client

import (
	"osu-client/client/opts"
	"osu-client/enum"
	"osu-client/model"
)

// GetNewsListing returns a list of news posts and related metadata.
func (c *OsuClient) GetNewsListing(opts opts.GetNewsListingOpts) (*model.NewsListing, error) {
	return handleResponse[model.NewsListing](c.getReq(enum.EndpointGetNews, opts))
}

// GetNewsPost returns details of the specified news post.
func (c *OsuClient) GetNewsPost(opts opts.GetNewsPostOpts) (*model.NewsPost, error) {
	if opts.PostID != nil && opts.PostSlug == nil && opts.Key == nil {
		idKey := enum.NewsPostKeyID
		opts.Key = &idKey
	}

	return handleResponse[model.NewsPost](c.getReq(enum.EndpointGetNewsPost, opts))
}
