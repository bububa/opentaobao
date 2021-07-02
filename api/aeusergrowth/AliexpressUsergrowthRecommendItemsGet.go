package aeusergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aeusergrowth"
)

// AliexpressUsergrowthRecommendItemsGet 第三方平台推荐商品
// aliexpress.usergrowth.recommend.items.get
//
// 第三方平台的推荐AE商品   场景：skin 、底部推荐等
func AliexpressUsergrowthRecommendItemsGet(clt *core.SDKClient, req *aeusergrowth.AliexpressUsergrowthRecommendItemsGetAPIRequest, session string) (*aeusergrowth.AliexpressUsergrowthRecommendItemsGetAPIResponse, error) {
	var resp aeusergrowth.AliexpressUsergrowthRecommendItemsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
