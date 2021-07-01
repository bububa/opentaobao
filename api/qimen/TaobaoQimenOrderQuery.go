package qimen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qimen"
)

/* 
根据收件人信息查询交易单号接口 
taobao.qimen.order.query

WMS 调用该接口，根据收件人信息查询平台交易订单号。
*/
func TaobaoQimenOrderQuery(clt *core.SDKClient, req *qimen.TaobaoQimenOrderQueryAPIRequest, session string) (*qimen.TaobaoQimenOrderQueryAPIResponse, error) {
    var resp qimen.TaobaoQimenOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
