package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasellervendororderlistAPIRequest 国际站服务市场订单列表接口 API请求
// alibaba.seller.vendor.order.list
//
// 返回服务商在服务市场的客户订单
type AlibabasellervendororderlistAPIRequest struct {
	model.Params
	// 查询参数
	_queryTradeDto *QueryTradeDto
}

// NewAlibabasellervendororderlistRequest 初始化AlibabasellervendororderlistAPIRequest对象
func NewAlibabasellervendororderlistRequest() *AlibabasellervendororderlistAPIRequest {
	return &AlibabasellervendororderlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasellervendororderlistAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.vendor.order.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasellervendororderlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasellervendororderlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryTradeDto is QueryTradeDto Setter
// 查询参数
func (r *AlibabasellervendororderlistAPIRequest) SetQueryTradeDto(_queryTradeDto *QueryTradeDto) error {
	r._queryTradeDto = _queryTradeDto
	r.Set("query_trade_dto", _queryTradeDto)
	return nil
}

// GetQueryTradeDto QueryTradeDto Getter
func (r AlibabasellervendororderlistAPIRequest) GetQueryTradeDto() *QueryTradeDto {
	return r._queryTradeDto
}
