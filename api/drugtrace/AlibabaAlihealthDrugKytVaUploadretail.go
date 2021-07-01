package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
零售单据上传接口 
alibaba.alihealth.drug.kyt.va.uploadretail

零售上传单据信息接口
*/
func AlibabaAlihealthDrugKytVaUploadretail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytVaUploadretailAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytVaUploadretailAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytVaUploadretailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
