package eleenterpriserestaurant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

// AlibabaEleEnterpriseRestaurantGet 查询餐厅信息
// alibaba.ele.enterprise.restaurant.get
//
// 查询餐厅信息
func AlibabaEleEnterpriseRestaurantGet(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantGetAPIRequest, resp *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
