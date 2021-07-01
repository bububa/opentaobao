package eleenterpriserestaurant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

/* 
查询餐厅信息 
alibaba.ele.enterprise.restaurant.get

查询餐厅信息
*/
func AlibabaEleEnterpriseRestaurantGet(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantGetAPIRequest, session string) (*eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantGetAPIResponse, error) {
    var resp eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
