package eleenterpriserestaurant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询餐厅菜单 APIRequest
alibaba.ele.enterprise.restaurant.multispecmenu

查询餐厅菜单
*/
type AlibabaEleEnterpriseRestaurantMultispecmenuRequest struct {
    model.Params

    // 餐厅ID
    erestaurantId   string 

}

func NewAlibabaEleEnterpriseRestaurantMultispecmenuRequest() *AlibabaEleEnterpriseRestaurantMultispecmenuRequest{
    return &AlibabaEleEnterpriseRestaurantMultispecmenuRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseRestaurantMultispecmenuRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.multispecmenu"
}

func (r AlibabaEleEnterpriseRestaurantMultispecmenuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseRestaurantMultispecmenuRequest) SetErestaurantId(erestaurantId string) error {
    r.erestaurantId = erestaurantId
    r.Set("erestaurant_id", erestaurantId)
    return nil
}

func (r AlibabaEleEnterpriseRestaurantMultispecmenuRequest) GetErestaurantId() string {
    return r.erestaurantId
}

