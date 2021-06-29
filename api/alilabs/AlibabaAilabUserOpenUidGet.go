package alilabs

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alilabs"
)

/* 
access token 获取精灵用户 id 
alibaba.ailab.user.open.uid.get

access token 获取精灵用户 id
*/
func AlibabaAilabUserOpenUidGet(clt *core.SDKClient, req *alilabs.AlibabaAilabUserOpenUidGetRequest, session string) (*alilabs.AlibabaAilabUserOpenUidGetAPIResponse, error) {
    var resp alilabs.AlibabaAilabUserOpenUidGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
