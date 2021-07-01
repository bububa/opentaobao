package shenjing

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shenjing"
)

/* 
访客PAD上传人脸 
alibaba.ib.shenjing.visitor.pad.uploadface

访客PAD端上传人脸。
*/
func AlibabaIbShenjingVisitorPadUploadface(clt *core.SDKClient, req *shenjing.AlibabaIbShenjingVisitorPadUploadfaceAPIRequest, session string) (*shenjing.AlibabaIbShenjingVisitorPadUploadfaceAPIResponse, error) {
    var resp shenjing.AlibabaIbShenjingVisitorPadUploadfaceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
