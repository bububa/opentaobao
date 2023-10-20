package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmemberqrcodeidentifyAPIRequest 扫码识别会员接口 API请求
// alibaba.wdk.member.qrcode.identify
//
// 根据用户输入的付款码（支付宝、盒马、淘宝）、商家等信息，查询当前用户的基本信息及对应会员卡信息
type AlibabawdkmemberqrcodeidentifyAPIRequest struct {
	model.Params
	// 付款码
	_qrCode string
}

// NewAlibabawdkmemberqrcodeidentifyRequest 初始化AlibabawdkmemberqrcodeidentifyAPIRequest对象
func NewAlibabawdkmemberqrcodeidentifyRequest() *AlibabawdkmemberqrcodeidentifyAPIRequest {
	return &AlibabawdkmemberqrcodeidentifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmemberqrcodeidentifyAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.member.qrcode.identify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmemberqrcodeidentifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmemberqrcodeidentifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQrCode is QrCode Setter
// 付款码
func (r *AlibabawdkmemberqrcodeidentifyAPIRequest) SetQrCode(_qrCode string) error {
	r._qrCode = _qrCode
	r.Set("qr_code", _qrCode)
	return nil
}

// GetQrCode QrCode Getter
func (r AlibabawdkmemberqrcodeidentifyAPIRequest) GetQrCode() string {
	return r._qrCode
}
