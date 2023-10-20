package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautofulfillmentdeliverysyn 交付状态及物流信息同步
// tmall.aliauto.fulfillment.delivery.syn
//
// 交付状态及物流信息同步
func Tmallaliautofulfillmentdeliverysyn(clt *core.SDKClient, req *tmallcar.TmallaliautofulfillmentdeliverysynAPIRequest, session string) (*tmallcar.TmallaliautofulfillmentdeliverysynAPIResponse, error) {
	var resp tmallcar.TmallaliautofulfillmentdeliverysynAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
