package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
零售单据上传接口 
alibaba.alihealth.drugtrace.top.lsyd.uploadretail

快易通多融零售上传接口
*/
func AlibabaAlihealthDrugtraceTopLsydUploadretail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydUploadretailAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
