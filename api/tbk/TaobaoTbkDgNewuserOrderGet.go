package tbk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tbk"
)

/* 
淘宝客-推广者-新用户订单明细查询 
taobao.tbk.dg.newuser.order.get

拉新API
*/
func TaobaoTbkDgNewuserOrderGet(clt *core.SDKClient, req *tbk.TaobaoTbkDgNewuserOrderGetRequest, session string) (*tbk.TaobaoTbkDgNewuserOrderGetResponse, error) {
    var resp tbk.TaobaoTbkDgNewuserOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
