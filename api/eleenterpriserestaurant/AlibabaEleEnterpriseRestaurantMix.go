package eleenterpriserestaurant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// AlibabaEleEnterpriseRestaurantMix 混合搜索店铺
// alibaba.ele.enterprise.restaurant.mix
//
// 混合搜索店铺
func AlibabaEleEnterpriseRestaurantMix(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMixAPIRequest, resp *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMixAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
