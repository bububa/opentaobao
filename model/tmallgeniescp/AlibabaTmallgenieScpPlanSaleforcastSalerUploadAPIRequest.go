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

// New
