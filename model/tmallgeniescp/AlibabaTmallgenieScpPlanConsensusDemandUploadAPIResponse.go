package tmallgeniescp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanconsensusdemanduploadAPIResponse 20-IBP共识需求回传接口 API返回值
// alibaba.tmallgenie.scp.plan.consensus.demand.upload
//
// IBP共识需求回传接口
type AlibabatmallgeniescpplanconsensusdemanduploadAPIResponse struct {
	model.CommonResponse
	AlibabatmallgeniescpplanconsensusdemanduploadAPIResponseModel
}

// AlibabatmallgeniescpplanconsensusdemanduploadAPIResponseModel is 20-IBP共识需求回传接口 成功返回结果
type AlibabatmallgeniescpplanconsensusdemanduploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_consensus_demand_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabatmallgeniescpplanconsensusdemanduploadResult `json:"result,omitempty" xml:"result,omitempty"`
}
