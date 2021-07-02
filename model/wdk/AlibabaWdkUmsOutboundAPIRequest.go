package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsOutboundAPIRequest 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) API请求
// alibaba.wdk.ums.outbound
//
// 出库-ERP下发单(新接口，包含调拨出库单和退货出库单等)
type AlibabaWdkUmsOutboundAPIRequest struct {
	model.Params
	// 出库-ERP下发单请求dto
	_erpOutputOrderDto *ErpOutputOrderDto
}

// NewAlibabaWdkUmsOutboundRequest 初始化AlibabaWdkUmsOutboundAPIRequest对象
func NewAlibabaWdkUmsOutboundRequest() *AlibabaWdkUmsOutboundAPIRequest {
	return &AlibabaWdkUmsOutboundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsOutboundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.outbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsOutboundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ErpOutputOrderDto Setter
// 出库-ERP下发单请求dto
func (r *AlibabaWdkUmsOutboundAPIRequest) SetErpOutputOrderDto(_erpOutputOrderDto *ErpOutputOrderDto) error {
	r._erpOutputOrderDto = _erpOutputOrderDto
	r.Set("erp_output_order_dto", _erpOutputOrderDto)
	return nil
}

// Get ErpOutputOrderDto Getter
func (r AlibabaWdkUmsOutboundAPIRequest) GetErpOutputOrderDto() *ErpOutputOrderDto {
	return r._erpOutputOrderDto
}
