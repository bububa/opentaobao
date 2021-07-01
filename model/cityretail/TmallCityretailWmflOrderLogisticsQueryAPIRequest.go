package cityretail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCityretailWmflOrderLogisticsQueryAPIRequest
完美履约订单物流状态查询接口 API请求
tmall.cityretail.wmfl.order.logistics.query

完美履约订单物流状态查询接口，该接口只能查询未完结的履约单以及完结的3天内订单 */
type TmallCityretailWmflOrderLogisticsQueryAPIRequest struct {
	model.Params
	// 订单号
	_mainOrderId string
}

// NewTmallCityretailWmflOrderLogisticsQueryRequest 初始化TmallCityretailWmflOrderLogisticsQueryAPIRequest对象
func NewTmallCityretailWmflOrderLogisticsQueryRequest() *TmallCityretailWmflOrderLogisticsQueryAPIRequest {
	return &TmallCityretailWmflOrderLogisticsQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCityretailWmflOrderLogisticsQueryAPIRequest) GetApiMethodName() string {
	return "tmall.cityretail.wmfl.order.logistics.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCityretailWmflOrderLogisticsQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MainOrderId Setter
// 订单号
func (r *TmallCityretailWmflOrderLogisticsQueryAPIRequest) SetMainOrderId(_mainOrderId string) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// Get MainOrderId Getter
func (r TmallCityretailWmflOrderLogisticsQueryAPIRequest) GetMainOrderId() string {
	return r._mainOrderId
}
