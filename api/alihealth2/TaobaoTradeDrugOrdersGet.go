package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Taobaotradedrugordersget 阿里健康获取某一药店全部订单
// taobao.trade.drug.orders.get
//
// 阿里健康获取某一药店全部订单
func Taobaotradedrugordersget(clt *core.SDKClient, req *alihealth2.TaobaotradedrugordersgetAPIRequest, session string) (*alihealth2.TaobaotradedrugordersgetAPIResponse, error) {
	var resp alihealth2.TaobaotradedrugordersgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
