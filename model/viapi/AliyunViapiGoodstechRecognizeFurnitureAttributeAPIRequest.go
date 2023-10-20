package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunviapigoodstechrecognizefurnitureattributeAPIRequest 家居属性识别 API请求
// aliyun.viapi.goodstech.recognize.furniture.attribute
//
// 识别输入的家居模型图的风格，目前支持16种风格识别。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunviapigoodstechrecognizefurnitureattributeAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunviapigoodstechrecognizefurnitureattributeRequest 初始化AliyunviapigoodstechrecognizefurnitureattributeAPIRequest对象
func NewAliyunviapigoodstechrecognizefurnitureattributeRequest() *AliyunviapigoodstechrecognizefurnitureattributeAPIRequest {
	return &AliyunviapigoodstechrecognizefurnitureattributeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunviapigoodstechrecognizefurnitureattributeAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.goodstech.recognize.furniture.attribute"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunviapigoodstechrecognizefurnitureattributeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunviapigoodstechrecognizefurnitureattributeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunviapigoodstechrecognizefurnitureattributeAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunviapigoodstechrecognizefurnitureattributeAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
