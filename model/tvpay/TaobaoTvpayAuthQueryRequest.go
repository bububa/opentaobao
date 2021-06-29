package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付授权查询 API请求
taobao.tvpay.auth.query

查询该用户在指定设备上是否有支付授权
*/
type TaobaoTvpayAuthQueryRequest struct {
    model.Params
    // 设备号
    deviceId   string
    // 来源
    from   string
    // 业务订单号
    bizOrderId   string
    // 是否淘系
    isTao   bool
    // 支付宝订单号
    orderNo   string
    // 外部订单号
    outOrderNo   string
}

// 初始化TaobaoTvpayAuthQueryRequest对象
func NewTaobaoTvpayAuthQueryRequest() *TaobaoTvpayAuthQueryRequest{
    return &TaobaoTvpayAuthQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAuthQueryRequest) GetApiMethodName() string {
    return "taobao.tvpay.auth.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAuthQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备号
func (r *TaobaoTvpayAuthQueryRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayAuthQueryRequest) GetDeviceId() string {
    return r.deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayAuthQueryRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

// From Getter
func (r TaobaoTvpayAuthQueryRequest) GetFrom() string {
    return r.from
}
// BizOrderId Setter
// 业务订单号
func (r *TaobaoTvpayAuthQueryRequest) SetBizOrderId(bizOrderId string) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoTvpayAuthQueryRequest) GetBizOrderId() string {
    return r.bizOrderId
}
// IsTao Setter
// 是否淘系
func (r *TaobaoTvpayAuthQueryRequest) SetIsTao(isTao bool) error {
    r.isTao = isTao
    r.Set("is_tao", isTao)
    return nil
}

// IsTao Getter
func (r TaobaoTvpayAuthQueryRequest) GetIsTao() bool {
    return r.isTao
}
// OrderNo Setter
// 支付宝订单号
func (r *TaobaoTvpayAuthQueryRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

// OrderNo Getter
func (r TaobaoTvpayAuthQueryRequest) GetOrderNo() string {
    return r.orderNo
}
// OutOrderNo Setter
// 外部订单号
func (r *TaobaoTvpayAuthQueryRequest) SetOutOrderNo(outOrderNo string) error {
    r.outOrderNo = outOrderNo
    r.Set("out_order_no", outOrderNo)
    return nil
}

// OutOrderNo Getter
func (r TaobaoTvpayAuthQueryRequest) GetOutOrderNo() string {
    return r.outOrderNo
}
