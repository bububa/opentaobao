package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberAddAgreement 添加用户协议记录接口
// alitrip.merchant.galaxy.member.add.agreement
//
// 记录用户是否授权协议
func AlitripMerchantGalaxyMemberAddAgreement(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberAddAgreementAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyMemberAddAgreementAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyMemberAddAgreementAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
