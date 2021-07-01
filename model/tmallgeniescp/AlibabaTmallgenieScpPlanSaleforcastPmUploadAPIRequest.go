package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest
18-销售预测数量（产管）回传接口 API请求
alibaba.tmallgenie.scp.plan.saleforcast.pm.upload

销售预测数量（产管）回传接口 */
type AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIRequest struct {
	model.Params
	// 入参
	_salesForecastRequest *SalesForecastRequest
}

// New
