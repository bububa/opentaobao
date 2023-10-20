package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaicontentbusinessgetthirdcyclevipstatusAPIRequest 天猫精灵商业化获取三方连续包会员状态 API请求
// alibaba.ai.content.business.get.third.cycle.vip.status
//
// 天猫精灵商业化获取三方连续包会员状态
type AlibabaaicontentbusinessgetthirdcyclevipstatusAPIRequest struct {
	model.Params
	// 获取三方连续包会员状态请求
	_getThirdVipStatusRequest *VipCycleThirdVipStatusRequest
}

// NewAlibabaaicontentbusinessgetthirdcyclevipstatusRequest 初始化AlibabaaicontentbusinessgetthirdcyclevipstatusAPIRequest对象
func NewAlibabaaicontentbusinessgetthirdcyclevipstatusRequest() *AlibabaaicontentbusinessgetthirdcyclevipstatusAPIRequest {
	return &AlibabaaicontentbusinessgetthirdcyclevipstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaicontentbusinessgetthirdcyclevipstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.content.business.get.third.cycle.vip.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaicontentbusinessgetthirdcyclevipstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaicontentbusinessgetthirdcyclevipstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGetThirdVipStatusRequest is GetThirdVipStatusRequest Setter
// 获取三方连续包会员状态请求
func (r *AlibabaaicontentbusinessgetthirdcyclevipstatusAPIRequest) SetGetThirdVipStatusRequest(_getThirdVipStatusRequest *VipCycleThirdVipStatusRequest) error {
	r._getThirdVipStatusRequest = _getThirdVipStatusRequest
	r.Set("get_third_vip_status_request", _getThirdVipStatusRequest)
	return nil
}

// GetGetThirdVipStatusRequest GetThirdVipStatusRequest Getter
func (r AlibabaaicontentbusinessgetthirdcyclevipstatusAPIRequest) GetGetThirdVipStatusRequest() *VipCycleThirdVipStatusRequest {
	return r._getThirdVipStatusRequest
}
