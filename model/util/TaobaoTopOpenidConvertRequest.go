package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
混淆nick转openid APIRequest
taobao.top.openid.convert

混淆nick转openid，生成混淆nick必须与当前请求的isv匹配
*/
type TaobaoTopOpenidConvertRequest struct {
    model.Params

    // 混淆nick转open_id
    mixNick   string 

}

func NewTaobaoTopOpenidConvertRequest() *TaobaoTopOpenidConvertRequest{
    return &TaobaoTopOpenidConvertRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTopOpenidConvertRequest) GetApiMethodName() string {
    return "taobao.top.openid.convert"
}

func (r TaobaoTopOpenidConvertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTopOpenidConvertRequest) SetMixNick(mixNick string) error {
    r.mixNick = mixNick
    r.Set("mix_nick", mixNick)
    return nil
}

func (r TaobaoTopOpenidConvertRequest) GetMixNick() string {
    return r.mixNick
}

