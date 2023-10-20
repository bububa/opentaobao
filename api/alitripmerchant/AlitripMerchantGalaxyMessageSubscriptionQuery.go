package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxymessagesubscriptionquery 查询用户是否有模版ID权限
// alitrip.merchant.galaxy.message.subscription.query
//
// 只是查询用户是否拥有权限ID
func Alitripmerchantgalaxymessagesubscriptionquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxymessagesubscriptionqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxymessagesubscriptionqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxymessagesubscriptionqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
