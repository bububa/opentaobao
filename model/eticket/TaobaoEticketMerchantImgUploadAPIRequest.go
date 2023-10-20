package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoeticketmerchantimguploadAPIRequest 码商上传二维码图片 API请求
// taobao.eticket.merchant.img.upload
//
// 电子凭证的码商可以通过这个接口，上传二维码图片
type TaobaoeticketmerchantimguploadAPIRequest struct {
	model.Params
	// 二维码图片
	_imgBytes *model.File
}

// NewTaobaoeticketmerchantimguploadRequest 初始化TaobaoeticketmerchantimguploadAPIRequest对象
func NewTaobaoeticketmerchantimguploadRequest() *TaobaoeticketmerchantimguploadAPIRequest {
	return &TaobaoeticketmerchantimguploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoeticketmerchantimguploadAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.img.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoeticketmerchantimguploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoeticketmerchantimguploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImgBytes is ImgBytes Setter
// 二维码图片
func (r *TaobaoeticketmerchantimguploadAPIRequest) SetImgBytes(_imgBytes *model.File) error {
	r._imgBytes = _imgBytes
	r.Set("img_bytes", _imgBytes)
	return nil
}

// GetImgBytes ImgBytes Getter
func (r TaobaoeticketmerchantimguploadAPIRequest) GetImgBytes() *model.File {
	return r._imgBytes
}
