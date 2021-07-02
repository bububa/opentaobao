package middleclaims

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMiddleClaimsbillReceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.middle.claimsbill.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMiddleClaimsbillReceiveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
