package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest 家居属性识别 API请求
// aliyun.viapi.goodstech.recognize.furniture.attribute
//
// 识别输入的家居模型图的风格，目前支持16种风格识别。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunViapiGoodstechRecognizeFurnitureAttributeRequest 初始化AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest对象
func NewAliyunViapiGoodstechRecognizeFurnitureAttributeRequest() *AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest {
	return &AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.goodstech.recognize.furniture.attribute"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
