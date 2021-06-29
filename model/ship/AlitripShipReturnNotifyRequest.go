package ship

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
船票退票退款回填接口 API请求
alitrip.ship.return.notify

此接口为接入商调用飞猪接口回填退票状态，飞猪平台给用户进行退票退款。飞猪平台保证数据幂等。
*/
type AlitripShipReturnNotifyRequest struct {
    model.Params
    // 退票请求入参
    confirmRefundRQ   *ShipAgentConfirmRefundRq
}

// 初始化AlitripShipReturnNotifyRequest对象
func NewAlitripShipReturnNotifyRequest() *AlitripShipReturnNotifyRequest{
    return &AlitripShipReturnNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripShipReturnNotifyRequest) GetApiMethodName() string {
    return "alitrip.ship.return.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripShipReturnNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConfirmRefundRQ Setter
// 退票请求入参
func (r *AlitripShipReturnNotifyRequest) SetConfirmRefundRQ(confirmRefundRQ *ShipAgentConfirmRefundRq) error {
    r.confirmRefundRQ = confirmRefundRQ
    r.Set("confirm_refund_r_q", confirmRefundRQ)
    return nil
}

// ConfirmRefundRQ Getter
func (r AlitripShipReturnNotifyRequest) GetConfirmRefundRQ() *ShipAgentConfirmRefundRq {
    return r.confirmRefundRQ
}
