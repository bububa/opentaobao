package uscesl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizApActivate 激活AP价签通讯模块
// taobao.uscesl.biz.ap.activate
//
// 激活AP价签通讯模块
func TaobaoUsceslBizApActivate(ctx context.Context, clt *core.SDKClient, req *uscesl.TaobaoUsceslBizApActivateAPIRequest, resp *uscesl.TaobaoUsceslBizApActivateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
