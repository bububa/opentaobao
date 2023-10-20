package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityAddressAdd 星河-营销抽奖奖品邮寄地址添加
// alitrip.merchant.galaxy.activity.address.add
//
// 营销抽奖活动，奖品邮寄地址填写
func AlitripMerchantGalaxyActivityAddressAdd(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityAddressAddAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyActivityAddressAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
