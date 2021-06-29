package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
民宿门店信息添加 
taobao.xhotel.bnbhouse.add

添加和更新民宿门店的信息
*/
func TaobaoXhotelBnbhouseAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbhouseAddRequest, session string) (*xhotelitem.TaobaoXhotelBnbhouseAddAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelBnbhouseAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
