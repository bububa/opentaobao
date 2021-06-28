package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
根据订单查询订单标签 
taobao.oc.tradetags.get

根据订单查询订单标签。<br/>
返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。<br/>
官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明<br/>
主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865<br/>
*/
func TaobaoOcTradetagsGet(clt *core.SDKClient, req *jst.TaobaoOcTradetagsGetRequest, session string) (*jst.TaobaoOcTradetagsGetAPIResponse, error) {
    var resp jst.TaobaoOcTradetagsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
