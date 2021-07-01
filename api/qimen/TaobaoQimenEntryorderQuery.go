package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
入库单查询接口 
taobao.qimen.entryorder.query

ERP调用接口，查询入库单信息;
*/
func TaobaoQimenEntryorderQuery(clt *core.SDKClient, req *qimen.TaobaoQimenEntryorderQueryAPIRequest, session string) (*qimen.TaobaoQimenEntryorderQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenEntryorderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
