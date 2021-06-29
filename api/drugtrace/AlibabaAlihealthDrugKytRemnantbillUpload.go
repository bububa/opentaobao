package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
零头出入库单据上传 
alibaba.alihealth.drug.kyt.remnantbill.upload

零头出入库单据上传
*/
func AlibabaAlihealthDrugKytRemnantbillUpload(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytRemnantbillUploadRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
