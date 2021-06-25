package category

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/category"
)

/* 
获取类目信息 
alibaba.wholesale.category.get

获取类目信息
*/
func AlibabaWholesaleCategoryGet(clt *core.SDKClient, req *category.AlibabaWholesaleCategoryGetRequest, session string) (*category.AlibabaWholesaleCategoryGetResponse, error) {
    var resp category.AlibabaWholesaleCategoryGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
