package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量判定服务是否可达 API请求
taobao.logistics.address.reachablebatch.get

批量判定服务是否可达
*/
type TaobaoLogisticsAddressReachablebatchGetAPIRequest struct {
    model.Params
    // 筛单用户输入地址结构
    _addressList   []AddressReachable
}

// 初始化TaobaoLogisticsAddressReachablebatchGetAPIRequest对象
func NewTaobaoLogisticsAddressReachablebatchGetRequest() *TaobaoLogisticsAddressReachablebatchGetAPIRequest{
    return &TaobaoLogisticsAddressReachablebatchGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressReachablebatchGetAPIRequest) GetApiMethodName() string {
    return "taobao.logistics.address.reachablebatch.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressReachablebatchGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AddressList Setter
// 筛单用户输入地址结构
func (r *TaobaoLogisticsAddressReachablebatchGetAPIRequest) SetAddressList(_addressList []AddressReachable) error {
    r._addressList = _addressList
    r.Set("address_list", _addressList)
    return nil
}

// AddressList Getter
func (r TaobaoLogisticsAddressReachablebatchGetAPIRequest) GetAddressList() []AddressReachable {
    return r._addressList
}
