package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanPayresultQuery 百川支付完成回调
// taobao.baichuan.payresult.query
//
// 百川支付完成回调
func TaobaoBaichuanPayresultQuery(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanPayresultQueryAPIRequest, resp *baichuan.TaobaoBaichuanPayresultQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
