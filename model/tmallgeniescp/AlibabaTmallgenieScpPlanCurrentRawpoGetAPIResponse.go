package tmallgeniescp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanCurrentRawpoGetAPIResponse 二级物料-PO数据同步 API返回值
// alibaba.tmallgenie.scp.plan.current.rawpo.get
//
// 二级物料-PO数据同步（WO-W[TL])
type AlibabaTmallgenieScpPlanCurrentRawpoGetAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanCurrentRawpoGetAPIResponseModel
}

// AlibabaTmallgenieScpPlanCurrentRawpoGetAPIResponseModel is 二级物料-PO数据同步 成功返回结果
type AlibabaTmallgenieScpPlanCurrentRawpoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_current_rawpo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
