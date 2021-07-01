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
type AlibabaEleEnterpriseRestaurantGetAPIRequest struct {
    model.Params
    // longitude和latitude用英文逗号分隔
    _geo   string
    // 餐厅ID
    _erestaurantId   string
}

// 初始化AlibabaEleEnterpriseRestaurantGetAPIRequest对象
func NewAlibabaEleEnterpriseRestaurantGetRequest() *AlibabaEleEnterpriseRestaurantGetAPIRequest{
    return &AlibabaEleEnterpriseRestaurantGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantGetAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Geo Setter
// longitude和latitude用英文逗号分隔
func (r *AlibabaEleEnterpriseRestaurantGetAPIRequest) SetGeo(_geo string) error {
    r._geo = _geo
    r.Set("geo", _geo)
    return nil
}

// Geo Getter
func (r AlibabaEleEnterpriseRestaurantGetAPIRequest) GetGeo() string {
    return r._geo
}
// ErestaurantId Setter
// 餐厅ID
func (r *AlibabaEleEnterpriseRestaurantGetAPIRequest) SetErestaurantId(_erestaurantId string) error {
    r._erestaurantId = _erestaurantId
    r.Set("erestaurant_id", _erestaurantId)
    return nil
}

// ErestaurantId Getter
func (r AlibabaEleEnterpriseRestaurantGetAPIRequest) GetErestaurantId() string {
    return r._erestaurantId
}
