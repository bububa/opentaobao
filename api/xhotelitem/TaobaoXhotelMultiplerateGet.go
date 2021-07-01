package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
复杂房价查询接口 
taobao.xhotel.multiplerate.get

查询复杂房价，支持通过入住人数，连住天数，商品信息，房价信息查询
*/
func TaobaoXhotelMultiplerateGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelMultiplerateGetAPIRequest, session string) (*xhotelitem.TaobaoXhotelMultiplerateGetAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelMultiplerateGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
