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
type TaobaoLogisticsAddressReachablebatchGetRequest struct {
    model.Params
    // 筛单用户输入地址结构
    addressList   []AddressReachable
}

// 初始化TaobaoLogisticsAddressReachablebatchGetRequest对象
func NewTaobaoLogisticsAddressReachablebatchGetRequest() *TaobaoLogisticsAddressReachablebatchGetRequest{
    return &TaobaoLogisticsAddressReachablebatchGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressReachablebatchGetRequest) GetApiMethodName() string {
    return "taobao.logistics.address.reachablebatch.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressReachablebatchGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AddressList Setter
// 筛单用户输入地址结构
func (r *TaobaoLogisticsAddressReachablebatchGetRequest) SetAddressList(addressList []AddressReachable) error {
    r.addressList = addressList
    r.Set("address_list", addressList)
    return nil
}

// AddressList Getter
func (r TaobaoLogisticsAddressReachablebatchGetRequest) GetAddressList() []AddressReachable {
    return r.addressList
}
