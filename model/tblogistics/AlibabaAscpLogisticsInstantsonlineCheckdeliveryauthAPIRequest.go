package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIRequest 同城配送在线下单检查授权 API请求
// alibaba.ascp.logistics.instantsonline.checkdeliveryauth
//
// 同城配送在线下单检查授权
type AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIRequest struct {
	model.Params
	// 业务类型，INSTANT_ONLINE：同城配送-在线下单
	_bizType string
}

// NewAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthRequest 初始化AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIRequest对象
func NewAlibabaAscpLogisticsInstantsonlineCheckdeliveryauthRequest() *AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIRequest {
	return &AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.instantsonline.checkdeliveryauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizType is BizType Setter
// 业务类型，INSTANT_ONLINE：同城配送-在线下单
func (r *AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaAscpLogisticsInstantsonlineCheckdeliveryauthAPIRequest) GetBizType() string {
	return r._bizType
}
