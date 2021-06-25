package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款打款 APIRequest
alibaba.wdk.reverse.refund

五道口退款打款开放能力接口
*/
type AlibabaWdkReverseRefundRequest struct {
    model.Params

    // 退款打款请求
    openRefundReqDTO   *OpenRefundReqDto 

}

func NewAlibabaWdkReverseRefundRequest() *AlibabaWdkReverseRefundRequest{
    return &AlibabaWdkReverseRefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkReverseRefundRequest) GetApiMethodName() string {
    return "alibaba.wdk.reverse.refund"
}

func (r AlibabaWdkReverseRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkReverseRefundRequest) SetOpenRefundReqDTO(openRefundReqDTO *OpenRefundReqDto) error {
    r.openRefundReqDTO = openRefundReqDTO
    r.Set("open_refund_req_d_t_o", openRefundReqDTO)
    return nil
}

func (r AlibabaWdkReverseRefundRequest) GetOpenRefundReqDTO() *OpenRefundReqDto {
    return r.openRefundReqDTO
}

