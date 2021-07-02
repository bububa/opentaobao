package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest 计算快递运费&下单参数校验 API请求
// alibaba.onetouch.logistics.express.charge.calculate
//
// 计算快递运费、下单参数校验
type AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest struct {
	model.Params
	// 请求参数对象
	_paramnQuery *PlaceOrderDto
}

// NewAlibabaOnetouchLogisticsExpressChargeCalculateRequest 初始化AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressChargeCalculateRequest() *AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest {
	return &AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.charge.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamnQuery is ParamnQuery Setter
// 请求参数对象
func (r *AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) SetParamnQuery(_paramnQuery *PlaceOrderDto) error {
	r._paramnQuery = _paramnQuery
	r.Set("paramn_query", _paramnQuery)
	return nil
}

// GetParamnQuery ParamnQuery Getter
func (r AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest) GetParamnQuery() *PlaceOrderDto {
	return r._paramnQuery
}
