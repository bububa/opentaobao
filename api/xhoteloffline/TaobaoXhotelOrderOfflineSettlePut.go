package xhoteloffline

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhoteloffline"
)

// TaobaoXhotelOrderOfflineSettlePut 线下信用住结账专用接口
// taobao.xhotel.order.offline.settle.put
//
// 线下信用住结账专用接口
func TaobaoXhotelOrderOfflineSettlePut(ctx context.Context, clt *core.SDKClient, req *xhoteloffline.TaobaoXhotelOrderOfflineSettlePutAPIRequest, resp *xhoteloffline.TaobaoXhotelOrderOfflineSettlePutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
