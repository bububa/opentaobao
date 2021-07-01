package aliexpress

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpress"
)

/* AliexpressFeedPostPublish
同步帖子
aliexpress.feed.post.publish

站外平台同步帖子至AE FEED域 */
func AliexpressFeedPostPublish(clt *core.SDKClient, req *aliexpress.AliexpressFeedPostPublishAPIRequest, session string) (*aliexpress.AliexpressFeedPostPublishAPIResponse, error) {
	var resp aliexpress.AliexpressFeedPostPublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
