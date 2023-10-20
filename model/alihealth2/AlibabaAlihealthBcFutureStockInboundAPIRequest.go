package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthbcfuturestockinboundAPIRequest 供应商上报期货库存 API请求
// alibaba.alihealth.bc.future.stock.inbound
//
// 供应商上报期货库存
type AlibabaalihealthbcfuturestockinboundAPIRequest struct {
	model.Params
	// 请求体
	_futureInboundReqDto *FutureInboundReqDto
}

// NewAlibabaalihealthbcfuturestockinboundRequest 初始化AlibabaalihealthbcfuturestockinboundAPIRequest对象
func NewAlibabaalihealthbcfuturestockinboundRequest() *AlibabaalihealthbcfuturestockinboundAPIRequest {
	return &AlibabaalihealthbcfuturestockinboundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthbcfuturestockinboundAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.bc.future.stock.inbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthbcfuturestockinboundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthbcfuturestockinboundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFutureInboundReqDto is FutureInboundReqDto Setter
// 请求体
func (r *AlibabaalihealthbcfuturestockinboundAPIRequest) SetFutureInboundReqDto(_futureInboundReqDto *FutureInboundReqDto) error {
	r._futureInboundReqDto = _futureInboundReqDto
	r.Set("future_inbound_req_dto", _futureInboundReqDto)
	return nil
}

// GetFutureInboundReqDto FutureInboundReqDto Getter
func (r AlibabaalihealthbcfuturestockinboundAPIRequest) GetFutureInboundReqDto() *FutureInboundReqDto {
	return r._futureInboundReqDto
}
