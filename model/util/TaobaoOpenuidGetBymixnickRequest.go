package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过mixnick转换openuid API请求
taobao.openuid.get.bymixnick

通过mixnick转换openuid
*/
type TaobaoOpenuidGetBymixnickRequest struct {
    model.Params
    // 无线类应用获取到的混淆的nick
    mixNick   string
}

// 初始化TaobaoOpenuidGetBymixnickRequest对象
func NewTaobaoOpenuidGetBymixnickRequest() *TaobaoOpenuidGetBymixnickRequest{
    return &TaobaoOpenuidGetBymixnickRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenuidGetBymixnickRequest) GetApiMethodName() string {
    return "taobao.openuid.get.bymixnick"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenuidGetBymixnickRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MixNick Setter
// 无线类应用获取到的混淆的nick
func (r *TaobaoOpenuidGetBymixnickRequest) SetMixNick(mixNick string) error {
    r.mixNick = mixNick
    r.Set("mix_nick", mixNick)
    return nil
}

// MixNick Getter
func (r TaobaoOpenuidGetBymixnickRequest) GetMixNick() string {
    return r.mixNick
}
