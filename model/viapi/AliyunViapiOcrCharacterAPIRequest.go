package viapi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiOcrCharacterAPIRequest 通用文字识别 API请求
// aliyun.viapi.ocr.character
//
// 获取通用的文字信息。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiOcrCharacterAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
	// 图片类型, ,取范围[0:ImageURL, 1:ImageContent]
	_imageType int64
	// 图片中文字的最小高度，单位像素
	_minHeight int64
	// 是否输出文字框的概率,取值范围[true:是, false:否]
	_outputProbability bool
}

// NewAliyunViapiOcrCharacterRequest 初始化AliyunViapiOcrCharacterAPIRequest对象
func NewAliyunViapiOcrCharacterRequest() *AliyunViapiOcrCharacterAPIRequest {
	return &AliyunViapiOcrCharacterAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunViapiOcrCharacterAPIRequest) Reset() {
	r._imageUrl = ""
	r._imageType = 0
	r._minHeight = 0
	r._outputProbability = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiOcrCharacterAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.ocr.character"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiOcrCharacterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunViapiOcrCharacterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiOcrCharacterAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunViapiOcrCharacterAPIRequest) GetImageUrl() string {
	return r._imageUrl
}

// SetImageType is ImageType Setter
// 图片类型, ,取范围[0:ImageURL, 1:ImageContent]
func (r *AliyunViapiOcrCharacterAPIRequest) SetImageType(_imageType int64) error {
	r._imageType = _imageType
	r.Set("image_type", _imageType)
	return nil
}

// GetImageType ImageType Getter
func (r AliyunViapiOcrCharacterAPIRequest) GetImageType() int64 {
	return r._imageType
}

// SetMinHeight is MinHeight Setter
// 图片中文字的最小高度，单位像素
func (r *AliyunViapiOcrCharacterAPIRequest) SetMinHeight(_minHeight int64) error {
	r._minHeight = _minHeight
	r.Set("min_height", _minHeight)
	return nil
}

// GetMinHeight MinHeight Getter
func (r AliyunViapiOcrCharacterAPIRequest) GetMinHeight() int64 {
	return r._minHeight
}

// SetOutputProbability is OutputProbability Setter
// 是否输出文字框的概率,取值范围[true:是, false:否]
func (r *AliyunViapiOcrCharacterAPIRequest) SetOutputProbability(_outputProbability bool) error {
	r._outputProbability = _outputProbability
	r.Set("output_probability", _outputProbability)
	return nil
}

// GetOutputProbability OutputProbability Getter
func (r AliyunViapiOcrCharacterAPIRequest) GetOutputProbability() bool {
	return r._outputProbability
}

var poolAliyunViapiOcrCharacterAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunViapiOcrCharacterRequest()
	},
}

// GetAliyunViapiOcrCharacterRequest 从 sync.Pool 获取 AliyunViapiOcrCharacterAPIRequest
func GetAliyunViapiOcrCharacterAPIRequest() *AliyunViapiOcrCharacterAPIRequest {
	return poolAliyunViapiOcrCharacterAPIRequest.Get().(*AliyunViapiOcrCharacterAPIRequest)
}

// ReleaseAliyunViapiOcrCharacterAPIRequest 将 AliyunViapiOcrCharacterAPIRequest 放入 sync.Pool
func ReleaseAliyunViapiOcrCharacterAPIRequest(v *AliyunViapiOcrCharacterAPIRequest) {
	v.Reset()
	poolAliyunViapiOcrCharacterAPIRequest.Put(v)
}
