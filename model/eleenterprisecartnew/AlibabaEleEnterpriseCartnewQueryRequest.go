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
    phone   string
    // 1212
    latitude   string
    // 1212
    longitude   string
    // 餐厅id
    erestaurantId   string
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
func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r AlibabaEleEnterpriseCartnewQueryRequest) GetPhone() string {
    return r.phone
}
// Latitude Setter
// 1212
func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r AlibabaEleEnterpriseCartnewQueryRequest) GetLatitude() string {
    return r.latitude
}
// Longitude Setter
// 1212
func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r AlibabaEleEnterpriseCartnewQueryRequest) GetLongitude() string {
    return r.longitude
}
// ErestaurantId Setter
// 餐厅id
func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetErestaurantId(erestaurantId string) error {
    r.erestaurantId = erestaurantId
    r.Set("erestaurant_id", erestaurantId)
    return nil
}

// ErestaurantId Getter
func (r AlibabaEleEnterpriseCartnewQueryRequest) GetErestaurantId() string {
    return r.erestaurantId
}
