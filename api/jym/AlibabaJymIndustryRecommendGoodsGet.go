package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymIndustryRecommendGoodsGet 获取交易猫推荐商品
// alibaba.jym.industry.recommend.goods.get
//
// 获取交易猫推荐商品
func AlibabaJymIndustryRecommendGoodsGet(clt *core.SDKClient, req *jym.AlibabaJymIndustryRecommendGoodsGetAPIRequest, session string) (*jym.AlibabaJymIndustryRecommendGoodsGetAPIResponse, error) {
	var resp jym.AlibabaJymIndustryRecommendGoodsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
