package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoStreetestSessionGet 根据获取压测用户的sessionKey
// taobao.streetest.session.get
//
// 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
func TaobaoStreetestSessionGet(ctx context.Context, clt *core.SDKClient, req *util.TaobaoStreetestSessionGetAPIRequest, resp *util.TaobaoStreetestSessionGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
