package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
民宿房东信息添加 
taobao.xhotel.bnbowner.add

添加和更新民宿房东的信息
*/
func TaobaoXhotelBnbownerAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbownerAddAPIRequest, session string) (*xhotelitem.TaobaoXhotelBnbownerAddAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelBnbownerAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
