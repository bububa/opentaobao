package eleenterpriserestaurant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// AlibabaEleEnterpriseRestaurantGet 查询餐厅信息
// alibaba.ele.enterprise.restaurant.get
//
// 查询餐厅信息
func AlibabaEleEnterpriseRestaurantGet(ctx context.Context, clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantGetAPIRequest, resp *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
