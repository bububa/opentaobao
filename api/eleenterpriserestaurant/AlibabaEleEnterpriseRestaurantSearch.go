package eleenterpriserestaurant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// AlibabaEleEnterpriseRestaurantSearch 餐厅列表
// alibaba.ele.enterprise.restaurant.search
//
// 餐厅列表
func AlibabaEleEnterpriseRestaurantSearch(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantSearchAPIRequest, resp *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
