package eleenterpriserestaurant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

/* 
餐厅列表 
alibaba.ele.enterprise.restaurant.search

餐厅列表
*/
func AlibabaEleEnterpriseRestaurantSearch(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantSearchRequest, session string) (*eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantSearchAPIResponse, error) {
    var resp eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
