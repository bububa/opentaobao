package viapi

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/viapi"
)

/* 
高清人体分割 
aliyun.viapi.imageseg.segmenthdbody

对输入图像中包含的人进行分割，输出结果透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
func AliyunViapiImagesegSegmenthdbody(clt *core.SDKClient, req *viapi.AliyunViapiImagesegSegmenthdbodyAPIRequest, session string) (*viapi.AliyunViapiImagesegSegmenthdbodyAPIResponse, error) {
    var resp viapi.AliyunViapiImagesegSegmenthdbodyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
