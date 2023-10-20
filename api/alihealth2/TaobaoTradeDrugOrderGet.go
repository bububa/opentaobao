package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Taobaotradedrugorderget 查看订单详情
// taobao.trade.drug.order.get
//
// 商家查看订单详情
func Taobaotradedrugorderget(clt *core.SDKClient, req *alihealth2.TaobaotradedrugordergetAPIRequest, session string) (*alihealth2.TaobaotradedrugordergetAPIResponse, error) {
	var resp alihealth2.TaobaotradedrugordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
