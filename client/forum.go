package client

import (
	"osu-client/client/opts"
	"osu-client/enum"
	"osu-client/model"
)

// GetTopicAndPosts returns a topic and its posts.
func (c *OsuClient) GetTopicAndPosts(opts opts.GetTopicAndPostsOpts) (*model.ForumTopicAndPosts, error) {
	return handleResponse[model.ForumTopicAndPosts](c.getReq(enum.EndpointGetTopicAndPosts, opts))
}
