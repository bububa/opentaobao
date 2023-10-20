package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest 19-销售预测数量（销管）回传接口 API请求
// alibaba.tmallgenie.scp.plan.saleforcast.saler.upload
//
// 销售预测数量（销管）回传接口
type AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest struct {
	model.Params
	// 入参
	_salesForecastRequest *SalesForecastRequest
}

// NewAlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest 初始化AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest() *AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest) Reset() {
	r._salesForecastRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.saleforcast.saler.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSalesForecastRequest is SalesForecastRequest Setter
// 入参
func (r *AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest) SetSalesForecastRequest(_salesForecastRequest *SalesForecastRequest) error {
	r._salesForecastRequest = _salesForecastRequest
	r.Set("sales_forecast_request", _salesForecastRequest)
	return nil
}

// GetSalesForecastRequest SalesForecastRequest Getter
func (r AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest) GetSalesForecastRequest() *SalesForecastRequest {
	return r._salesForecastRequest
}

var poolAlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanSaleforcastSalerUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest
func GetAlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest() *AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest 将 AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest(v *AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIRequest.Put(v)
}
