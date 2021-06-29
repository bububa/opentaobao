package alihealthoutflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthoutflow"
)

/* 
异步开方处方详情 
alibaba.alihealth.asyncprescribe.prescription.detail

异步开方处方查询
*/
func AlibabaAlihealthAsyncprescribePrescriptionDetail(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthAsyncprescribePrescriptionDetailRequest, session string) (*alihealthoutflow.AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse, error) {
    var resp alihealthoutflow.AlibabaAlihealthAsyncprescribePrescriptionDetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
