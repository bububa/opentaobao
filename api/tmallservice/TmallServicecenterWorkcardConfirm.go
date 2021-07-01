package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
服务商确认服务完成 
tmall.servicecenter.workcard.confirm

提供给外部合作服务商，用于通知天猫，告知寄修服务厂内操作全部完成
*/
func TmallServicecenterWorkcardConfirm(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardConfirmAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardConfirmAPIResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
