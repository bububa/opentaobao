package shenjing

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shenjing"
)

/* 
访客通过PAD提交访客码 
alibaba.ib.shenjing.visitor.pad.fetchcodeverify

访客通过PAD提交访客码，录脸进入园区。
*/
func AlibabaIbShenjingVisitorPadFetchcodeverify(clt *core.SDKClient, req *shenjing.AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest, session string) (*shenjing.AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse, error) {
    var resp shenjing.AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
