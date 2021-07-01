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
type TaobaoTvpayAuthQueryAPIRequest struct {
    model.Params
    // 设备号
    _deviceId   string
    // 来源
    _from   string
    // 业务订单号
    _bizOrderId   string
    // 是否淘系
    _isTao   bool
    // 支付宝订单号
    _orderNo   string
    // 外部订单号
    _outOrderNo   string
}

// 初始化TaobaoTvpayAuthQueryAPIRequest对象
func NewTaobaoTvpayAuthQueryRequest() *TaobaoTvpayAuthQueryAPIRequest{
    return &TaobaoTvpayAuthQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAuthQueryAPIRequest) GetApiMethodName() string {
    return "taobao.tvpay.auth.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAuthQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备号
func (r *TaobaoTvpayAuthQueryAPIRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetDeviceId() string {
    return r._deviceId
}
// From Setter
// 来源
func (r *TaobaoTvpayAuthQueryAPIRequest) SetFrom(_from string) error {
    r._from = _from
    r.Set("from", _from)
    return nil
}

// From Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetFrom() string {
    return r._from
}
// BizOrderId Setter
// 业务订单号
func (r *TaobaoTvpayAuthQueryAPIRequest) SetBizOrderId(_bizOrderId string) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetBizOrderId() string {
    return r._bizOrderId
}
// IsTao Setter
// 是否淘系
func (r *TaobaoTvpayAuthQueryAPIRequest) SetIsTao(_isTao bool) error {
    r._isTao = _isTao
    r.Set("is_tao", _isTao)
    return nil
}

// IsTao Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetIsTao() bool {
    return r._isTao
}
// OrderNo Setter
// 支付宝订单号
func (r *TaobaoTvpayAuthQueryAPIRequest) SetOrderNo(_orderNo string) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetOrderNo() string {
    return r._orderNo
}
// OutOrderNo Setter
// 外部订单号
func (r *TaobaoTvpayAuthQueryAPIRequest) SetOutOrderNo(_outOrderNo string) error {
    r._outOrderNo = _outOrderNo
    r.Set("out_order_no", _outOrderNo)
    return nil
}

// OutOrderNo Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetOutOrderNo() string {
    return r._outOrderNo
}
