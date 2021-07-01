package eleenterpriserestaurant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest
查询餐厅菜单 API请求
alibaba.ele.enterprise.restaurant.multispecmenu

查询餐厅菜单 */
type AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest struct {
	model.Params
	// 餐厅ID
	_erestaurantId string
}

// New
