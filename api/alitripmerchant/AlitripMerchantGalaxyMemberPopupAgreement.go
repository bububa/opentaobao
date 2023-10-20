package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberPopupAgreement 小程序唤起协议弹窗
// alitrip.merchant.galaxy.member.popup.agreement
//
// 用户进入小程序后，根据用户是否授权协议，判断是否唤起协议弹窗
func AlitripMerchantGalaxyMemberPopupAgreement(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberPopupAgreementAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyMemberPopupAgreementAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyMemberPopupAgreementAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
