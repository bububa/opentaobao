package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
入库单确认接口 
taobao.qimen.entryorder.confirm

WMS调用接口，回传入库单信息;
*/
func TaobaoQimenEntryorderConfirm(clt *core.SDKClient, req *qimen.TaobaoQimenEntryorderConfirmRequest, session string) (*qimen.TaobaoQimenEntryorderConfirmAPIResponse, error) {
    var resp qimen.TaobaoQimenEntryorderConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
