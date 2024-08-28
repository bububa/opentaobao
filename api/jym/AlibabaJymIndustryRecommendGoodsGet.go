package jym

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymIndustryRecommendGoodsGet 获取交易猫推荐商品
// alibaba.jym.industry.recommend.goods.get
//
// 获取交易猫推荐商品
func AlibabaJymIndustryRecommendGoodsGet(ctx context.Context, clt *core.SDKClient, req *jym.AlibabaJymIndustryRecommendGoodsGetAPIRequest, resp *jym.AlibabaJymIndustryRecommendGoodsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
