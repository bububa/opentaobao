package eticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketQrcodeUploadAPIRequest 码商二维码图片上传 API请求
// taobao.vmarket.eticket.qrcode.upload
//
// 电子凭证的码商可以通过这个接口，上传他们发送的二维码图片
type TaobaoVmarketEticketQrcodeUploadAPIRequest struct {
	model.Params
	// 上传的图片byte[]  小于300K，图片尺寸400*400以内
	_imgBytes *model.File
	// 码商ID
	_codeMerchantId int64
}

// NewTaobaoVmarketEticketQrcodeUploadRequest 初始化TaobaoVmarketEticketQrcodeUploadAPIRequest对象
func NewTaobaoVmarketEticketQrcodeUploadRequest() *TaobaoVmarketEticketQrcodeUploadAPIRequest {
	return &TaobaoVmarketEticketQrcodeUploadAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoVmarketEticketQrcodeUploadAPIRequest) Reset() {
	r._imgBytes = nil
	r._codeMerchantId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketQrcodeUploadAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.qrcode.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketQrcodeUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVmarketEticketQrcodeUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImgBytes is ImgBytes Setter
// 上传的图片byte[]  小于300K，图片尺寸400*400以内
func (r *TaobaoVmarketEticketQrcodeUploadAPIRequest) SetImgBytes(_imgBytes *model.File) error {
	r._imgBytes = _imgBytes
	r.Set("img_bytes", _imgBytes)
	return nil
}

// GetImgBytes ImgBytes Getter
func (r TaobaoVmarketEticketQrcodeUploadAPIRequest) GetImgBytes() *model.File {
	return r._imgBytes
}

// SetCodeMerchantId is CodeMerchantId Setter
// 码商ID
func (r *TaobaoVmarketEticketQrcodeUploadAPIRequest) SetCodeMerchantId(_codeMerchantId int64) error {
	r._codeMerchantId = _codeMerchantId
	r.Set("code_merchant_id", _codeMerchantId)
	return nil
}

// GetCodeMerchantId CodeMerchantId Getter
func (r TaobaoVmarketEticketQrcodeUploadAPIRequest) GetCodeMerchantId() int64 {
	return r._codeMerchantId
}

var poolTaobaoVmarketEticketQrcodeUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoVmarketEticketQrcodeUploadRequest()
	},
}

// GetTaobaoVmarketEticketQrcodeUploadRequest 从 sync.Pool 获取 TaobaoVmarketEticketQrcodeUploadAPIRequest
func GetTaobaoVmarketEticketQrcodeUploadAPIRequest() *TaobaoVmarketEticketQrcodeUploadAPIRequest {
	return poolTaobaoVmarketEticketQrcodeUploadAPIRequest.Get().(*TaobaoVmarketEticketQrcodeUploadAPIRequest)
}

// ReleaseTaobaoVmarketEticketQrcodeUploadAPIRequest 将 TaobaoVmarketEticketQrcodeUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoVmarketEticketQrcodeUploadAPIRequest(v *TaobaoVmarketEticketQrcodeUploadAPIRequest) {
	v.Reset()
	poolTaobaoVmarketEticketQrcodeUploadAPIRequest.Put(v)
}
