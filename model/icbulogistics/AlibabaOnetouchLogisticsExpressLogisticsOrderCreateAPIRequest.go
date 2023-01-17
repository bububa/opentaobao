package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest 快递下单 API请求
// alibaba.onetouch.logistics.express.logistics.order.create
//
// 快递下单
type AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest struct {
	model.Params
	// 请求参数对象
	_paramnQuery *PlaceOrderDto
}

// NewAlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest 初始化AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressLogisticsOrderCreateRequest() *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest {
	return &AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.logistics.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamnQuery is ParamnQuery Setter
// 请求参数对象
func (r *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest) SetParamnQuery(_paramnQuery *PlaceOrderDto) error {
	r._paramnQuery = _paramnQuery
	r.Set("paramn_query", _paramnQuery)
	return nil
}

// GetParamnQuery ParamnQuery Getter
func (r AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest) GetParamnQuery() *PlaceOrderDto {
	return r._paramnQuery
}
