package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
民宿新增房源 
taobao.xhotel.bnbroomtype.add

添加民宿房源
*/
func TaobaoXhotelBnbroomtypeAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbroomtypeAddRequest, session string) (*xhotelitem.TaobaoXhotelBnbroomtypeAddAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelBnbroomtypeAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
