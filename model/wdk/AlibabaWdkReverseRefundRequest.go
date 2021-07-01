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
type AlibabaWdkReverseRefundAPIRequest struct {
    model.Params
    // 退款打款请求
    _openRefundReqDTO   *OpenRefundReqDTO
}

// 初始化AlibabaWdkReverseRefundAPIRequest对象
func NewAlibabaWdkReverseRefundRequest() *AlibabaWdkReverseRefundAPIRequest{
    return &AlibabaWdkReverseRefundAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseRefundAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.reverse.refund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseRefundAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenRefundReqDTO Setter
// 退款打款请求
func (r *AlibabaWdkReverseRefundAPIRequest) SetOpenRefundReqDTO(_openRefundReqDTO *OpenRefundReqDTO) error {
    r._openRefundReqDTO = _openRefundReqDTO
    r.Set("open_refund_req_d_t_o", _openRefundReqDTO)
    return nil
}

// OpenRefundReqDTO Getter
func (r AlibabaWdkReverseRefundAPIRequest) GetOpenRefundReqDTO() *OpenRefundReqDTO {
    return r._openRefundReqDTO
}
