package tmallgeniescp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse 10-同步库存现有量 API返回值
// alibaba.tmallgenie.scp.plan.inventor.qty.get
//
// 同步库存现有量
type AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanInventorQtyGetAPIResponseModel
}

// AlibabaTmallgenieScpPlanInventorQtyGetAPIResponseModel is 10-同步库存现有量 成功返回结果
type AlibabaTmallgenieScpPlanInventorQtyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_inventor_qty_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
