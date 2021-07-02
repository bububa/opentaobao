package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// TaobaoXhotelOrderOfflineSettleCancel 线下信用住取消结账专用接口
// taobao.xhotel.order.offline.settle.cancel
//
// 线下信用住取消结账专用接口
func TaobaoXhotelOrderOfflineSettleCancel(clt *core.SDKClient, req *xhotel.TaobaoXhotelOrderOfflineSettleCancelAPIRequest, session string) (*xhotel.TaobaoXhotelOrderOfflineSettleCancelAPIResponse, error) {
	var resp xhotel.TaobaoXhotelOrderOfflineSettleCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
