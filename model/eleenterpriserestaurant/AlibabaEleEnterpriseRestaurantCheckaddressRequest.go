package eleenterpriserestaurant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
检查地址是否在餐厅配送范围内 APIRequest
alibaba.ele.enterprise.restaurant.checkaddress

检查地址是否在餐厅配送范围内
*/
type AlibabaEleEnterpriseRestaurantCheckaddressRequest struct {
    model.Params

    // 餐厅Id
    erestaurantId   string 

    // [{"longitude": 1, "latitude": 2}], json 字符串, 每个元素是 一个 dict{longitude, latitude, …} 其他字段原样返回
    addresses   string 

}

func NewAlibabaEleEnterpriseRestaurantCheckaddressRequest() *AlibabaEleEnterpriseRestaurantCheckaddressRequest{
    return &AlibabaEleEnterpriseRestaurantCheckaddressRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseRestaurantCheckaddressRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.checkaddress"
}

func (r AlibabaEleEnterpriseRestaurantCheckaddressRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseRestaurantCheckaddressRequest) SetErestaurantId(erestaurantId string) error {
    r.erestaurantId = erestaurantId
    r.Set("erestaurant_id", erestaurantId)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantCheckaddressRequest) GetErestaurantId() string {
    return r.erestaurantId
}

func (r *AlibabaEleEnterpriseRestaurantCheckaddressRequest) SetAddresses(addresses string) error {
    r.addresses = addresses
    r.Set("addresses", addresses)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantCheckaddressRequest) GetAddresses() string {
    return r.addresses
}

