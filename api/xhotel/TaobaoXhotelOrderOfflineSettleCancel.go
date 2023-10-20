package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// TaobaoXhotelOrderOfflineSettleCancel 线下信用住取消结账专用接口
// taobao.xhotel.order.offline.settle.cancel
//
// 线下信用住取消结账专用接口
func TaobaoXhotelOrderOfflineSettleCancel(clt *core.SDKClient, req *xhotel.TaobaoXhotelOrderOfflineSettleCancelAPIRequest, resp *xhotel.TaobaoXhotelOrderOfflineSettleCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
