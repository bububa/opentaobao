package blackvip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
88VIP用户信息查询 API请求
taobao.blackvip.userinfo.get

查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等
*/
type TaobaoBlackvipUserinfoGetRequest struct {
    model.Params
}

// 初始化TaobaoBlackvipUserinfoGetRequest对象
func NewTaobaoBlackvipUserinfoGetRequest() *TaobaoBlackvipUserinfoGetRequest{
    return &TaobaoBlackvipUserinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBlackvipUserinfoGetRequest) GetApiMethodName() string {
    return "taobao.blackvip.userinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBlackvipUserinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
