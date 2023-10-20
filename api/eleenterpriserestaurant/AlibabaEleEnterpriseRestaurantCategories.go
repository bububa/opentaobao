package eleenterpriserestaurant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// AlibabaEleEnterpriseRestaurantCategories 餐厅分类
// alibaba.ele.enterprise.restaurant.categories
//
// 餐厅分类
func AlibabaEleEnterpriseRestaurantCategories(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCategoriesAPIRequest, resp *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCategoriesAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
