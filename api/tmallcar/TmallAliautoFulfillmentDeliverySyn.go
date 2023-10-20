package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoFulfillmentDeliverySyn 交付状态及物流信息同步
// tmall.aliauto.fulfillment.delivery.syn
//
// 交付状态及物流信息同步
func TmallAliautoFulfillmentDeliverySyn(clt *core.SDKClient, req *tmallcar.TmallAliautoFulfillmentDeliverySynAPIRequest, resp *tmallcar.TmallAliautoFulfillmentDeliverySynAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
