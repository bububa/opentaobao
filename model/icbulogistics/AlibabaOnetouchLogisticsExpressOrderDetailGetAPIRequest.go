package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpressorderdetailgetAPIRequest 订单详细信息(面单及仓库信息) API请求
// alibaba.onetouch.logistics.express.order.detail.get
//
// 订单详细信息(面单及仓库信息)
type AlibabaonetouchlogisticsexpressorderdetailgetAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *LogisticsOrderQueryDto
}

// NewAlibabaonetouchlogisticsexpressorderdetailgetRequest 初始化AlibabaonetouchlogisticsexpressorderdetailgetAPIRequest对象
func NewAlibabaonetouchlogisticsexpressorderdetailgetRequest() *AlibabaonetouchlogisticsexpressorderdetailgetAPIRequest {
	return &AlibabaonetouchlogisticsexpressorderdetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaonetouchlogisticsexpressorderdetailgetAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.order.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaonetouchlogisticsexpressorderdetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaonetouchlogisticsexpressorderdetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQuery is ParamQuery Setter
// 请求参数
func (r *AlibabaonetouchlogisticsexpressorderdetailgetAPIRequest) SetParamQuery(_paramQuery *LogisticsOrderQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// GetParamQuery ParamQuery Getter
func (r AlibabaonetouchlogisticsexpressorderdetailgetAPIRequest) GetParamQuery() *LogisticsOrderQueryDto {
	return r._paramQuery
}
