package aliospay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款接口 API请求
aliyun.alios.pay.refund

商户用来发起退款的接口
*/
type AliyunAliosPayRefundRequest struct {
    model.Params
    // 请求参数
    _refundRequest   *RefundRequest
}

// 初始化AliyunAliosPayRefundRequest对象
func NewAliyunAliosPayRefundRequest() *AliyunAliosPayRefundRequest{
    return &AliyunAliosPayRefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunAliosPayRefundRequest) GetApiMethodName() string {
    return "aliyun.alios.pay.refund"
}

// IRequest interface 方法, 获取API参数
func (r AliyunAliosPayRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundRequest Setter
// 请求参数
func (r *AliyunAliosPayRefundRequest) SetRefundRequest(_refundRequest *RefundRequest) error {
    r._refundRequest = _refundRequest
    r.Set("refund_request", _refundRequest)
    return nil
}

// RefundRequest Getter
func (r AliyunAliosPayRefundRequest) GetRefundRequest() *RefundRequest {
    return r._refundRequest
}
