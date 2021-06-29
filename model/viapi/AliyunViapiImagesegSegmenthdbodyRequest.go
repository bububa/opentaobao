package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
高清人体分割 APIRequest
aliyun.viapi.imageseg.segmenthdbody

对输入图像中包含的人进行分割，输出结果透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiImagesegSegmenthdbodyRequest struct {
    model.Params

    // 待检测图片链接
    imageUrl   string 

}

func NewAliyunViapiImagesegSegmenthdbodyRequest() *AliyunViapiImagesegSegmenthdbodyRequest{
    return &AliyunViapiImagesegSegmenthdbodyRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunViapiImagesegSegmenthdbodyRequest) GetApiMethodName() string {
    return "aliyun.viapi.imageseg.segmenthdbody"
}

func (r AliyunViapiImagesegSegmenthdbodyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunViapiImagesegSegmenthdbodyRequest) SetImageUrl(imageUrl string) error {
    r.imageUrl = imageUrl
    r.Set("image_url", imageUrl)
    return nil
}

func (r AliyunViapiImagesegSegmenthdbodyRequest) GetImageUrl() string {
    return r.imageUrl
}

