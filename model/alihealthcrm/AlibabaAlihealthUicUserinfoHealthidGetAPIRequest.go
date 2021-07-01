package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthUicUserinfoHealthidGetAPIRequest
获取健康id API请求
alibaba.alihealth.uic.userinfo.healthid.get

根据支付宝用户ID获取用户健康ID */
type AlibabaAlihealthUicUserinfoHealthidGetAPIRequest struct {
	model.Params
	// 支付宝用户ID
	_alipayUserId string
}

// NewAlibabaAlihealthUicUserinfoHealthidGetRequest 初始化AlibabaAlihealthUicUserinfoHealthidGetAPIRequest对象
func NewAlibabaAlihealthUicUserinfoHealthidGetRequest() *AlibabaAlihealthUicUserinfoHealthidGetAPIRequest {
	return &AlibabaAlihealthUicUserinfoHealthidGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthUicUserinfoHealthidGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.uic.userinfo.healthid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthUicUserinfoHealthidGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlipayUserId Setter
// 支付宝用户ID
func (r *AlibabaAlihealthUicUserinfoHealthidGetAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// Get AlipayUserId Getter
func (r AlibabaAlihealthUicUserinfoHealthidGetAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}
