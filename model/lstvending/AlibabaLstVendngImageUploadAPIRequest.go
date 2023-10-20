package lstvending

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstVendngImageUploadAPIRequest) Reset() {
	r._imgBytes = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVendngImageUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vendng.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVendngImageUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstVendngImageUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImgBytes is ImgBytes Setter
// 图片文件字节数组
func (r *AlibabaLstVendngImageUploadAPIRequest) SetImgBytes(_imgBytes *model.File) error {
	r._imgBytes = _imgBytes
	r.Set("img_bytes", _imgBytes)
	return nil
}

// GetImgBytes ImgBytes Getter
func (r AlibabaLstVendngImageUploadAPIRequest) GetImgBytes() *model.File {
	return r._imgBytes
}

var poolAlibabaLstVendngImageUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstVendngImageUploadRequest()
	},
}

// GetAlibabaLstVendngImageUploadRequest 从 sync.Pool 获取 AlibabaLstVendngImageUploadAPIRequest
func GetAlibabaLstVendngImageUploadAPIRequest() *AlibabaLstVendngImageUploadAPIRequest {
	return poolAlibabaLstVendngImageUploadAPIRequest.Get().(*AlibabaLstVendngImageUploadAPIRequest)
}

// ReleaseAlibabaLstVendngImageUploadAPIRequest 将 AlibabaLstVendngImageUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstVendngImageUploadAPIRequest(v *AlibabaLstVendngImageUploadAPIRequest) {
	v.Reset()
	poolAlibabaLstVendngImageUploadAPIRequest.Put(v)
}
