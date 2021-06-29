package alidoc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方详情 APIRequest
alibaba.alihealth.rx.prescription.detail

获取处方结构化信息
*/
type AlibabaAlihealthRxPrescriptionDetailRequest struct {
    model.Params

    // 查询参数
    query   *RxPrescriptionQuery 

}

func NewAlibabaAlihealthRxPrescriptionDetailRequest() *AlibabaAlihealthRxPrescriptionDetailRequest{
    return &AlibabaAlihealthRxPrescriptionDetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthRxPrescriptionDetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.rx.prescription.detail"
}

func (r AlibabaAlihealthRxPrescriptionDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthRxPrescriptionDetailRequest) SetQuery(query *RxPrescriptionQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaAlihealthRxPrescriptionDetailRequest) GetQuery() *RxPrescriptionQuery {
    return r.query
}

