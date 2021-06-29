package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
人脸检测定位 APIRequest
aliyun.viapi.facebody.detectface

输入图片之后，在人脸检测定位返回结果的基础上，识别各个检测人脸的四种属性，返回性别（男/女）、年龄、表情（笑/不笑）、眼镜（戴/不戴）；并可返回人脸的1024维深度学习特征、(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiFacebodyDetectfaceRequest struct {
    model.Params

    // 图片url地址(http/https)
    imageUrl   string 

    // 0: 通过url识别，参数image_url不为空；1: 通过图片content识别，参数content不为空 支持图片格式：JPEG、JPG、BMP、PNG
    imageType   int64 

}

func NewAliyunViapiFacebodyDetectfaceRequest() *AliyunViapiFacebodyDetectfaceRequest{
    return &AliyunViapiFacebodyDetectfaceRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunViapiFacebodyDetectfaceRequest) GetApiMethodName() string {
    return "aliyun.viapi.facebody.detectface"
}

func (r AliyunViapiFacebodyDetectfaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunViapiFacebodyDetectfaceRequest) SetImageUrl(imageUrl string) error {
    r.imageUrl = imageUrl
    r.Set("image_url", imageUrl)
    return nil
}

func (r AliyunViapiFacebodyDetectfaceRequest) GetImageUrl() string {
    return r.imageUrl
}

func (r *AliyunViapiFacebodyDetectfaceRequest) SetImageType(imageType int64) error {
    r.imageType = imageType
    r.Set("image_type", imageType)
    return nil
}

func (r AliyunViapiFacebodyDetectfaceRequest) GetImageType() int64 {
    return r.imageType
}

