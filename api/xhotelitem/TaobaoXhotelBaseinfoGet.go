package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
酒店基础信息查询接口 
taobao.xhotel.baseinfo.get

酒店基础信息(酒店/房型/房价定义)查询接口， 包括 酒店房型可售, 以及 hid 下 的标准房型列表
*/
func TaobaoXhotelBaseinfoGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBaseinfoGetRequest, session string) (*xhotelitem.TaobaoXhotelBaseinfoGetAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelBaseinfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
