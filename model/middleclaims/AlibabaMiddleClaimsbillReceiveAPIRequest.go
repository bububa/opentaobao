package middleclaims

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamiddleclaimsbillreceiveAPIRequest 国际化中台服务域接收理赔账单 API请求
// alibaba.middle.claimsbill.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔打款的处理后，将该打款结果返回至服务域
type AlibabamiddleclaimsbillreceiveAPIRequest struct {
	model.Params
	// 理赔账单实体
	_claimsBillDto *ClaimsBillDto
}

// NewAlibabamiddleclaimsbillreceiveRequest 初始化AlibabamiddleclaimsbillreceiveAPIRequest对象
func NewAlibabamiddleclaimsbillreceiveRequest() *AlibabamiddleclaimsbillreceiveAPIRequest {
	return &AlibabamiddleclaimsbillreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamiddleclaimsbillreceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.middle.claimsbill.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamiddleclaimsbillreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamiddleclaimsbillreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClaimsBillDto is ClaimsBillDto Setter
// 理赔账单实体
func (r *AlibabamiddleclaimsbillreceiveAPIRequest) SetClaimsBillDto(_claimsBillDto *ClaimsBillDto) error {
	r._claimsBillDto = _claimsBillDto
	r.Set("claims_bill_dto", _claimsBillDto)
	return nil
}

// GetClaimsBillDto ClaimsBillDto Getter
func (r AlibabamiddleclaimsbillreceiveAPIRequest) GetClaimsBillDto() *ClaimsBillDto {
	return r._claimsBillDto
}
