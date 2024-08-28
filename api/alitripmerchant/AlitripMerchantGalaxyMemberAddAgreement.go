package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberAddAgreement 添加用户协议记录接口
// alitrip.merchant.galaxy.member.add.agreement
//
// 记录用户是否授权协议
func AlitripMerchantGalaxyMemberAddAgreement(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberAddAgreementAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberAddAgreementAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
