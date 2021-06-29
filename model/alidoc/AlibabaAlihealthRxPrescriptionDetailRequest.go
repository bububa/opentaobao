package alidoc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方详情 API请求
alibaba.alihealth.rx.prescription.detail

获取处方结构化信息
*/
type AlibabaAlihealthRxPrescriptionDetailRequest struct {
    model.Params
    // 查询参数
    query   *RxPrescriptionQuery
}

// 初始化AlibabaAlihealthRxPrescriptionDetailRequest对象
func NewAlibabaAlihealthRxPrescriptionDetailRequest() *AlibabaAlihealthRxPrescriptionDetailRequest{
    return &AlibabaAlihealthRxPrescriptionDetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthRxPrescriptionDetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.rx.prescription.detail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthRxPrescriptionDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询参数
func (r *AlibabaAlihealthRxPrescriptionDetailRequest) SetQuery(query *RxPrescriptionQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaAlihealthRxPrescriptionDetailRequest) GetQuery() *RxPrescriptionQuery {
    return r.query
}
