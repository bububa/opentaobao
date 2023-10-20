package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyVerifySignature 微信账号生物认证
// alitrip.merchant.galaxy.verify.signature
//
// 在用户注册等场景下，如果账号风险等级过高，需要进行生物认证
func AlitripMerchantGalaxyVerifySignature(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyVerifySignatureAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyVerifySignatureAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
