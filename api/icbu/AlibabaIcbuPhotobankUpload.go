package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
图片银行图片上传开放接口 
alibaba.icbu.photobank.upload

图片银行图片上传开放接口
*/
func AlibabaIcbuPhotobankUpload(clt *core.SDKClient, req *icbu.AlibabaIcbuPhotobankUploadRequest, session string) (*icbu.AlibabaIcbuPhotobankUploadResponse, error) {
    var resp icbu.AlibabaIcbuPhotobankUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
