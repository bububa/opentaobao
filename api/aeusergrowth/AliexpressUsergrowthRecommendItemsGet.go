package aeusergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aeusergrowth"
)

// Aliexpressusergrowthrecommenditemsget 第三方平台推荐商品
// aliexpress.usergrowth.recommend.items.get
//
// 第三方平台的推荐AE商品   场景：skin 、底部推荐等
func Aliexpressusergrowthrecommenditemsget(clt *core.SDKClient, req *aeusergrowth.AliexpressusergrowthrecommenditemsgetAPIRequest, session string) (*aeusergrowth.AliexpressusergrowthrecommenditemsgetAPIResponse, error) {
	var resp aeusergrowth.AliexpressusergrowthrecommenditemsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
