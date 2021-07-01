package eleenterpriserestaurant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest
检查地址是否在餐厅配送范围内 API请求
alibaba.ele.enterprise.restaurant.checkaddress

检查地址是否在餐厅配送范围内 */
type AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest struct {
	model.Params
	// 餐厅Id
	_erestaurantId string
	// [{"longitude": 1, "latitude": 2}], json 字符串, 每个元素是 一个 dict{longitude, latitude, …} 其他字段原样返回
	_addresses string
}

// New
