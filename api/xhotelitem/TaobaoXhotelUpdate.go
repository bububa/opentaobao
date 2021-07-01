package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
酒店更新接口（ID不存在自动新增） 
taobao.xhotel.update

酒店更新接口
*/
func TaobaoXhotelUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelUpdateAPIRequest, session string) (*xhotelitem.TaobaoXhotelUpdateAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
