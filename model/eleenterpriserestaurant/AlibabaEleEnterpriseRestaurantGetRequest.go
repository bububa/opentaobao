package eleenterpriserestaurant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询餐厅信息 APIRequest
alibaba.ele.enterprise.restaurant.get

查询餐厅信息
*/
type AlibabaEleEnterpriseRestaurantGetRequest struct {
    model.Params

    // longitude和latitude用英文逗号分隔
    geo   string 

    // 餐厅ID
    erestaurantId   string 

}

func NewAlibabaEleEnterpriseRestaurantGetRequest() *AlibabaEleEnterpriseRestaurantGetRequest{
    return &AlibabaEleEnterpriseRestaurantGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseRestaurantGetRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.get"
}

func (r AlibabaEleEnterpriseRestaurantGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseRestaurantGetRequest) SetGeo(geo string) error {
    r.geo = geo
    r.Set("geo", geo)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantGetRequest) GetGeo() string {
    return r.geo
}

func (r *AlibabaEleEnterpriseRestaurantGetRequest) SetErestaurantId(erestaurantId string) error {
    r.erestaurantId = erestaurantId
    r.Set("erestaurant_id", erestaurantId)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantGetRequest) GetErestaurantId() string {
    return r.erestaurantId
}

