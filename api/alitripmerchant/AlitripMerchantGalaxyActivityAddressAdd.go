package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyactivityaddressadd 星河-营销抽奖奖品邮寄地址添加
// alitrip.merchant.galaxy.activity.address.add
//
// 营销抽奖活动，奖品邮寄地址填写
func Alitripmerchantgalaxyactivityaddressadd(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyactivityaddressaddAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyactivityaddressaddAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyactivityaddressaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
