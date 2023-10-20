package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunviapiocrcharacterAPIRequest 通用文字识别 API请求
// aliyun.viapi.ocr.character
//
// 获取通用的文字信息。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunviapiocrcharacterAPIRequest struct {
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

// NewAliyunviapiocrcharacterRequest 初始化AliyunviapiocrcharacterAPIRequest对象
func NewAliyunviapiocrcharacterRequest() *AliyunviapiocrcharacterAPIRequest {
	return &AliyunviapiocrcharacterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunviapiocrcharacterAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.ocr.character"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunviapiocrcharacterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunviapiocrcharacterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunviapiocrcharacterAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunviapiocrcharacterAPIRequest) GetImageUrl() string {
	return r._imageUrl
}

// SetImageType is ImageType Setter
// 图片类型, ,取范围[0:ImageURL, 1:ImageContent]
func (r *AliyunviapiocrcharacterAPIRequest) SetImageType(_imageType int64) error {
	r._imageType = _imageType
	r.Set("image_type", _imageType)
	return nil
}

// GetImageType ImageType Getter
func (r AliyunviapiocrcharacterAPIRequest) GetImageType() int64 {
	return r._imageType
}

// SetMinHeight is MinHeight Setter
// 图片中文字的最小高度，单位像素
func (r *AliyunviapiocrcharacterAPIRequest) SetMinHeight(_minHeight int64) error {
	r._minHeight = _minHeight
	r.Set("min_height", _minHeight)
	return nil
}

// GetMinHeight MinHeight Getter
func (r AliyunviapiocrcharacterAPIRequest) GetMinHeight() int64 {
	return r._minHeight
}

// SetOutputProbability is OutputProbability Setter
// 是否输出文字框的概率,取值范围[true:是, false:否]
func (r *AliyunviapiocrcharacterAPIRequest) SetOutputProbability(_outputProbability bool) error {
	r._outputProbability = _outputProbability
	r.Set("output_probability", _outputProbability)
	return nil
}

// GetOutputProbability OutputProbability Getter
func (r AliyunviapiocrcharacterAPIRequest) GetOutputProbability() bool {
	return r._outputProbability
}
