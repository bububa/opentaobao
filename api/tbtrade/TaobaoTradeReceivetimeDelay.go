package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// Taobaotradereceivetimedelay 延长交易收货时间
// taobao.trade.receivetime.delay
//
// 延长交易收货时间
func Taobaotradereceivetimedelay(clt *core.SDKClient, req *tbtrade.TaobaotradereceivetimedelayAPIRequest, session string) (*tbtrade.TaobaotradereceivetimedelayAPIResponse, error) {
	var resp tbtrade.TaobaotradereceivetimedelayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
