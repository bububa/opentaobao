package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
蜂鸟询用户T 
alibaba.ele.fengniao.user.time.query

蜂鸟询用户T
*/
func AlibabaEleFengniaoUserTimeQuery(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoUserTimeQueryRequest, session string) (*logistic.AlibabaEleFengniaoUserTimeQueryResponse, error) {
    var resp logistic.AlibabaEleFengniaoUserTimeQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
