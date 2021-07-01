package viapi

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/viapi"
)

/* 
通用分割 
aliyun.viapi.imageseg.segment.commonimage

识别输入图像中的视觉中心物体轮廓，与背景进行分离，返回分割后的前景物体图（4通道png透明图）。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
func AliyunViapiImagesegSegmentCommonimage(clt *core.SDKClient, req *viapi.AliyunViapiImagesegSegmentCommonimageAPIRequest, session string) (*viapi.AliyunViapiImagesegSegmentCommonimageAPIResponse, error) {
    var resp viapi.AliyunViapiImagesegSegmentCommonimageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
