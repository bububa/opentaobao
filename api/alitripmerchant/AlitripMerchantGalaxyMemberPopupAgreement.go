package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberPopupAgreement 小程序唤起协议弹窗
// alitrip.merchant.galaxy.member.popup.agreement
//
// 用户进入小程序后，根据用户是否授权协议，判断是否唤起协议弹窗
func AlitripMerchantGalaxyMemberPopupAgreement(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberPopupAgreementAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberPopupAgreementAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
