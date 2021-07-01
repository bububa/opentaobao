package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
调拨单查询 
taobao.qimen.transferorder.query

调拨单查询
*/
func TaobaoQimenTransferorderQuery(clt *core.SDKClient, req *qimen.TaobaoQimenTransferorderQueryAPIRequest, session string) (*qimen.TaobaoQimenTransferorderQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenTransferorderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
