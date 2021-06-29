package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方机构查询订单详情接口 API请求
alibaba.alihealth.medical.order.query

查询订单详情，包括评价
*/
type AlibabaAlihealthMedicalOrderQueryRequest struct {
    model.Params
    // 请求入参
    _requestInfo   *OrderQueryRequestDTO
}

// 初始化AlibabaAlihealthMedicalOrderQueryRequest对象
func NewAlibabaAlihealthMedicalOrderQueryRequest() *AlibabaAlihealthMedicalOrderQueryRequest{
    return &AlibabaAlihealthMedicalOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestInfo Setter
// 请求入参
func (r *AlibabaAlihealthMedicalOrderQueryRequest) SetRequestInfo(_requestInfo *OrderQueryRequestDTO) error {
    r._requestInfo = _requestInfo
    r.Set("request_info", _requestInfo)
    return nil
}

// RequestInfo Getter
func (r AlibabaAlihealthMedicalOrderQueryRequest) GetRequestInfo() *OrderQueryRequestDTO {
    return r._requestInfo
}
