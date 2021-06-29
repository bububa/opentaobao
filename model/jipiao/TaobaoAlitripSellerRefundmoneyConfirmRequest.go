package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】确认退款 API请求
taobao.alitrip.seller.refundmoney.confirm

代理人确认退票申请单的退款
*/
type TaobaoAlitripSellerRefundmoneyConfirmRequest struct {
    model.Params
    // 申请单id
    applyId   int64
}

// 初始化TaobaoAlitripSellerRefundmoneyConfirmRequest对象
func NewTaobaoAlitripSellerRefundmoneyConfirmRequest() *TaobaoAlitripSellerRefundmoneyConfirmRequest{
    return &TaobaoAlitripSellerRefundmoneyConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundmoneyConfirmRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refundmoney.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundmoneyConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 申请单id
func (r *TaobaoAlitripSellerRefundmoneyConfirmRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripSellerRefundmoneyConfirmRequest) GetApplyId() int64 {
    return r.applyId
}
