package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxymessagesubscriptionstorage 存储模版ID
// alitrip.merchant.galaxy.message.subscription.storage
//
// 消息订阅中的消息模版的存储
func Alitripmerchantgalaxymessagesubscriptionstorage(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxymessagesubscriptionstorageAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxymessagesubscriptionstorageAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxymessagesubscriptionstorageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
