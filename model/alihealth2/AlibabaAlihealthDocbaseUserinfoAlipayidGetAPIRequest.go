package alihealth2

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.docbase.userinfo.alipayid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
