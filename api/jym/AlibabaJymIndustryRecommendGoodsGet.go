package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// Alibabajymindustryrecommendgoodsget 获取交易猫推荐商品
// alibaba.jym.industry.recommend.goods.get
//
// 获取交易猫推荐商品
func Alibabajymindustryrecommendgoodsget(clt *core.SDKClient, req *jym.AlibabajymindustryrecommendgoodsgetAPIRequest, session string) (*jym.AlibabajymindustryrecommendgoodsgetAPIResponse, error) {
	var resp jym.AlibabajymindustryrecommendgoodsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
