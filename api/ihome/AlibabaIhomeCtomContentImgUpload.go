package ihome

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ihome"
)

/* 
实拍图投稿链路上传图片 
alibaba.ihome.ctom.content.img.upload

实拍图投稿链路上传图片
*/
func AlibabaIhomeCtomContentImgUpload(clt *core.SDKClient, req *ihome.AlibabaIhomeCtomContentImgUploadRequest, session string) (*ihome.AlibabaIhomeCtomContentImgUploadAPIResponse, error) {
    var resp ihome.AlibabaIhomeCtomContentImgUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
