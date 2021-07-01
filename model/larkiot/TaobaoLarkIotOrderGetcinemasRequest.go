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
type TaobaoLarkIotOrderGetcinemasAPIRequest struct {
    model.Params
    // 渠道编码
    _channelCode   string
}

// 初始化TaobaoLarkIotOrderGetcinemasAPIRequest对象
func NewTaobaoLarkIotOrderGetcinemasRequest() *TaobaoLarkIotOrderGetcinemasAPIRequest{
    return &TaobaoLarkIotOrderGetcinemasAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLarkIotOrderGetcinemasAPIRequest) GetApiMethodName() string {
    return "taobao.lark.iot.order.getcinemas"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLarkIotOrderGetcinemasAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelCode Setter
// 渠道编码
func (r *TaobaoLarkIotOrderGetcinemasAPIRequest) SetChannelCode(_channelCode string) error {
    r._channelCode = _channelCode
    r.Set("channel_code", _channelCode)
    return nil
}

// ChannelCode Getter
func (r TaobaoLarkIotOrderGetcinemasAPIRequest) GetChannelCode() string {
    return r._channelCode
}
