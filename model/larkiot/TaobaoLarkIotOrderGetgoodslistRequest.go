package larkiot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
iot渠道获取卖品信息 APIRequest
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

func NewTaobaoLarkIotOrderGetgoodslistRequest() *TaobaoLarkIotOrderGetgoodslistRequest{
    return &TaobaoLarkIotOrderGetgoodslistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLarkIotOrderGetgoodslistRequest) GetApiMethodName() string {
    return "taobao.lark.iot.order.getgoodslist"
}

func (r TaobaoLarkIotOrderGetgoodslistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLarkIotOrderGetgoodslistRequest) SetChannelCode(channelCode string) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

func (r TaobaoLarkIotOrderGetgoodslistRequest) GetChannelCode() string {
    return r.channelCode
}

func (r *TaobaoLarkIotOrderGetgoodslistRequest) SetCinemaLinkId(cinemaLinkId string) error {
    r.cinemaLinkId = cinemaLinkId
    r.Set("cinema_link_id", cinemaLinkId)
    return nil
}

func (r TaobaoLarkIotOrderGetgoodslistRequest) GetCinemaLinkId() string {
    return r.cinemaLinkId
}

