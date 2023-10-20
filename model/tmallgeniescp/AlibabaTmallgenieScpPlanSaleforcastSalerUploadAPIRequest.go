package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplansaleforcastsaleruploadAPIRequest 19-销售预测数量（销管）回传接口 API请求
// alibaba.tmallgenie.scp.plan.saleforcast.saler.upload
//
// 销售预测数量（销管）回传接口
type AlibabatmallgeniescpplansaleforcastsaleruploadAPIRequest struct {
	model.Params
	// 入参
	_salesForecastRequest *SalesForecastRequest
}

// NewAlibabatmallgeniescpplansaleforcastsaleruploadRequest 初始化AlibabatmallgeniescpplansaleforcastsaleruploadAPIRequest对象
func NewAlibabatmallgeniescpplansaleforcastsaleruploadRequest() *AlibabatmallgeniescpplansaleforcastsaleruploadAPIRequest {
	return &AlibabatmallgeniescpplansaleforcastsaleruploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplansaleforcastsaleruploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.saleforcast.saler.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplansaleforcastsaleruploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplansaleforcastsaleruploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSalesForecastRequest is SalesForecastRequest Setter
// 入参
func (r *AlibabatmallgeniescpplansaleforcastsaleruploadAPIRequest) SetSalesForecastRequest(_salesForecastRequest *SalesForecastRequest) error {
	r._salesForecastRequest = _salesForecastRequest
	r.Set("sales_forecast_request", _salesForecastRequest)
	return nil
}

// GetSalesForecastRequest SalesForecastRequest Getter
func (r AlibabatmallgeniescpplansaleforcastsaleruploadAPIRequest) GetSalesForecastRequest() *SalesForecastRequest {
	return r._salesForecastRequest
}
