package tmallcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcar"
)

/* 
接收授信结果通知 
tmall.aliauto.autofinance.credit.receive

天猫汽车的金融业务场景中，需要接收外部ISV对用户授信申请的通知结果.
*/
func TmallAliautoAutofinanceCreditReceive(clt *core.SDKClient, req *tmallcar.TmallAliautoAutofinanceCreditReceiveAPIRequest, session string) (*tmallcar.TmallAliautoAutofinanceCreditReceiveAPIResponse, error) {
    var resp tmallcar.TmallAliautoAutofinanceCreditReceiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
