package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest 订单详细信息(面单及仓库信息) API请求
// alibaba.onetouch.logistics.express.order.detail.get
//
// 订单详细信息(面单及仓库信息)
type AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *LogisticsOrderQueryDto
}

// NewAlibabaOnetouchLogisticsExpressOrderDetailGetRequest 初始化AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressOrderDetailGetRequest() *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest {
	return &AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.order.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamQuery is ParamQuery Setter
// 请求参数
func (r *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) SetParamQuery(_paramQuery *LogisticsOrderQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// GetParamQuery ParamQuery Getter
func (r AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest) GetParamQuery() *LogisticsOrderQueryDto {
	return r._paramQuery
}
