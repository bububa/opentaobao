package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIRequest 同城配送在线下单检查授权 API请求
// alibaba.ascp.logistics.instantsonline.checkdeliveryauth
//
// 同城配送在线下单检查授权
type AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIRequest struct {
	model.Params
	// 业务类型，INSTANT_ONLINE：同城配送-在线下单
	_bizType string
}

// NewAlibabaascplogisticsinstantsonlinecheckdeliveryauthRequest 初始化AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIRequest对象
func NewAlibabaascplogisticsinstantsonlinecheckdeliveryauthRequest() *AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIRequest {
	return &AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.instantsonline.checkdeliveryauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizType is BizType Setter
// 业务类型，INSTANT_ONLINE：同城配送-在线下单
func (r *AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIRequest) GetBizType() string {
	return r._bizType
}
