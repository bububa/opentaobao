package aliexpress

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpress"
)

// AliexpressFeedPostPublish 同步帖子
// aliexpress.feed.post.publish
//
// 站外平台同步帖子至AE FEED域
func AliexpressFeedPostPublish(ctx context.Context, clt *core.SDKClient, req *aliexpress.AliexpressFeedPostPublishAPIRequest, resp *aliexpress.AliexpressFeedPostPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
