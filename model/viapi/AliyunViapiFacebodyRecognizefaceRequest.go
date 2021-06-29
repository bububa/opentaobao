package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
人脸属性识别 API请求
aliyun.viapi.facebody.recognizeface

输入图片之后，在人脸检测定位返回结果的基础上，识别各个检测人脸的四种属性，返回性别（男/女）、年龄、表情（笑/不笑）、眼镜（戴/不戴）；并可返回人脸的1024维深度学习特征。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiFacebodyRecognizefaceRequest struct {
    model.Params
    // 待检测图片链接
    imageUrl   string
    // 图片类型, ,取值范围[0:ImageURL, 1:ImageContent]
    imageType   int64
}

// 初始化AliyunViapiFacebodyRecognizefaceRequest对象
func NewAliyunViapiFacebodyRecognizefaceRequest() *AliyunViapiFacebodyRecognizefaceRequest{
    return &AliyunViapiFacebodyRecognizefaceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiFacebodyRecognizefaceRequest) GetApiMethodName() string {
    return "aliyun.viapi.facebody.recognizeface"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiFacebodyRecognizefaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiFacebodyRecognizefaceRequest) SetImageUrl(imageUrl string) error {
    r.imageUrl = imageUrl
    r.Set("image_url", imageUrl)
    return nil
}

// ImageUrl Getter
func (r AliyunViapiFacebodyRecognizefaceRequest) GetImageUrl() string {
    return r.imageUrl
}
// ImageType Setter
// 图片类型, ,取值范围[0:ImageURL, 1:ImageContent]
func (r *AliyunViapiFacebodyRecognizefaceRequest) SetImageType(imageType int64) error {
    r.imageType = imageType
    r.Set("image_type", imageType)
    return nil
}

// ImageType Getter
func (r AliyunViapiFacebodyRecognizefaceRequest) GetImageType() int64 {
    return r.imageType
}
