package eleenterpriserestaurant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// AlibabaEleEnterpriseRestaurantMix 混合搜索店铺
// alibaba.ele.enterprise.restaurant.mix
//
// 混合搜索店铺
func AlibabaEleEnterpriseRestaurantMix(ctx context.Context, clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMixAPIRequest, resp *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMixAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
