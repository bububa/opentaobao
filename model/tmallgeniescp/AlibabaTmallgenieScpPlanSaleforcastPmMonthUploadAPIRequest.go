package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIRequest 24-销售月预测数量（产管）回传-月度 API请求
// alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload
//
// 销售月预测数量（产管）回传-月度
type AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIRequest struct {
	model.Params
	// 入参
	_salesForecastRequest *SalesForecastRequest
}

// NewAlibabatmallgeniescpplansaleforcastpmmonthuploadRequest 初始化AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIRequest对象
func NewAlibabatmallgeniescpplansaleforcastpmmonthuploadRequest() *AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIRequest {
	return &AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSalesForecastRequest is SalesForecastRequest Setter
// 入参
func (r *AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIRequest) SetSalesForecastRequest(_salesForecastRequest *SalesForecastRequest) error {
	r._salesForecastRequest = _salesForecastRequest
	r.Set("sales_forecast_request", _salesForecastRequest)
	return nil
}

// GetSalesForecastRequest SalesForecastRequest Getter
func (r AlibabatmallgeniescpplansaleforcastpmmonthuploadAPIRequest) GetSalesForecastRequest() *SalesForecastRequest {
	return r._salesForecastRequest
}
