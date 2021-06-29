package eleenterpriserestaurant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询餐厅信息 API请求
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

// 初始化AlibabaEleEnterpriseRestaurantGetRequest对象
func NewAlibabaEleEnterpriseRestaurantGetRequest() *AlibabaEleEnterpriseRestaurantGetRequest{
    return &AlibabaEleEnterpriseRestaurantGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantGetRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Geo Setter
// longitude和latitude用英文逗号分隔
func (r *AlibabaEleEnterpriseRestaurantGetRequest) SetGeo(geo string) error {
    r.geo = geo
    r.Set("geo", geo)
    return nil
}

// Geo Getter
func (r AlibabaEleEnterpriseRestaurantGetRequest) GetGeo() string {
    return r.geo
}
// ErestaurantId Setter
// 餐厅ID
func (r *AlibabaEleEnterpriseRestaurantGetRequest) SetErestaurantId(erestaurantId string) error {
    r.erestaurantId = erestaurantId
    r.Set("erestaurant_id", erestaurantId)
    return nil
}

// ErestaurantId Getter
func (r AlibabaEleEnterpriseRestaurantGetRequest) GetErestaurantId() string {
    return r.erestaurantId
}
