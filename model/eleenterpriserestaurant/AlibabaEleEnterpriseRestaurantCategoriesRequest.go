package eleenterpriserestaurant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
餐厅分类 APIRequest
alibaba.ele.enterprise.restaurant.categories

餐厅分类
*/
type AlibabaEleEnterpriseRestaurantCategoriesRequest struct {
    model.Params

}

func NewAlibabaEleEnterpriseRestaurantCategoriesRequest() *AlibabaEleEnterpriseRestaurantCategoriesRequest{
    return &AlibabaEleEnterpriseRestaurantCategoriesRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseRestaurantCategoriesRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.categories"
}

func (r AlibabaEleEnterpriseRestaurantCategoriesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


