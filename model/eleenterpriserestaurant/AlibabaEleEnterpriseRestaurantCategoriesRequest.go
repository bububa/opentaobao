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
type AlibabaEleEnterpriseRestaurantCategoriesRequest struct {
    model.Params
}

// 初始化AlibabaEleEnterpriseRestaurantCategoriesRequest对象
func NewAlibabaEleEnterpriseRestaurantCategoriesRequest() *AlibabaEleEnterpriseRestaurantCategoriesRequest{
    return &AlibabaEleEnterpriseRestaurantCategoriesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantCategoriesRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.restaurant.categories"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantCategoriesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
