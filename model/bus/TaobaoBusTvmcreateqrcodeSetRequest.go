package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自助机生成支付宝支付二维码 API请求
taobao.bus.tvmcreateqrcode.set

用于汽车票线下自助机调用获取支付宝的二维码
*/
type TaobaoBusTvmcreateqrcodeSetRequest struct {
    model.Params
    // 飞猪订单号
    alitripOrderId   string
    // 超时时间（分钟）
    timeoutExpress   int64
}

// 初始化TaobaoBusTvmcreateqrcodeSetRequest对象
func NewTaobaoBusTvmcreateqrcodeSetRequest() *TaobaoBusTvmcreateqrcodeSetRequest{
    return &TaobaoBusTvmcreateqrcodeSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmcreateqrcodeSetRequest) GetApiMethodName() string {
    return "taobao.bus.tvmcreateqrcode.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmcreateqrcodeSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusTvmcreateqrcodeSetRequest) SetAlitripOrderId(alitripOrderId string) error {
    r.alitripOrderId = alitripOrderId
    r.Set("alitrip_order_id", alitripOrderId)
    return nil
}

// AlitripOrderId Getter
func (r TaobaoBusTvmcreateqrcodeSetRequest) GetAlitripOrderId() string {
    return r.alitripOrderId
}
// TimeoutExpress Setter
// 超时时间（分钟）
func (r *TaobaoBusTvmcreateqrcodeSetRequest) SetTimeoutExpress(timeoutExpress int64) error {
    r.timeoutExpress = timeoutExpress
    r.Set("timeout_express", timeoutExpress)
    return nil
}

// TimeoutExpress Getter
func (r TaobaoBusTvmcreateqrcodeSetRequest) GetTimeoutExpress() int64 {
    return r.timeoutExpress
}
