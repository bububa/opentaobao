package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
医院上传出库信息 
alibaba.alihealth.drug.kyt.yy.uploadretail

医院上传出库信息接口
*/
func AlibabaAlihealthDrugKytYyUploadretail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYyUploadretailRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytYyUploadretailAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytYyUploadretailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}