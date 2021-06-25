package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件上预创建天猫订单 APIResponse
tmall.device.trade.precreate

智能硬件上预创建天猫订单。
1，use_open_price不再需要传入，使用unit_price传入单价。
2，订单默认5分钟自动关闭，没有付款的订单在手机淘宝不可见。
3，同一个码只运行一个用户扫码，多个用户扫一个码会报错 订单不存在。
*/
type TmallDeviceTradePrecreateAPIResponse struct {
    model.CommonResponse
    Response *TmallDeviceTradePrecreateResponse `json:"tmall_device_trade_precreate_response,omitempty"`
}

type TmallDeviceTradePrecreateResponse struct {

    // 链接有效结束时间
    LifeEnd   string `json:"life_end,omitempty"`

    // 链接有效起始时间
    LifeStart   string `json:"life_start,omitempty"`

    // 链接二维码图片
    QrCode   string `json:"qr_code,omitempty"`

    // 短链接
    ShortUrl   string `json:"short_url,omitempty"`

}
