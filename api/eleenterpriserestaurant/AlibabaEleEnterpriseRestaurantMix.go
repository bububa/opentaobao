package eleenterpriserestaurant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterpriserestaurant"
)

/* 
混合搜索店铺 
alibaba.ele.enterprise.restaurant.mix

混合搜索店铺
*/
func AlibabaEleEnterpriseRestaurantMix(clt *core.SDKClient, req *eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMixRequest, session string) (*eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMixAPIResponse, error) {
    var resp eleenterpriserestaurant.AlibabaEleEnterpriseRestaurantMixAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
