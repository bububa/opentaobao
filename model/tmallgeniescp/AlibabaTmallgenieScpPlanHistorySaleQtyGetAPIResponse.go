package tmallgeniescp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIResponse
【已废除】同步历史的销售数据 API返回值
alibaba.tmallgenie.scp.plan.history.sale.qty.get

同步历史的销售数据 */
type AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIResponseModel
}

// AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIResponseModel is 【已废除】同步历史的销售数据 成功返回结果
type AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_history_sale_qty_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
