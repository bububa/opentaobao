package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
企业下载列表 
alibaba.alihealth.drug.download.entlist

获取企业的下载文件列表
*/
func AlibabaAlihealthDrugDownloadEntlist(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadEntlistAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugDownloadEntlistAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugDownloadEntlistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
