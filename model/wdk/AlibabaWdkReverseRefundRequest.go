package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款打款 API请求
alibaba.wdk.reverse.refund

五道口退款打款开放能力接口
*/
type AlibabaWdkReverseRefundRequest struct {
    model.Params
    // 退款打款请求
    _openRefundReqDTO   *OpenRefundReqDTO
}

// 初始化AlibabaWdkReverseRefundRequest对象
func NewAlibabaWdkReverseRefundRequest() *AlibabaWdkReverseRefundRequest{
    return &AlibabaWdkReverseRefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseRefundRequest) GetApiMethodName() string {
    return "alibaba.wdk.reverse.refund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenRefundReqDTO Setter
// 退款打款请求
func (r *AlibabaWdkReverseRefundRequest) SetOpenRefundReqDTO(_openRefundReqDTO *OpenRefundReqDTO) error {
    r._openRefundReqDTO = _openRefundReqDTO
    r.Set("open_refund_req_d_t_o", _openRefundReqDTO)
    return nil
}

// OpenRefundReqDTO Getter
func (r AlibabaWdkReverseRefundRequest) GetOpenRefundReqDTO() *OpenRefundReqDTO {
    return r._openRefundReqDTO
}
