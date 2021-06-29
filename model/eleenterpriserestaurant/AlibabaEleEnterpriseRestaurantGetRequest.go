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
    _geo   string
    // 餐厅ID
    _erestaurantId   string
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
func (r *AlibabaEleEnterpriseRestaurantGetRequest) SetGeo(_geo string) error {
    r._geo = _geo
    r.Set("geo", _geo)
    return nil
}

// Geo Getter
func (r AlibabaEleEnterpriseRestaurantGetRequest) GetGeo() string {
    return r._geo
}
// ErestaurantId Setter
// 餐厅ID
func (r *AlibabaEleEnterpriseRestaurantGetRequest) SetErestaurantId(_erestaurantId string) error {
    r._erestaurantId = _erestaurantId
    r.Set("erestaurant_id", _erestaurantId)
    return nil
}

// ErestaurantId Getter
func (r AlibabaEleEnterpriseRestaurantGetRequest) GetErestaurantId() string {
    return r._erestaurantId
}
