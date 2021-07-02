package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendngImageUploadAPIRequest 售货机商品图片上传 API请求
// alibaba.lst.vendng.image.upload
//
// 零售通自动售货机商品图片上传接口，主要为ISV厂商提供图片同步的通道，从而建立统一的商品图片库。
type AlibabaLstVendngImageUploadAPIRequest struct {
	model.Params
	// 图片文件字节数组
	_imgBytes *model.File
}

// NewAlibabaLstVendngImageUploadRequest 初始化AlibabaLstVendngImageUploadAPIRequest对象
func NewAlibabaLstVendngImageUploadRequest() *AlibabaLstVendngImageUploadAPIRequest {
	return &AlibabaLstVendngImageUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVendngImageUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vendng.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVendngImageUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ImgBytes Setter
// 图片文件字节数组
func (r *AlibabaLstVendngImageUploadAPIRequest) SetImgBytes(_imgBytes *model.File) error {
	r._imgBytes = _imgBytes
	r.Set("img_bytes", _imgBytes)
	return nil
}

// Get ImgBytes Getter
func (r AlibabaLstVendngImageUploadAPIRequest) GetImgBytes() *model.File {
	return r._imgBytes
}
