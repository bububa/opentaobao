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
type TaobaoElifeLifecardRefundAPIRequest struct {
    model.Params
    // 请求参数
    _refundRequest   *RefundRequest
}

// 初始化TaobaoElifeLifecardRefundAPIRequest对象
func NewTaobaoElifeLifecardRefundRequest() *TaobaoElifeLifecardRefundAPIRequest{
    return &TaobaoElifeLifecardRefundAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoElifeLifecardRefundAPIRequest) GetApiMethodName() string {
    return "taobao.elife.lifecard.refund"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoElifeLifecardRefundAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundRequest Setter
// 请求参数
func (r *TaobaoElifeLifecardRefundAPIRequest) SetRefundRequest(_refundRequest *RefundRequest) error {
    r._refundRequest = _refundRequest
    r.Set("refund_request", _refundRequest)
    return nil
}

// RefundRequest Getter
func (r TaobaoElifeLifecardRefundAPIRequest) GetRefundRequest() *RefundRequest {
    return r._refundRequest
}
