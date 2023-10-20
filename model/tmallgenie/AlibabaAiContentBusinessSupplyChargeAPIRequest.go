package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaicontentbusinesssupplychargeAPIRequest 供销商品充值接口 API请求
// alibaba.ai.content.business.supply.charge
//
// 供销商品充值接口
type AlibabaaicontentbusinesssupplychargeAPIRequest struct {
	model.Params
	// 入参对象
	_memberChargeRequest *MemberChargeRequest
}

// NewAlibabaaicontentbusinesssupplychargeRequest 初始化AlibabaaicontentbusinesssupplychargeAPIRequest对象
func NewAlibabaaicontentbusinesssupplychargeRequest() *AlibabaaicontentbusinesssupplychargeAPIRequest {
	return &AlibabaaicontentbusinesssupplychargeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaicontentbusinesssupplychargeAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.content.business.supply.charge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaicontentbusinesssupplychargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaicontentbusinesssupplychargeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemberChargeRequest is MemberChargeRequest Setter
// 入参对象
func (r *AlibabaaicontentbusinesssupplychargeAPIRequest) SetMemberChargeRequest(_memberChargeRequest *MemberChargeRequest) error {
	r._memberChargeRequest = _memberChargeRequest
	r.Set("member_charge_request", _memberChargeRequest)
	return nil
}

// GetMemberChargeRequest MemberChargeRequest Getter
func (r AlibabaaicontentbusinesssupplychargeAPIRequest) GetMemberChargeRequest() *MemberChargeRequest {
	return r._memberChargeRequest
}
