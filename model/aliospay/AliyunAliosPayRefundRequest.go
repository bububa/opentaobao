package aliospay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款接口 APIRequest
aliyun.alios.pay.refund

商户用来发起退款的接口
*/
type AliyunAliosPayRefundRequest struct {
    model.Params

    // 请求参数
    refundRequest   *RefundRequest 

}

func NewAliyunAliosPayRefundRequest() *AliyunAliosPayRefundRequest{
    return &AliyunAliosPayRefundRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunAliosPayRefundRequest) GetApiMethodName() string {
    return "aliyun.alios.pay.refund"
}

func (r AliyunAliosPayRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunAliosPayRefundRequest) SetRefundRequest(refundRequest *RefundRequest) error {
    r.refundRequest = refundRequest
    r.Set("refund_request", refundRequest)
    return nil
}

func (r AliyunAliosPayRefundRequest) GetRefundRequest() *RefundRequest {
    return r.refundRequest
}

