package eleenterpriserestaurant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseRestaurantGetAPIRequest
查询餐厅信息 API请求
alibaba.ele.enterprise.restaurant.get

查询餐厅信息 */
type AlibabaEleEnterpriseRestaurantGetAPIRequest struct {
	model.Params
	// longitude和latitude用英文逗号分隔
	_geo string
	// 餐厅ID
	_erestaurantId string
}

// New
