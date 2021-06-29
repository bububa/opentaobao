package eleenterprisecartnew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新版购物车查询 APIRequest
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

func NewAlibabaEleEnterpriseCartnewQueryRequest() *AlibabaEleEnterpriseCartnewQueryRequest{
    return &AlibabaEleEnterpriseCartnewQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseCartnewQueryRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.cartnew.query"
}

func (r AlibabaEleEnterpriseCartnewQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r AlibabaEleEnterpriseCartnewQueryRequest) GetPhone() string {
    return r.phone
}

func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r AlibabaEleEnterpriseCartnewQueryRequest) GetLatitude() string {
    return r.latitude
}

func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r AlibabaEleEnterpriseCartnewQueryRequest) GetLongitude() string {
    return r.longitude
}

func (r *AlibabaEleEnterpriseCartnewQueryRequest) SetErestaurantId(erestaurantId string) error {
    r.erestaurantId = erestaurantId
    r.Set("erestaurant_id", erestaurantId)
    return nil
}

func (r AlibabaEleEnterpriseCartnewQueryRequest) GetErestaurantId() string {
    return r.erestaurantId
}

