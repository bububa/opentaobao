package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest 根据健康ID获取支付宝ID API请求
// alibaba.alihealth.docbase.userinfo.alipayid.get
//
// 根据健康ID获取支付宝ID
type AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest struct {
	model.Params
	// 阿里健康ID
	_alihealthUserId string
	// 渠道alipay taobao uc gaode
	_appChannel string
}

// NewAlibabaAlihealthDocbaseUserinfoAlipayidGetRequest 初始化AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest对象
func NewAlibabaAlihealthDocbaseUserinfoAlipayidGetRequest() *AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest {
	return &AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest) Reset() {
	r._alihealthUserId = ""
	r._appChannel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.docbase.userinfo.alipayid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlihealthUserId is AlihealthUserId Setter
// 阿里健康ID
func (r *AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest) SetAlihealthUserId(_alihealthUserId string) error {
	r._alihealthUserId = _alihealthUserId
	r.Set("alihealth_user_id", _alihealthUserId)
	return nil
}

// GetAlihealthUserId AlihealthUserId Getter
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest) GetAlihealthUserId() string {
	return r._alihealthUserId
}

// SetAppChannel is AppChannel Setter
// 渠道alipay taobao uc gaode
func (r *AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest) SetAppChannel(_appChannel string) error {
	r._appChannel = _appChannel
	r.Set("app_channel", _appChannel)
	return nil
}

// GetAppChannel AppChannel Getter
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest) GetAppChannel() string {
	return r._appChannel
}

var poolAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDocbaseUserinfoAlipayidGetRequest()
	},
}

// GetAlibabaAlihealthDocbaseUserinfoAlipayidGetRequest 从 sync.Pool 获取 AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest
func GetAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest() *AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest {
	return poolAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest.Get().(*AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest)
}

// ReleaseAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest 将 AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest(v *AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest.Put(v)
}
