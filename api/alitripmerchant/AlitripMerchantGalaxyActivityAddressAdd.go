package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityAddressAdd 星河-营销抽奖奖品邮寄地址添加
// alitrip.merchant.galaxy.activity.address.add
//
// 营销抽奖活动，奖品邮寄地址填写
func AlitripMerchantGalaxyActivityAddressAdd(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityAddressAddAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyActivityAddressAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
