package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiContentBusinessSupplyChargeAPIRequest 供销商品充值接口 API请求
// alibaba.ai.content.business.supply.charge
//
// 供销商品充值接口
type AlibabaAiContentBusinessSupplyChargeAPIRequest struct {
	model.Params
	// 入参对象
	_memberChargeRequest *MemberChargeRequest
}

// NewAlibabaAiContentBusinessSupplyChargeRequest 初始化AlibabaAiContentBusinessSupplyChargeAPIRequest对象
func NewAlibabaAiContentBusinessSupplyChargeRequest() *AlibabaAiContentBusinessSupplyChargeAPIRequest {
	return &AlibabaAiContentBusinessSupplyChargeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAiContentBusinessSupplyChargeAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.content.business.supply.charge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAiContentBusinessSupplyChargeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMemberChargeRequest is MemberChargeRequest Setter
// 入参对象
func (r *AlibabaAiContentBusinessSupplyChargeAPIRequest) SetMemberChargeRequest(_memberChargeRequest *MemberChargeRequest) error {
	r._memberChargeRequest = _memberChargeRequest
	r.Set("member_charge_request", _memberChargeRequest)
	return nil
}

// GetMemberChargeRequest MemberChargeRequest Getter
func (r AlibabaAiContentBusinessSupplyChargeAPIRequest) GetMemberChargeRequest() *MemberChargeRequest {
	return r._memberChargeRequest
}
