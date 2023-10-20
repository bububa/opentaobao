package tmallgenie

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAiContentBusinessSupplyChargeAPIRequest) Reset() {
	r._memberChargeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAiContentBusinessSupplyChargeAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.content.business.supply.charge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAiContentBusinessSupplyChargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAiContentBusinessSupplyChargeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAiContentBusinessSupplyChargeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAiContentBusinessSupplyChargeRequest()
	},
}

// GetAlibabaAiContentBusinessSupplyChargeRequest 从 sync.Pool 获取 AlibabaAiContentBusinessSupplyChargeAPIRequest
func GetAlibabaAiContentBusinessSupplyChargeAPIRequest() *AlibabaAiContentBusinessSupplyChargeAPIRequest {
	return poolAlibabaAiContentBusinessSupplyChargeAPIRequest.Get().(*AlibabaAiContentBusinessSupplyChargeAPIRequest)
}

// ReleaseAlibabaAiContentBusinessSupplyChargeAPIRequest 将 AlibabaAiContentBusinessSupplyChargeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAiContentBusinessSupplyChargeAPIRequest(v *AlibabaAiContentBusinessSupplyChargeAPIRequest) {
	v.Reset()
	poolAlibabaAiContentBusinessSupplyChargeAPIRequest.Put(v)
}
