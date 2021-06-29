package elife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌惠卡券冲正退还 API请求
taobao.elife.lifecard.refund

淘宝生活汇消费卡虚拟卡，线下冲正退货接口
*/
type TaobaoElifeLifecardRefundRequest struct {
    model.Params
    // 请求参数
    refundRequest   *RefundRequest
}

// 初始化TaobaoElifeLifecardRefundRequest对象
func NewTaobaoElifeLifecardRefundRequest() *TaobaoElifeLifecardRefundRequest{
    return &TaobaoElifeLifecardRefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoElifeLifecardRefundRequest) GetApiMethodName() string {
    return "taobao.elife.lifecard.refund"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoElifeLifecardRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundRequest Setter
// 请求参数
func (r *TaobaoElifeLifecardRefundRequest) SetRefundRequest(refundRequest *RefundRequest) error {
    r.refundRequest = refundRequest
    r.Set("refund_request", refundRequest)
    return nil
}

// RefundRequest Getter
func (r TaobaoElifeLifecardRefundRequest) GetRefundRequest() *RefundRequest {
    return r.refundRequest
}
