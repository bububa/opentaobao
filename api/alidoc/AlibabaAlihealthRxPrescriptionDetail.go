package alidoc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alidoc"
)

/* 
处方详情 
alibaba.alihealth.rx.prescription.detail

获取处方结构化信息
*/
func AlibabaAlihealthRxPrescriptionDetail(clt *core.SDKClient, req *alidoc.AlibabaAlihealthRxPrescriptionDetailRequest, session string) (*alidoc.AlibabaAlihealthRxPrescriptionDetailAPIResponse, error) {
    var resp alidoc.AlibabaAlihealthRxPrescriptionDetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
