package tmallcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcar"
)

/* 
服务工单状态更新 
tmall.aliauto.receipt.state.update

二轮车服务工单状态更新
*/
func TmallAliautoReceiptStateUpdate(clt *core.SDKClient, req *tmallcar.TmallAliautoReceiptStateUpdateRequest, session string) (*tmallcar.TmallAliautoReceiptStateUpdateAPIResponse, error) {
    var resp tmallcar.TmallAliautoReceiptStateUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
