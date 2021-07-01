package tmallcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcar"
)

/* 
isv查询服务工单详情 
tmall.aliauto.service.receipt.get

isv查询服务工单详情
*/
func TmallAliautoServiceReceiptGet(clt *core.SDKClient, req *tmallcar.TmallAliautoServiceReceiptGetAPIRequest, session string) (*tmallcar.TmallAliautoServiceReceiptGetAPIResponse, error) {
    var resp tmallcar.TmallAliautoServiceReceiptGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
