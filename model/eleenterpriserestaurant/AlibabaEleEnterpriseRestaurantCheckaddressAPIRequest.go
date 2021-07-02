package eleenterpriserestaurant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest 检查地址是否在餐厅配送范围内 API请求
// alibaba.ele.enterprise.restaurant.checkaddress
//
// 检查地址是否在餐厅配送范围内
type AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest struct {
	model.Params
	// 餐厅Id
	_erestaurantId string
	// [{"longitude": 1, "latitude": 2}], json 字符串, 每个元素是 一个 dict{longitude, latitude, …} 其他字段原样返回
	_addresses string
}

// NewAlibabaEleEnterpriseRestaurantCheckaddressRequest 初始化AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest对象
func NewAlibabaEleEnterpriseRestaurantCheckaddressRequest() *AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest {
	return &AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.restaurant.checkaddress"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ErestaurantId Setter
// 餐厅Id
func (r *AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest) SetErestaurantId(_erestaurantId string) error {
	r._erestaurantId = _erestaurantId
	r.Set("erestaurant_id", _erestaurantId)
	return nil
}

// Get ErestaurantId Getter
func (r AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest) GetErestaurantId() string {
	return r._erestaurantId
}

// Set is Addresses Setter
// [{"longitude": 1, "latitude": 2}], json 字符串, 每个元素是 一个 dict{longitude, latitude, …} 其他字段原样返回
func (r *AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest) SetAddresses(_addresses string) error {
	r._addresses = _addresses
	r.Set("addresses", _addresses)
	return nil
}

// Get Addresses Getter
func (r AlibabaEleEnterpriseRestaurantCheckaddressAPIRequest) GetAddresses() string {
	return r._addresses
}
