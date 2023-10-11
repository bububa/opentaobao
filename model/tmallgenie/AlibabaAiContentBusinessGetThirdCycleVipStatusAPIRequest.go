package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest 天猫精灵商业化获取三方连续包会员状态 API请求
// alibaba.ai.content.business.get.third.cycle.vip.status
//
// 天猫精灵商业化获取三方连续包会员状态
type AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest struct {
	model.Params
	// 获取三方连续包会员状态请求
	_getThirdVipStatusRequest *VipCycleThirdVipStatusRequest
}

// NewAlibabaAiContentBusinessGetThirdCycleVipStatusRequest 初始化AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest对象
func NewAlibabaAiContentBusinessGetThirdCycleVipStatusRequest() *AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest {
	return &AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.content.business.get.third.cycle.vip.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGetThirdVipStatusRequest is GetThirdVipStatusRequest Setter
// 获取三方连续包会员状态请求
func (r *AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest) SetGetThirdVipStatusRequest(_getThirdVipStatusRequest *VipCycleThirdVipStatusRequest) error {
	r._getThirdVipStatusRequest = _getThirdVipStatusRequest
	r.Set("get_third_vip_status_request", _getThirdVipStatusRequest)
	return nil
}

// GetGetThirdVipStatusRequest GetThirdVipStatusRequest Getter
func (r AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest) GetGetThirdVipStatusRequest() *VipCycleThirdVipStatusRequest {
	return r._getThirdVipStatusRequest
}
