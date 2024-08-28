package eleenterpriserestaurant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// AlibabaEleEnterpriseRestaurantCheckaddress 检查地址是否在餐厅配送范围内
// alibaba.ele.enterprise.restaurant.checkaddress
//
// 检查地址是否在餐厅配送范围内
func AlibabaEleEnterpriseRestaurantCheckaddress(ctx context.Context, clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest, resp *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
