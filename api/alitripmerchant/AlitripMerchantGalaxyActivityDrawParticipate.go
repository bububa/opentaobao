package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyactivitydrawparticipate 星河-幸运抽奖活动参与
// alitrip.merchant.galaxy.activity.draw.participate
//
// 雅高小程序幸运抽奖活动页面用户进行抽奖，根据活动规则返回抽奖结果
func Alitripmerchantgalaxyactivitydrawparticipate(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyactivitydrawparticipateAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyactivitydrawparticipateAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyactivitydrawparticipateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
