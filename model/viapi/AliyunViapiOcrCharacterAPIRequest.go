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
type AliyunViapiOcrCharacterAPIRequest struct {
    model.Params
    // 待检测图片链接
    _imageUrl   string
    // 图片类型, ,取范围[0:ImageURL, 1:ImageContent]
    _imageType   int64
    // 图片中文字的最小高度，单位像素
    _minHeight   int64
    // 是否输出文字框的概率,取值范围[true:是, false:否]
    _outputProbability   bool
}

// 初始化AliyunViapiOcrCharacterAPIRequest对象
func NewAliyunViapiOcrCharacterRequest() *AliyunViapiOcrCharacterAPIRequest{
    return &AliyunViapiOcrCharacterAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiOcrCharacterAPIRequest) GetApiMethodName() string {
    return "aliyun.viapi.ocr.character"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiOcrCharacterAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiOcrCharacterAPIRequest) SetImageUrl(_imageUrl string) error {
    r._imageUrl = _imageUrl
    r.Set("image_url", _imageUrl)
    return nil
}

// ImageUrl Getter
func (r AliyunViapiOcrCharacterAPIRequest) GetImageUrl() string {
    return r._imageUrl
}
// ImageType Setter
// 图片类型, ,取范围[0:ImageURL, 1:ImageContent]
func (r *AliyunViapiOcrCharacterAPIRequest) SetImageType(_imageType int64) error {
    r._imageType = _imageType
    r.Set("image_type", _imageType)
    return nil
}

// ImageType Getter
func (r AliyunViapiOcrCharacterAPIRequest) GetImageType() int64 {
    return r._imageType
}
// MinHeight Setter
// 图片中文字的最小高度，单位像素
func (r *AliyunViapiOcrCharacterAPIRequest) SetMinHeight(_minHeight int64) error {
    r._minHeight = _minHeight
    r.Set("min_height", _minHeight)
    return nil
}

// MinHeight Getter
func (r AliyunViapiOcrCharacterAPIRequest) GetMinHeight() int64 {
    return r._minHeight
}
// OutputProbability Setter
// 是否输出文字框的概率,取值范围[true:是, false:否]
func (r *AliyunViapiOcrCharacterAPIRequest) SetOutputProbability(_outputProbability bool) error {
    r._outputProbability = _outputProbability
    r.Set("output_probability", _outputProbability)
    return nil
}

// OutputProbability Getter
func (r AliyunViapiOcrCharacterAPIRequest) GetOutputProbability() bool {
    return r._outputProbability
}
