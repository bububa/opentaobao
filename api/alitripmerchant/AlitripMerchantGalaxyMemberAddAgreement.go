package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxymemberaddagreement 添加用户协议记录接口
// alitrip.merchant.galaxy.member.add.agreement
//
// 记录用户是否授权协议
func Alitripmerchantgalaxymemberaddagreement(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxymemberaddagreementAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxymemberaddagreementAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxymemberaddagreementAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
