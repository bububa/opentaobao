package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付查询订单状态 API请求
taobao.tvpay.order.query

tv支付查询订单状态
*/
type TaobaoTvpayOrderQueryRequest struct {
    model.Params
    // 设备id
    _deviceId   string
    // 来源
    _from   string
    // 业务订单号
    _bizOrderId   string
    // 是否淘系
    _isTao   bool
    // 支付宝订单号
    _orderNo   string
    // 订单类型
    _orderType   string
    // 外部订单号
    _outOrderNo   string
    // 牌照方
    _license   string
}

// 初始化TaobaoTvpayOrderQueryRequest对象
func NewTaobaoTvpayOrderQueryRequest() *TaobaoTvpayOrderQueryRequest{
    return &TaobaoTvpayOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayOrderQueryRequest) GetApiMethodName() string {
    return "taobao.tvpay.order.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备id
func (r *TaobaoTvpayOrderQueryRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayOrderQueryRequest) GetDeviceId() string {
    return r._deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayOrderQueryRequest) SetFrom(_from string) error {
    r._from = _from
    r.Set("from", _from)
    return nil
}

// From Getter
func (r TaobaoTvpayOrderQueryRequest) GetFrom() string {
    return r._from
}
// BizOrderId Setter
// 业务订单号
func (r *TaobaoTvpayOrderQueryRequest) SetBizOrderId(_bizOrderId string) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoTvpayOrderQueryRequest) GetBizOrderId() string {
    return r._bizOrderId
}
// IsTao Setter
// 是否淘系
func (r *TaobaoTvpayOrderQueryRequest) SetIsTao(_isTao bool) error {
    r._isTao = _isTao
    r.Set("is_tao", _isTao)
    return nil
}

// IsTao Getter
func (r TaobaoTvpayOrderQueryRequest) GetIsTao() bool {
    return r._isTao
}
// OrderNo Setter
// 支付宝订单号
func (r *TaobaoTvpayOrderQueryRequest) SetOrderNo(_orderNo string) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r TaobaoTvpayOrderQueryRequest) GetOrderNo() string {
    return r._orderNo
}
// OrderType Setter
// 订单类型
func (r *TaobaoTvpayOrderQueryRequest) SetOrderType(_orderType string) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r TaobaoTvpayOrderQueryRequest) GetOrderType() string {
    return r._orderType
}
// OutOrderNo Setter
// 外部订单号
func (r *TaobaoTvpayOrderQueryRequest) SetOutOrderNo(_outOrderNo string) error {
    r._outOrderNo = _outOrderNo
    r.Set("out_order_no", _outOrderNo)
    return nil
}

// OutOrderNo Getter
func (r TaobaoTvpayOrderQueryRequest) GetOutOrderNo() string {
    return r._outOrderNo
}
// License Setter
// 牌照方
func (r *TaobaoTvpayOrderQueryRequest) SetLicense(_license string) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r TaobaoTvpayOrderQueryRequest) GetLicense() string {
    return r._license
}
