package middleclaims

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMiddleClaimsbillReceiveAPIRequest 国际化中台服务域接收理赔账单 API请求
// alibaba.middle.claimsbill.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔打款的处理后，将该打款结果返回至服务域
type AlibabaMiddleClaimsbillReceiveAPIRequest struct {
	model.Params
	// 理赔账单实体
	_claimsBillDto *ClaimsBillDto
}

// NewAlibabaMiddleClaimsbillReceiveRequest 初始化AlibabaMiddleClaimsbillReceiveAPIRequest对象
func NewAlibabaMiddleClaimsbillReceiveRequest() *AlibabaMiddleClaimsbillReceiveAPIRequest {
	return &AlibabaMiddleClaimsbillReceiveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMiddleClaimsbillReceiveAPIRequest) Reset() {
	r._claimsBillDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMiddleClaimsbillReceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.middle.claimsbill.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMiddleClaimsbillReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMiddleClaimsbillReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClaimsBillDto is ClaimsBillDto Setter
// 理赔账单实体
func (r *AlibabaMiddleClaimsbillReceiveAPIRequest) SetClaimsBillDto(_claimsBillDto *ClaimsBillDto) error {
	r._claimsBillDto = _claimsBillDto
	r.Set("claims_bill_dto", _claimsBillDto)
	return nil
}

// GetClaimsBillDto ClaimsBillDto Getter
func (r AlibabaMiddleClaimsbillReceiveAPIRequest) GetClaimsBillDto() *ClaimsBillDto {
	return r._claimsBillDto
}

var poolAlibabaMiddleClaimsbillReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMiddleClaimsbillReceiveRequest()
	},
}

// GetAlibabaMiddleClaimsbillReceiveRequest 从 sync.Pool 获取 AlibabaMiddleClaimsbillReceiveAPIRequest
func GetAlibabaMiddleClaimsbillReceiveAPIRequest() *AlibabaMiddleClaimsbillReceiveAPIRequest {
	return poolAlibabaMiddleClaimsbillReceiveAPIRequest.Get().(*AlibabaMiddleClaimsbillReceiveAPIRequest)
}

// ReleaseAlibabaMiddleClaimsbillReceiveAPIRequest 将 AlibabaMiddleClaimsbillReceiveAPIRequest 放入 sync.Pool
func ReleaseAlibabaMiddleClaimsbillReceiveAPIRequest(v *AlibabaMiddleClaimsbillReceiveAPIRequest) {
	v.Reset()
	poolAlibabaMiddleClaimsbillReceiveAPIRequest.Put(v)
}
