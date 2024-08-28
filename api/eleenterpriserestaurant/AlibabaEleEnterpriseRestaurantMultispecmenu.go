package eleenterpriserestaurant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// AlibabaEleEnterpriseRestaurantMultispecmenu 查询餐厅菜单
// alibaba.ele.enterprise.restaurant.multispecmenu
//
// 查询餐厅菜单
func AlibabaEleEnterpriseRestaurantMultispecmenu(ctx context.Context, clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest, resp *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
