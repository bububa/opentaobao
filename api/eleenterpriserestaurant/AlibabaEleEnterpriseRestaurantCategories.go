package eleenterpriserestaurant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

/* 
餐厅分类 
alibaba.ele.enterprise.restaurant.categories

餐厅分类
*/
func AlibabaEleEnterpriseRestaurantCategories(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCategoriesAPIRequest, session string) (*eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCategoriesAPIResponse, error) {
    var resp eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCategoriesAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
