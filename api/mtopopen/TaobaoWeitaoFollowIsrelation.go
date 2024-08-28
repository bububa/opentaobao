package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoWeitaoFollowIsrelation 微淘是否关注
// taobao.weitao.follow.isrelation
//
// 判断用户是否关注对应的公共账号
func TaobaoWeitaoFollowIsrelation(ctx context.Context, clt *core.SDKClient, req *mtopopen.TaobaoWeitaoFollowIsrelationAPIRequest, resp *mtopopen.TaobaoWeitaoFollowIsrelationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
