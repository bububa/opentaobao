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
type TaobaoAlitripSellerRefundRefusereturnRequest struct {
    model.Params
    // 申请单ID
    applyId   int64
    // 拒绝理由
    reason   string
}

// 初始化TaobaoAlitripSellerRefundRefusereturnRequest对象
func NewTaobaoAlitripSellerRefundRefusereturnRequest() *TaobaoAlitripSellerRefundRefusereturnRequest{
    return &TaobaoAlitripSellerRefundRefusereturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundRefusereturnRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refund.refusereturn"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundRefusereturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 申请单ID
func (r *TaobaoAlitripSellerRefundRefusereturnRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripSellerRefundRefusereturnRequest) GetApplyId() int64 {
    return r.applyId
}
// Reason Setter
// 拒绝理由
func (r *TaobaoAlitripSellerRefundRefusereturnRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

// Reason Getter
func (r TaobaoAlitripSellerRefundRefusereturnRequest) GetReason() string {
    return r.reason
}
