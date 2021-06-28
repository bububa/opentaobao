package blackvip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/blackvip"
)

/* 
88VIP用户信息查询 
taobao.blackvip.userinfo.get

查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等
*/
func TaobaoBlackvipUserinfoGet(clt *core.SDKClient, req *blackvip.TaobaoBlackvipUserinfoGetRequest, session string) (*blackvip.TaobaoBlackvipUserinfoGetAPIResponse, error) {
    var resp blackvip.TaobaoBlackvipUserinfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
