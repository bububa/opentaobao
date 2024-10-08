package caipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/caipiao"
)

// TaobaoCaipiaoSignstatusCheck 检查用户是否签署支付宝代购协议
// taobao.caipiao.signstatus.check
//
// 检查用户是否签署了支付宝代扣协议。如果签署了，返回true; 如果没签署，返回false, 同时返回签署代扣协议的Url。
func TaobaoCaipiaoSignstatusCheck(ctx context.Context, clt *core.SDKClient, req *caipiao.TaobaoCaipiaoSignstatusCheckAPIRequest, resp *caipiao.TaobaoCaipiaoSignstatusCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
