package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMemberQrcodeIdentifyAPIRequest 扫码识别会员接口 API请求
// alibaba.wdk.member.qrcode.identify
//
// 根据用户输入的付款码（支付宝、盒马、淘宝）、商家等信息，查询当前用户的基本信息及对应会员卡信息
type AlibabaWdkMemberQrcodeIdentifyAPIRequest struct {
	model.Params
	// 付款码
	_qrCode string
}

// NewAlibabaWdkMemberQrcodeIdentifyRequest 初始化AlibabaWdkMemberQrcodeIdentifyAPIRequest对象
func NewAlibabaWdkMemberQrcodeIdentifyRequest() *AlibabaWdkMemberQrcodeIdentifyAPIRequest {
	return &AlibabaWdkMemberQrcodeIdentifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMemberQrcodeIdentifyAPIRequest) Reset() {
	r._qrCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMemberQrcodeIdentifyAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.member.qrcode.identify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMemberQrcodeIdentifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMemberQrcodeIdentifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQrCode is QrCode Setter
// 付款码
func (r *AlibabaWdkMemberQrcodeIdentifyAPIRequest) SetQrCode(_qrCode string) error {
	r._qrCode = _qrCode
	r.Set("qr_code", _qrCode)
	return nil
}

// GetQrCode QrCode Getter
func (r AlibabaWdkMemberQrcodeIdentifyAPIRequest) GetQrCode() string {
	return r._qrCode
}

var poolAlibabaWdkMemberQrcodeIdentifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMemberQrcodeIdentifyRequest()
	},
}

// GetAlibabaWdkMemberQrcodeIdentifyRequest 从 sync.Pool 获取 AlibabaWdkMemberQrcodeIdentifyAPIRequest
func GetAlibabaWdkMemberQrcodeIdentifyAPIRequest() *AlibabaWdkMemberQrcodeIdentifyAPIRequest {
	return poolAlibabaWdkMemberQrcodeIdentifyAPIRequest.Get().(*AlibabaWdkMemberQrcodeIdentifyAPIRequest)
}

// ReleaseAlibabaWdkMemberQrcodeIdentifyAPIRequest 将 AlibabaWdkMemberQrcodeIdentifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMemberQrcodeIdentifyAPIRequest(v *AlibabaWdkMemberQrcodeIdentifyAPIRequest) {
	v.Reset()
	poolAlibabaWdkMemberQrcodeIdentifyAPIRequest.Put(v)
}
