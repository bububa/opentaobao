package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest
19-销售预测数量（销管）回传接口 API请求
alibaba.tmallgenie.scp.plan.saleforcast.saler.upload

销售预测数量（销管）回传接口 */
type AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest struct {
	model.Params
	// 入参
	_salesForecastRequest *SalesForecastRequest
}

// NewAlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest 初始化AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest() *AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.saleforcast.saler.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SalesForecastRequest Setter
// 入参
func (r *AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest) SetSalesForecastRequest(_salesForecastRequest *SalesForecastRequest) error {
	r._salesForecastRequest = _salesForecastRequest
	r.Set("sales_forecast_request", _salesForecastRequest)
	return nil
}

// Get SalesForecastRequest Getter
func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest) GetSalesForecastRequest() *SalesForecastRequest {
	return r._salesForecastRequest
}
