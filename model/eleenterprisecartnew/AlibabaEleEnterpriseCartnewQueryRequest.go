package eleenterprisecartnew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新版购物车查询 API请求
alibaba.ele.enterprise.cartnew.query

新版购物车查询
*/
type AlibabaEleEnterpriseCartnewQueryRequest struct {
    model.Params
    // 1212
    _phone   string
    // 1212
    _latitude   string
    // 1212
    _longitude   string
    // 餐厅id
    _erestaurantId   string
}

// 初始化AlibabaEleEnterpriseCartnewQueryRequest对象
func NewAlibabaEleEnterpriseCartnewQueryRequest() *AlibabaEleEnterpriseCartnewQueryRequest{
    return &AlibabaEleEnterpriseCartnewQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseCartnewQueryRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.cartnew.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseCartnewQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Phone Setter
// 1212
func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseCartnewQueryRequest) GetPhone() string {
    return r._phone
}
// Latitude Setter
// 1212
func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r AlibabaEleEnterpriseCartnewQueryRequest) GetLatitude() string {
    return r._latitude
}
// Longitude Setter
// 1212
func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r AlibabaEleEnterpriseCartnewQueryRequest) GetLongitude() string {
    return r._longitude
}
// ErestaurantId Setter
// 餐厅id
func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetErestaurantId(_erestaurantId string) error {
    r._erestaurantId = _erestaurantId
    r.Set("erestaurant_id", _erestaurantId)
    return nil
}

// ErestaurantId Getter
func (r AlibabaEleEnterpriseCartnewQueryRequest) GetErestaurantId() string {
    return r._erestaurantId
}
