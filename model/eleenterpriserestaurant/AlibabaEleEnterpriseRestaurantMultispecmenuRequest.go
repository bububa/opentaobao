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
type AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest struct {
    model.Params
    // 餐厅ID
    _erestaurantId   string
}

// 初始化AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest对象
func NewAlibabaEleEnterpriseRestaurantMultispecmenuRequest() *AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest{
    return &AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.multispecmenu"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ErestaurantId Setter
// 餐厅ID
func (r *AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest) SetErestaurantId(_erestaurantId string) error {
    r._erestaurantId = _erestaurantId
    r.Set("erestaurant_id", _erestaurantId)
    return nil
}

// ErestaurantId Getter
func (r AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest) GetErestaurantId() string {
    return r._erestaurantId
}
