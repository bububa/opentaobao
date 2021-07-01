package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest
24-销售月预测数量（产管）回传-月度 API请求
alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload

销售月预测数量（产管）回传-月度 */
type AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIRequest struct {
	model.Params
	// 入参
	_salesForecastRequest *SalesForecastRequest
}

// New
