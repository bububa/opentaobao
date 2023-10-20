package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleapprizeorderfulfillment 鉴定担保资金订单履约
// alibaba.idle.apprize.order.fulfillment
//
// 服务商针对自己的服务订单进行履约
func Alibabaidleapprizeorderfulfillment(clt *core.SDKClient, req *idle.AlibabaidleapprizeorderfulfillmentAPIRequest, session string) (*idle.AlibabaidleapprizeorderfulfillmentAPIResponse, error) {
	var resp idle.AlibabaidleapprizeorderfulfillmentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
