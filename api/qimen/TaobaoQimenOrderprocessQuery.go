package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
订单流水查询接口 
taobao.qimen.orderprocess.query

ERP调用订单流水查询接口
*/
func TaobaoQimenOrderprocessQuery(clt *core.SDKClient, req *qimen.TaobaoQimenOrderprocessQueryRequest, session string) (*qimen.TaobaoQimenOrderprocessQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenOrderprocessQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
