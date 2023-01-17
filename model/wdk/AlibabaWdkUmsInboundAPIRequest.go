package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsInboundAPIRequest 入库-ERP下发单 API请求
// alibaba.wdk.ums.inbound
//
// 入库-ERP下发单
type AlibabaWdkUmsInboundAPIRequest struct {
	model.Params
	// 1
	_erpArrivalnoticeDto *ErpArrivalNoticeDto
}

// NewAlibabaWdkUmsInboundRequest 初始化AlibabaWdkUmsInboundAPIRequest对象
func NewAlibabaWdkUmsInboundRequest() *AlibabaWdkUmsInboundAPIRequest {
	return &AlibabaWdkUmsInboundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsInboundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.inbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsInboundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkUmsInboundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErpArrivalnoticeDto is ErpArrivalnoticeDto Setter
// 1
func (r *AlibabaWdkUmsInboundAPIRequest) SetErpArrivalnoticeDto(_erpArrivalnoticeDto *ErpArrivalNoticeDto) error {
	r._erpArrivalnoticeDto = _erpArrivalnoticeDto
	r.Set("erp_arrivalnotice_dto", _erpArrivalnoticeDto)
	return nil
}

// GetErpArrivalnoticeDto ErpArrivalnoticeDto Getter
func (r AlibabaWdkUmsInboundAPIRequest) GetErpArrivalnoticeDto() *ErpArrivalNoticeDto {
	return r._erpArrivalnoticeDto
}
