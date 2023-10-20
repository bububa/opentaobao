package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumsinboundAPIRequest 入库-ERP下发单 API请求
// alibaba.wdk.ums.inbound
//
// 入库-ERP下发单
type AlibabawdkumsinboundAPIRequest struct {
	model.Params
	// 1
	_erpArrivalnoticeDto *ErpArrivalNoticeDto
}

// NewAlibabawdkumsinboundRequest 初始化AlibabawdkumsinboundAPIRequest对象
func NewAlibabawdkumsinboundRequest() *AlibabawdkumsinboundAPIRequest {
	return &AlibabawdkumsinboundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkumsinboundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.inbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkumsinboundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkumsinboundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErpArrivalnoticeDto is ErpArrivalnoticeDto Setter
// 1
func (r *AlibabawdkumsinboundAPIRequest) SetErpArrivalnoticeDto(_erpArrivalnoticeDto *ErpArrivalNoticeDto) error {
	r._erpArrivalnoticeDto = _erpArrivalnoticeDto
	r.Set("erp_arrivalnotice_dto", _erpArrivalnoticeDto)
	return nil
}

// GetErpArrivalnoticeDto ErpArrivalnoticeDto Getter
func (r AlibabawdkumsinboundAPIRequest) GetErpArrivalnoticeDto() *ErpArrivalNoticeDto {
	return r._erpArrivalnoticeDto
}
