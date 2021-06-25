package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
获取流量档位 
alibaba.aliqin.flow.wallet.grade

获取直充流量档位
*/
func AlibabaAliqinFlowWalletGrade(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletGradeRequest, session string) (*alicom.AlibabaAliqinFlowWalletGradeResponse, error) {
    var resp alicom.AlibabaAliqinFlowWalletGradeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
