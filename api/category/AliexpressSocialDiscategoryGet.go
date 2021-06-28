package category

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/category"
)

/* 
展示类目获取接口 
aliexpress.social.discategory.get

AE展示类目获取接口
*/
func AliexpressSocialDiscategoryGet(clt *core.SDKClient, req *category.AliexpressSocialDiscategoryGetRequest, session string) (*category.AliexpressSocialDiscategoryGetAPIResponse, error) {
    var resp category.AliexpressSocialDiscategoryGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
