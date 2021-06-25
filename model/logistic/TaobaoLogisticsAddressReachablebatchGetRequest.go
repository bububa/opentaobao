package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量判定服务是否可达 APIRequest
taobao.logistics.address.reachablebatch.get

批量判定服务是否可达
*/
type TaobaoLogisticsAddressReachablebatchGetRequest struct {
    model.Params

    // 筛单用户输入地址结构
    addressList   []AddressReachable 

}

func NewTaobaoLogisticsAddressReachablebatchGetRequest() *TaobaoLogisticsAddressReachablebatchGetRequest{
    return &TaobaoLogisticsAddressReachablebatchGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsAddressReachablebatchGetRequest) GetApiMethodName() string {
    return "taobao.logistics.address.reachablebatch.get"
}

func (r TaobaoLogisticsAddressReachablebatchGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsAddressReachablebatchGetRequest) SetAddressList(addressList []AddressReachable) error {
    r.addressList = addressList
    r.Set("address_list", addressList)
    return nil
}

func (r TaobaoLogisticsAddressReachablebatchGetRequest) GetAddressList() []AddressReachable {
    return r.addressList
}

