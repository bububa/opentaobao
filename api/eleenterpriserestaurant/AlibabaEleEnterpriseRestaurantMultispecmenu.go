package eleenterpriserestaurant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

/* 
查询餐厅菜单 
alibaba.ele.enterprise.restaurant.multispecmenu

查询餐厅菜单
*/
func AlibabaEleEnterpriseRestaurantMultispecmenu(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest, session string) (*eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse, error) {
    var resp eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMultispecmenuAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
