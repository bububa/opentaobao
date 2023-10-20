package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaJymFulfillmentCardCallback 外部商家卡密结果回调
// alibaba.jym.fulfillment.card.callback
//
// 外部商家卡密结果回调
func AlibabaJymFulfillmentCardCallback(clt *core.SDKClient, req *trade.AlibabaJymFulfillmentCardCallbackAPIRequest, session string) (*trade.AlibabaJymFulfillmentCardCallbackAPIResponse, error) {
	var resp trade.AlibabaJymFulfillmentCardCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
