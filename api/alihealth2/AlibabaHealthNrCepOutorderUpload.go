package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
线上订单收货验收单、出入库单据生成接口 
alibaba.health.nr.cep.outorder.upload

线上订单收货验收单、出入库单据生成接口
*/
func AlibabaHealthNrCepOutorderUpload(clt *core.SDKClient, req *alihealth2.AlibabaHealthNrCepOutorderUploadRequest, session string) (*alihealth2.AlibabaHealthNrCepOutorderUploadAPIResponse, error) {
    var resp alihealth2.AlibabaHealthNrCepOutorderUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
