package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoWeitaoFeedIsrelation 是否关注
// taobao.weitao.feed.isrelation
//
// 判断用户是否关注对应的公共账号
func TaobaoWeitaoFeedIsrelation(ctx context.Context, clt *core.SDKClient, req *mtopopen.TaobaoWeitaoFeedIsrelationAPIRequest, resp *mtopopen.TaobaoWeitaoFeedIsrelationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
