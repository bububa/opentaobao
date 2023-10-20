package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvendngimageuploadAPIRequest 售货机商品图片上传 API请求
// alibaba.lst.vendng.image.upload
//
// 零售通自动售货机商品图片上传接口，主要为ISV厂商提供图片同步的通道，从而建立统一的商品图片库。
type AlibabalstvendngimageuploadAPIRequest struct {
	model.Params
	// 图片文件字节数组
	_imgBytes *model.File
}

// NewAlibabalstvendngimageuploadRequest 初始化AlibabalstvendngimageuploadAPIRequest对象
func NewAlibabalstvendngimageuploadRequest() *AlibabalstvendngimageuploadAPIRequest {
	return &AlibabalstvendngimageuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstvendngimageuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vendng.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstvendngimageuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstvendngimageuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImgBytes is ImgBytes Setter
// 图片文件字节数组
func (r *AlibabalstvendngimageuploadAPIRequest) SetImgBytes(_imgBytes *model.File) error {
	r._imgBytes = _imgBytes
	r.Set("img_bytes", _imgBytes)
	return nil
}

// GetImgBytes ImgBytes Getter
func (r AlibabalstvendngimageuploadAPIRequest) GetImgBytes() *model.File {
	return r._imgBytes
}
