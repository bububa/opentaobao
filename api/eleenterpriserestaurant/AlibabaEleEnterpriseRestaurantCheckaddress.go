package eleenterpriserestaurant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

/* 
检查地址是否在餐厅配送范围内 
alibaba.ele.enterprise.restaurant.checkaddress

检查地址是否在餐厅配送范围内
*/
func AlibabaEleEnterpriseRestaurantCheckaddress(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCheckaddressRequest, session string) (*eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse, error) {
    var resp eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
