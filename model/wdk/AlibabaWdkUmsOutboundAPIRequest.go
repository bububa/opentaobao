package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkUmsOutboundAPIRequest) Reset() {
	r._erpOutputOrderDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsOutboundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.outbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsOutboundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkUmsOutboundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErpOutputOrderDto is ErpOutputOrderDto Setter
// 出库-ERP下发单请求dto
func (r *AlibabaWdkUmsOutboundAPIRequest) SetErpOutputOrderDto(_erpOutputOrderDto *ErpOutputOrderDto) error {
	r._erpOutputOrderDto = _erpOutputOrderDto
	r.Set("erp_output_order_dto", _erpOutputOrderDto)
	return nil
}

// GetErpOutputOrderDto ErpOutputOrderDto Getter
func (r AlibabaWdkUmsOutboundAPIRequest) GetErpOutputOrderDto() *ErpOutputOrderDto {
	return r._erpOutputOrderDto
}

var poolAlibabaWdkUmsOutboundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkUmsOutboundRequest()
	},
}

// GetAlibabaWdkUmsOutboundRequest 从 sync.Pool 获取 AlibabaWdkUmsOutboundAPIRequest
func GetAlibabaWdkUmsOutboundAPIRequest() *AlibabaWdkUmsOutboundAPIRequest {
	return poolAlibabaWdkUmsOutboundAPIRequest.Get().(*AlibabaWdkUmsOutboundAPIRequest)
}

// ReleaseAlibabaWdkUmsOutboundAPIRequest 将 AlibabaWdkUmsOutboundAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkUmsOutboundAPIRequest(v *AlibabaWdkUmsOutboundAPIRequest) {
	v.Reset()
	poolAlibabaWdkUmsOutboundAPIRequest.Put(v)
}
