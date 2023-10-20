package eleenterpriserestaurant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// AlibabaEleEnterpriseRestaurantCheckaddress 检查地址是否在餐厅配送范围内
// alibaba.ele.enterprise.restaurant.checkaddress
//
// 检查地址是否在餐厅配送范围内
func AlibabaEleEnterpriseRestaurantCheckaddress(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest, resp *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
