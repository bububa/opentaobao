package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsAddressReachablebatchGetAPIRequest
批量判定服务是否可达 API请求
taobao.logistics.address.reachablebatch.get

批量判定服务是否可达 */
type TaobaoLogisticsAddressReachablebatchGetAPIRequest struct {
	model.Params
	// 筛单用户输入地址结构
	_addressList []AddressReachable
}

// NewTaobaoLogisticsAddressReachablebatchGetRequest 初始化TaobaoLogisticsAddressReachablebatchGetAPIRequest对象
func NewTaobaoLogisticsAddressReachablebatchGetRequest() *TaobaoLogisticsAddressReachablebatchGetAPIRequest {
	return &TaobaoLogisticsAddressReachablebatchGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressReachablebatchGetAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.address.reachablebatch.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressReachablebatchGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AddressList Setter
// 筛单用户输入地址结构
func (r *TaobaoLogisticsAddressReachablebatchGetAPIRequest) SetAddressList(_addressList []AddressReachable) error {
	r._addressList = _addressList
	r.Set("address_list", _addressList)
	return nil
}

// Get AddressList Getter
func (r TaobaoLogisticsAddressReachablebatchGetAPIRequest) GetAddressList() []AddressReachable {
	return r._addressList
}
