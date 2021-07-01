package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
酒店新增接口（ID重复自动更新） 
taobao.xhotel.add

添加酒店或更新酒店
*/
func TaobaoXhotelAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelAddAPIRequest, session string) (*xhotelitem.TaobaoXhotelAddAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
