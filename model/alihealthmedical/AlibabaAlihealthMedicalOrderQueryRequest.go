package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方机构查询订单详情接口 APIRequest
alibaba.alihealth.medical.order.query

查询订单详情，包括评价
*/
type AlibabaAlihealthMedicalOrderQueryRequest struct {
    model.Params

    // 请求入参
    requestInfo   *OrderQueryRequestDTO 

}

func NewAlibabaAlihealthMedicalOrderQueryRequest() *AlibabaAlihealthMedicalOrderQueryRequest{
    return &AlibabaAlihealthMedicalOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMedicalOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.order.query"
}

func (r AlibabaAlihealthMedicalOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMedicalOrderQueryRequest) SetRequestInfo(requestInfo *OrderQueryRequestDTO) error {
    r.requestInfo = requestInfo
    r.Set("request_info", requestInfo)
    return nil
}

func (r AlibabaAlihealthMedicalOrderQueryRequest) GetRequestInfo() *OrderQueryRequestDTO {
    return r.requestInfo
}

