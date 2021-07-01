package eleenterpriserestaurant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
餐厅分类 API请求
alibaba.ele.enterprise.restaurant.categories

餐厅分类
*/
type AlibabaEleEnterpriseRestaurantCategoriesAPIRequest struct {
    model.Params
}

// 初始化AlibabaEleEnterpriseRestaurantCategoriesAPIRequest对象
func NewAlibabaEleEnterpriseRestaurantCategoriesRequest() *AlibabaEleEnterpriseRestaurantCategoriesAPIRequest{
    return &AlibabaEleEnterpriseRestaurantCategoriesAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantCategoriesAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.categories"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantCategoriesAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
