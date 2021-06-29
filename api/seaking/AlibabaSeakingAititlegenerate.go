package seaking

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/seaking"
)

/* 
标题智能优化 
alibaba.seaking.aititlegenerate

标题智能优化
*/
func AlibabaSeakingAititlegenerate(clt *core.SDKClient, req *seaking.AlibabaSeakingAititlegenerateRequest, session string) (*seaking.AlibabaSeakingAititlegenerateAPIResponse, error) {
    var resp seaking.AlibabaSeakingAititlegenerateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
