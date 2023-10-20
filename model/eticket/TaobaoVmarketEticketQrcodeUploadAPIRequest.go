package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketeticketqrcodeuploadAPIRequest 码商二维码图片上传 API请求
// taobao.vmarket.eticket.qrcode.upload
//
// 电子凭证的码商可以通过这个接口，上传他们发送的二维码图片
type TaobaovmarketeticketqrcodeuploadAPIRequest struct {
	model.Params
	// 上传的图片byte[]  小于300K，图片尺寸400*400以内
	_imgBytes *model.File
	// 码商ID
	_codeMerchantId int64
}

// NewTaobaovmarketeticketqrcodeuploadRequest 初始化TaobaovmarketeticketqrcodeuploadAPIRequest对象
func NewTaobaovmarketeticketqrcodeuploadRequest() *TaobaovmarketeticketqrcodeuploadAPIRequest {
	return &TaobaovmarketeticketqrcodeuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovmarketeticketqrcodeuploadAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.qrcode.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovmarketeticketqrcodeuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovmarketeticketqrcodeuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImgBytes is ImgBytes Setter
// 上传的图片byte[]  小于300K，图片尺寸400*400以内
func (r *TaobaovmarketeticketqrcodeuploadAPIRequest) SetImgBytes(_imgBytes *model.File) error {
	r._imgBytes = _imgBytes
	r.Set("img_bytes", _imgBytes)
	return nil
}

// GetImgBytes ImgBytes Getter
func (r TaobaovmarketeticketqrcodeuploadAPIRequest) GetImgBytes() *model.File {
	return r._imgBytes
}

// SetCodeMerchantId is CodeMerchantId Setter
// 码商ID
func (r *TaobaovmarketeticketqrcodeuploadAPIRequest) SetCodeMerchantId(_codeMerchantId int64) error {
	r._codeMerchantId = _codeMerchantId
	r.Set("code_merchant_id", _codeMerchantId)
	return nil
}

// GetCodeMerchantId CodeMerchantId Getter
func (r TaobaovmarketeticketqrcodeuploadAPIRequest) GetCodeMerchantId() int64 {
	return r._codeMerchantId
}
