package icburfq

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icburfq"
)

/* 
上传附件获取附件files_str 
alibaba.icbu.annex.upload

上传附件获取附件files_str
*/
func AlibabaIcbuAnnexUpload(clt *core.SDKClient, req *icburfq.AlibabaIcbuAnnexUploadAPIRequest, session string) (*icburfq.AlibabaIcbuAnnexUploadAPIResponse, error) {
    var resp icburfq.AlibabaIcbuAnnexUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
