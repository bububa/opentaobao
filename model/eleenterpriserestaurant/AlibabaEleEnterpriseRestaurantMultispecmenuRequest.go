package eleenterpriserestaurant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询餐厅菜单 API请求
alibaba.ele.enterprise.restaurant.multispecmenu

查询餐厅菜单
*/
type AlibabaEleEnterpriseRestaurantMultispecmenuRequest struct {
    model.Params
    // 餐厅ID
    erestaurantId   string
}

// 初始化AlibabaEleEnterpriseRestaurantMultispecmenuRequest对象
func NewAlibabaEleEnterpriseRestaurantMultispecmenuRequest() *AlibabaEleEnterpriseRestaurantMultispecmenuRequest{
    return &AlibabaEleEnterpriseRestaurantMultispecmenuRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantMultispecmenuRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.multispecmenu"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantMultispecmenuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ErestaurantId Setter
// 餐厅ID
func (r *AlibabaEleEnterpriseRestaurantMultispecmenuRequest) SetErestaurantId(erestaurantId string) error {
    r.erestaurantId = erestaurantId
    r.Set("erestaurant_id", erestaurantId)
    return nil
}

// ErestaurantId Getter
func (r AlibabaEleEnterpriseRestaurantMultispecmenuRequest) GetErestaurantId() string {
    return r.erestaurantId
}
