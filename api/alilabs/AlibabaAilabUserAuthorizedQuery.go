package alilabs

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alilabs"
)

/* 
查询授权状态接口 
alibaba.ailab.user.authorized.query

查询三方用户授权状态
*/
func AlibabaAilabUserAuthorizedQuery(clt *core.SDKClient, req *alilabs.AlibabaAilabUserAuthorizedQueryAPIRequest, session string) (*alilabs.AlibabaAilabUserAuthorizedQueryAPIResponse, error) {
    var resp alilabs.AlibabaAilabUserAuthorizedQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
