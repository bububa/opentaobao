package larkiot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取iot渠道开放的影院 APIRequest
taobao.lark.iot.order.getcinemas

iot渠道拉取有权限访问的影院
*/
type TaobaoLarkIotOrderGetcinemasRequest struct {
    model.Params

    // 渠道编码
    channelCode   string 

}

func NewTaobaoLarkIotOrderGetcinemasRequest() *TaobaoLarkIotOrderGetcinemasRequest{
    return &TaobaoLarkIotOrderGetcinemasRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLarkIotOrderGetcinemasRequest) GetApiMethodName() string {
    return "taobao.lark.iot.order.getcinemas"
}

func (r TaobaoLarkIotOrderGetcinemasRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLarkIotOrderGetcinemasRequest) SetChannelCode(channelCode string) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

func (r TaobaoLarkIotOrderGetcinemasRequest) GetChannelCode() string {
    return r.channelCode
}

