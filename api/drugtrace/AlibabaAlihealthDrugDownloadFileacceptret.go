package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
企业上传回执 
alibaba.alihealth.drug.download.fileacceptret

拿到企业下载回执，将企业已下载的和未下载成功的条目都相应的改变状态
*/
func AlibabaAlihealthDrugDownloadFileacceptret(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadFileacceptretAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugDownloadFileacceptretAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugDownloadFileacceptretAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
