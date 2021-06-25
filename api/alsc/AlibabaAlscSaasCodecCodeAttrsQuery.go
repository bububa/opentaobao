package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
码业务属性查询 
alibaba.alsc.saas.codec.code.attrs.query

码业务属性查询
*/
func AlibabaAlscSaasCodecCodeAttrsQuery(clt *core.SDKClient, req *alsc.AlibabaAlscSaasCodecCodeAttrsQueryRequest, session string) (*alsc.AlibabaAlscSaasCodecCodeAttrsQueryResponse, error) {
    var resp alsc.AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
