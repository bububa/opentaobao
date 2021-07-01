package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】拒绝退票 API请求
taobao.alitrip.seller.refund.refusereturn

拒绝退票
*/
type TaobaoAlitripSellerRefundRefusereturnAPIRequest struct {
    model.Params
    // 申请单ID
    _applyId   int64
    // 拒绝理由
    _reason   string
}

// 初始化TaobaoAlitripSellerRefundRefusereturnAPIRequest对象
func NewTaobaoAlitripSellerRefundRefusereturnRequest() *TaobaoAlitripSellerRefundRefusereturnAPIRequest{
    return &TaobaoAlitripSellerRefundRefusereturnAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundRefusereturnAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refund.refusereturn"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundRefusereturnAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 申请单ID
func (r *TaobaoAlitripSellerRefundRefusereturnAPIRequest) SetApplyId(_applyId int64) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripSellerRefundRefusereturnAPIRequest) GetApplyId() int64 {
    return r._applyId
}
// Reason Setter
// 拒绝理由
func (r *TaobaoAlitripSellerRefundRefusereturnAPIRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r TaobaoAlitripSellerRefundRefusereturnAPIRequest) GetReason() string {
    return r._reason
}
