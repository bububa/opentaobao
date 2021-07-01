package alisports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alisports"
)

/* 
计算补助金额 
alibaba.alisports.efsp.countsubsidy

计算补助金额
*/
func AlibabaAlisportsEfspCountsubsidy(clt *core.SDKClient, req *alisports.AlibabaAlisportsEfspCountsubsidyAPIRequest, session string) (*alisports.AlibabaAlisportsEfspCountsubsidyAPIResponse, error) {
    var resp alisports.AlibabaAlisportsEfspCountsubsidyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
