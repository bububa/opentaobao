package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyQueryDrawSummary 星河-抽奖活动概要列表查询
// alitrip.merchant.galaxy.query.draw.summary
//
// 雅高小程序抽奖活动列表概要查询
func AlitripMerchantGalaxyQueryDrawSummary(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyQueryDrawSummaryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyQueryDrawSummaryAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyQueryDrawSummaryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
