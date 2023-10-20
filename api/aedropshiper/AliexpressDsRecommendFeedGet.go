package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpressdsrecommendfeedget 获取推荐商品信息流接口
// aliexpress.ds.recommend.feed.get
//
// 获取推荐商品信息流
func Aliexpressdsrecommendfeedget(clt *core.SDKClient, req *aedropshiper.AliexpressdsrecommendfeedgetAPIRequest, session string) (*aedropshiper.AliexpressdsrecommendfeedgetAPIResponse, error) {
	var resp aedropshiper.AliexpressdsrecommendfeedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
