package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
酒店查询接口 
taobao.xhotel.get

酒店查询接口
*/
func TaobaoXhotelGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelGetRequest, session string) (*xhotelitem.TaobaoXhotelGetAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
