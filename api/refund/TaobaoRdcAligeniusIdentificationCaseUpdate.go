package refund

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/refund"
)

/* 
鉴定工单信息同步 
taobao.rdc.aligenius.identification.case.update

同步商家鉴定工单信息
*/
func TaobaoRdcAligeniusIdentificationCaseUpdate(clt *core.SDKClient, req *refund.TaobaoRdcAligeniusIdentificationCaseUpdateRequest, session string) (*refund.TaobaoRdcAligeniusIdentificationCaseUpdateAPIResponse, error) {
    var resp refund.TaobaoRdcAligeniusIdentificationCaseUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
