package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpresschargecalculateAPIRequest 计算快递运费&下单参数校验 API请求
// alibaba.onetouch.logistics.express.charge.calculate
//
// 计算快递运费、下单参数校验
type AlibabaonetouchlogisticsexpresschargecalculateAPIRequest struct {
	model.Params
	// 请求参数对象
	_paramnQuery *PlaceOrderDto
}

// NewAlibabaonetouchlogisticsexpresschargecalculateRequest 初始化AlibabaonetouchlogisticsexpresschargecalculateAPIRequest对象
func NewAlibabaonetouchlogisticsexpresschargecalculateRequest() *AlibabaonetouchlogisticsexpresschargecalculateAPIRequest {
	return &AlibabaonetouchlogisticsexpresschargecalculateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaonetouchlogisticsexpresschargecalculateAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.charge.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaonetouchlogisticsexpresschargecalculateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaonetouchlogisticsexpresschargecalculateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamnQuery is ParamnQuery Setter
// 请求参数对象
func (r *AlibabaonetouchlogisticsexpresschargecalculateAPIRequest) SetParamnQuery(_paramnQuery *PlaceOrderDto) error {
	r._paramnQuery = _paramnQuery
	r.Set("paramn_query", _paramnQuery)
	return nil
}

// GetParamnQuery ParamnQuery Getter
func (r AlibabaonetouchlogisticsexpresschargecalculateAPIRequest) GetParamnQuery() *PlaceOrderDto {
	return r._paramnQuery
}
