package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
复杂价格删除 
taobao.xhotel.multiplerate.delete

酒店产品库rate删除
*/
func TaobaoXhotelMultiplerateDelete(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelMultiplerateDeleteRequest, session string) (*xhotelitem.TaobaoXhotelMultiplerateDeleteAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelMultiplerateDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
