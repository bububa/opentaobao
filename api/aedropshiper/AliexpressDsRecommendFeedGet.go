package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsRecommendFeedGet 获取推荐商品信息流接口
// aliexpress.ds.recommend.feed.get
//
// 获取推荐商品信息流
func AliexpressDsRecommendFeedGet(clt *core.SDKClient, req *aedropshiper.AliexpressDsRecommendFeedGetAPIRequest, session string) (*aedropshiper.AliexpressDsRecommendFeedGetAPIResponse, error) {
	var resp aedropshiper.AliexpressDsRecommendFeedGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
