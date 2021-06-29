package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
混淆nick转openid API请求
taobao.top.openid.convert

混淆nick转openid，生成混淆nick必须与当前请求的isv匹配
*/
type TaobaoTopOpenidConvertRequest struct {
    model.Params
    // 混淆nick转open_id
    mixNick   string
}

// 初始化TaobaoTopOpenidConvertRequest对象
func NewTaobaoTopOpenidConvertRequest() *TaobaoTopOpenidConvertRequest{
    return &TaobaoTopOpenidConvertRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTopOpenidConvertRequest) GetApiMethodName() string {
    return "taobao.top.openid.convert"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTopOpenidConvertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MixNick Setter
// 混淆nick转open_id
func (r *TaobaoTopOpenidConvertRequest) SetMixNick(mixNick string) error {
    r.mixNick = mixNick
    r.Set("mix_nick", mixNick)
    return nil
}

// MixNick Getter
func (r TaobaoTopOpenidConvertRequest) GetMixNick() string {
    return r.mixNick
}
