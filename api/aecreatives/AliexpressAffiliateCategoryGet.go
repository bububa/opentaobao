package aecreatives

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aecreatives"
)

/* 
AE流量推广类目信息获取API 
aliexpress.affiliate.category.get

获取AE流量推广类目的API
*/
func AliexpressAffiliateCategoryGet(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateCategoryGetRequest, session string) (*aecreatives.AliexpressAffiliateCategoryGetAPIResponse, error) {
    var resp aecreatives.AliexpressAffiliateCategoryGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
