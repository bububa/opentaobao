package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthbcfuturestockoutboundAPIRequest 供应商期货出库 API请求
// alibaba.alihealth.bc.future.stock.outbound
//
// 供应商期货出库
type AlibabaalihealthbcfuturestockoutboundAPIRequest struct {
	model.Params
	// 请求体
	_futureOutboundReqDto *FutureOutboundReqDto
}

// NewAlibabaalihealthbcfuturestockoutboundRequest 初始化AlibabaalihealthbcfuturestockoutboundAPIRequest对象
func NewAlibabaalihealthbcfuturestockoutboundRequest() *AlibabaalihealthbcfuturestockoutboundAPIRequest {
	return &AlibabaalihealthbcfuturestockoutboundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthbcfuturestockoutboundAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.bc.future.stock.outbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthbcfuturestockoutboundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthbcfuturestockoutboundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFutureOutboundReqDto is FutureOutboundReqDto Setter
// 请求体
func (r *AlibabaalihealthbcfuturestockoutboundAPIRequest) SetFutureOutboundReqDto(_futureOutboundReqDto *FutureOutboundReqDto) error {
	r._futureOutboundReqDto = _futureOutboundReqDto
	r.Set("future_outbound_req_dto", _futureOutboundReqDto)
	return nil
}

// GetFutureOutboundReqDto FutureOutboundReqDto Getter
func (r AlibabaalihealthbcfuturestockoutboundAPIRequest) GetFutureOutboundReqDto() *FutureOutboundReqDto {
	return r._futureOutboundReqDto
}
