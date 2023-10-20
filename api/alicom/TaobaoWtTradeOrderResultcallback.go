package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Taobaowttradeorderresultcallback 商家回调接口
// taobao.wt.trade.order.resultcallback
//
// 阿里通信定制服务，商家发货后进行调用该接口，用于自动发货并确认收货
func Taobaowttradeorderresultcallback(clt *core.SDKClient, req *alicom.TaobaowttradeorderresultcallbackAPIRequest, session string) (*alicom.TaobaowttradeorderresultcallbackAPIResponse, error) {
	var resp alicom.TaobaowttradeorderresultcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
