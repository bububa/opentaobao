package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
内购服务确认单申请接口 
taobao.fuwu.sp.confirm.apply

isv能通过该接口发起确认申请单
*/
func TaobaoFuwuSpConfirmApply(clt *core.SDKClient, req *servicecenter.TaobaoFuwuSpConfirmApplyRequest, session string) (*servicecenter.TaobaoFuwuSpConfirmApplyResponse, error) {
    var resp servicecenter.TaobaoFuwuSpConfirmApplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
