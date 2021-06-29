package alisports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alisports"
)

/* 
用户核销 
alibaba.alisports.efsp.userwriteoff

用户核销
*/
func AlibabaAlisportsEfspUserwriteoff(clt *core.SDKClient, req *alisports.AlibabaAlisportsEfspUserwriteoffRequest, session string) (*alisports.AlibabaAlisportsEfspUserwriteoffAPIResponse, error) {
    var resp alisports.AlibabaAlisportsEfspUserwriteoffAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
