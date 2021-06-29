package ship

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单信息回填(出票回调) API请求
alitrip.ship.order.notify

此接口为接入商调用飞猪旅行接口回填票号、密码(验证码)等订单信息。接口根据alitripOrderId幂等。若第一次调用失败，后续调用仍然可以回填票号、密码(验证码)成功。第一次调用成功后，后续调用会直接返回第一次的调用结果，不会再产生更新操作。多张票同时出票回填时，保证原子性，只允许全部成功或者全部失败，不能存在部分成功或者失败
*/
type AlitripShipOrderNotifyRequest struct {
    model.Params
    // 出票入参
    _confirmBookRQ   *ShipAgentConfirmBookRq
}

// 初始化AlitripShipOrderNotifyRequest对象
func NewAlitripShipOrderNotifyRequest() *AlitripShipOrderNotifyRequest{
    return &AlitripShipOrderNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripShipOrderNotifyRequest) GetApiMethodName() string {
    return "alitrip.ship.order.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripShipOrderNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConfirmBookRQ Setter
// 出票入参
func (r *AlitripShipOrderNotifyRequest) SetConfirmBookRQ(_confirmBookRQ *ShipAgentConfirmBookRq) error {
    r._confirmBookRQ = _confirmBookRQ
    r.Set("confirm_book_r_q", _confirmBookRQ)
    return nil
}

// ConfirmBookRQ Getter
func (r AlitripShipOrderNotifyRequest) GetConfirmBookRQ() *ShipAgentConfirmBookRq {
    return r._confirmBookRQ
}
