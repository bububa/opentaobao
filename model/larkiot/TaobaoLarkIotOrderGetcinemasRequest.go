package larkiot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取iot渠道开放的影院 API请求
taobao.lark.iot.order.getcinemas

iot渠道拉取有权限访问的影院
*/
type TaobaoLarkIotOrderGetcinemasRequest struct {
    model.Params
    // 渠道编码
    channelCode   string
}

// 初始化TaobaoLarkIotOrderGetcinemasRequest对象
func NewTaobaoLarkIotOrderGetcinemasRequest() *TaobaoLarkIotOrderGetcinemasRequest{
    return &TaobaoLarkIotOrderGetcinemasRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLarkIotOrderGetcinemasRequest) GetApiMethodName() string {
    return "taobao.lark.iot.order.getcinemas"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLarkIotOrderGetcinemasRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelCode Setter
// 渠道编码
func (r *TaobaoLarkIotOrderGetcinemasRequest) SetChannelCode(channelCode string) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

// ChannelCode Getter
func (r TaobaoLarkIotOrderGetcinemasRequest) GetChannelCode() string {
    return r.channelCode
}
