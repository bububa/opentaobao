package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumsoutboundAPIRequest 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) API请求
// alibaba.wdk.ums.outbound
//
// 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
type AlibabawdkumsoutboundAPIRequest struct {
	model.Params
	// 出库-ERP下发单请求dto
	_erpOutputOrderDto *ErpOutputOrderDto
}

// NewAlibabawdkumsoutboundRequest 初始化AlibabawdkumsoutboundAPIRequest对象
func NewAlibabawdkumsoutboundRequest() *AlibabawdkumsoutboundAPIRequest {
	return &AlibabawdkumsoutboundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkumsoutboundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.outbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkumsoutboundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkumsoutboundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErpOutputOrderDto is ErpOutputOrderDto Setter
// 出库-ERP下发单请求dto
func (r *AlibabawdkumsoutboundAPIRequest) SetErpOutputOrderDto(_erpOutputOrderDto *ErpOutputOrderDto) error {
	r._erpOutputOrderDto = _erpOutputOrderDto
	r.Set("erp_output_order_dto", _erpOutputOrderDto)
	return nil
}

// GetErpOutputOrderDto ErpOutputOrderDto Getter
func (r AlibabawdkumsoutboundAPIRequest) GetErpOutputOrderDto() *ErpOutputOrderDto {
	return r._erpOutputOrderDto
}
