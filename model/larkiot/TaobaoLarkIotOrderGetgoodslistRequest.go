package larkiot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
iot渠道获取卖品信息 API请求
taobao.lark.iot.order.getgoodslist

iot无人超市服务商通过接口获取影院的可售卖品数据
*/
type TaobaoLarkIotOrderGetgoodslistRequest struct {
    model.Params
    // 渠道编码
    channelCode   string
    // 影院内码
    cinemaLinkId   string
}

// 初始化TaobaoLarkIotOrderGetgoodslistRequest对象
func NewTaobaoLarkIotOrderGetgoodslistRequest() *TaobaoLarkIotOrderGetgoodslistRequest{
    return &TaobaoLarkIotOrderGetgoodslistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLarkIotOrderGetgoodslistRequest) GetApiMethodName() string {
    return "taobao.lark.iot.order.getgoodslist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLarkIotOrderGetgoodslistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelCode Setter
// 渠道编码
func (r *TaobaoLarkIotOrderGetgoodslistRequest) SetChannelCode(channelCode string) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

// ChannelCode Getter
func (r TaobaoLarkIotOrderGetgoodslistRequest) GetChannelCode() string {
    return r.channelCode
}
// CinemaLinkId Setter
// 影院内码
func (r *TaobaoLarkIotOrderGetgoodslistRequest) SetCinemaLinkId(cinemaLinkId string) error {
    r.cinemaLinkId = cinemaLinkId
    r.Set("cinema_link_id", cinemaLinkId)
    return nil
}

// CinemaLinkId Getter
func (r TaobaoLarkIotOrderGetgoodslistRequest) GetCinemaLinkId() string {
    return r.cinemaLinkId
}
