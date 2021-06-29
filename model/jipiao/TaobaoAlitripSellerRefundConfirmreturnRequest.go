package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】确认退票 API请求
taobao.alitrip.seller.refund.confirmreturn

确认退票
*/
type TaobaoAlitripSellerRefundConfirmreturnRequest struct {
    model.Params
    // 退票申请单
    applyId   int64
}

// 初始化TaobaoAlitripSellerRefundConfirmreturnRequest对象
func NewTaobaoAlitripSellerRefundConfirmreturnRequest() *TaobaoAlitripSellerRefundConfirmreturnRequest{
    return &TaobaoAlitripSellerRefundConfirmreturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundConfirmreturnRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refund.confirmreturn"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundConfirmreturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 退票申请单
func (r *TaobaoAlitripSellerRefundConfirmreturnRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r TaobaoAlitripSellerRefundConfirmreturnRequest) GetApplyId() int64 {
    return r.applyId
}
