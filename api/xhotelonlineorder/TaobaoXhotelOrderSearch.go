package xhotelonlineorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* 
酒店产品库订单查询 
taobao.xhotel.order.search

酒店产品库订单查询功能，查询90天内的订单
*/
func TaobaoXhotelOrderSearch(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderSearchAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelOrderSearchAPIResponse, error) {
    var resp xhotelonlineorder.TaobaoXhotelOrderSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
