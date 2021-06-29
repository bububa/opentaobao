package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通用文字识别 API请求
aliyun.viapi.ocr.character

获取通用的文字信息。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiOcrCharacterRequest struct {
    model.Params
    // 待检测图片链接
    imageUrl   string
    // 图片类型, ,取范围[0:ImageURL, 1:ImageContent]
    imageType   int64
    // 图片中文字的最小高度，单位像素
    minHeight   int64
    // 是否输出文字框的概率,取值范围[true:是, false:否]
    outputProbability   bool
}

// 初始化AliyunViapiOcrCharacterRequest对象
func NewAliyunViapiOcrCharacterRequest() *AliyunViapiOcrCharacterRequest{
    return &AliyunViapiOcrCharacterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiOcrCharacterRequest) GetApiMethodName() string {
    return "aliyun.viapi.ocr.character"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiOcrCharacterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiOcrCharacterRequest) SetImageUrl(imageUrl string) error {
    r.imageUrl = imageUrl
    r.Set("image_url", imageUrl)
    return nil
}

// ImageUrl Getter
func (r AliyunViapiOcrCharacterRequest) GetImageUrl() string {
    return r.imageUrl
}
// ImageType Setter
// 图片类型, ,取范围[0:ImageURL, 1:ImageContent]
func (r *AliyunViapiOcrCharacterRequest) SetImageType(imageType int64) error {
    r.imageType = imageType
    r.Set("image_type", imageType)
    return nil
}

// ImageType Getter
func (r AliyunViapiOcrCharacterRequest) GetImageType() int64 {
    return r.imageType
}
// MinHeight Setter
// 图片中文字的最小高度，单位像素
func (r *AliyunViapiOcrCharacterRequest) SetMinHeight(minHeight int64) error {
    r.minHeight = minHeight
    r.Set("min_height", minHeight)
    return nil
}

// MinHeight Getter
func (r AliyunViapiOcrCharacterRequest) GetMinHeight() int64 {
    return r.minHeight
}
// OutputProbability Setter
// 是否输出文字框的概率,取值范围[true:是, false:否]
func (r *AliyunViapiOcrCharacterRequest) SetOutputProbability(outputProbability bool) error {
    r.outputProbability = outputProbability
    r.Set("output_probability", outputProbability)
    return nil
}

// OutputProbability Getter
func (r AliyunViapiOcrCharacterRequest) GetOutputProbability() bool {
    return r.outputProbability
}
