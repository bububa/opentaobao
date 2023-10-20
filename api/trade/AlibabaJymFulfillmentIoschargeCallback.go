package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaJymFulfillmentIoschargeCallback 代充充值回调
// alibaba.jym.fulfillment.ioscharge.callback
//
// 代充充值回调
func AlibabaJymFulfillmentIoschargeCallback(clt *core.SDKClient, req *trade.AlibabaJymFulfillmentIoschargeCallbackAPIRequest, resp *trade.AlibabaJymFulfillmentIoschargeCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
