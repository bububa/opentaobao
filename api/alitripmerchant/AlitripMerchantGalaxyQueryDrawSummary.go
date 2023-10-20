package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyquerydrawsummary 星河-抽奖活动概要列表查询
// alitrip.merchant.galaxy.query.draw.summary
//
// 雅高小程序抽奖活动列表概要查询
func Alitripmerchantgalaxyquerydrawsummary(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyquerydrawsummaryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyquerydrawsummaryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyquerydrawsummaryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
