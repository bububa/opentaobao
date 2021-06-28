package iot

import (
    "encoding/xml"

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
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_device_trade_precreate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 链接有效结束时间
    
    LifeEnd   string `json:"life_end,omitempty" xml:"