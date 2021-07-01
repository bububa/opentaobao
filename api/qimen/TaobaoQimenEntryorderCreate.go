package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
入库单创建接口 
taobao.qimen.entryorder.create

ERP调用接口，创建入库单;
*/
func TaobaoQimenEntryorderCreate(clt *core.SDKClient, req *qimen.TaobaoQimenEntryorderCreateAPIRequest, session string) (*qimen.TaobaoQimenEntryorderCreateAPIResponse, error) {
    var resp qimen.TaobaoQimenEntryorderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
