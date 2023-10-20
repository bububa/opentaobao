package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanBomUploadAPIResponse 计划BOM同步 API返回值
// alibaba.tmallgenie.scp.plan.bom.upload
//
// 计划BOM同步
type AlibabaTmallgenieScpPlanBomUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanBomUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanBomUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanBomUploadAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanBomUploadAPIResponseModel is 计划BOM同步 成功返回结果
type AlibabaTmallgenieScpPlanBomUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_bom_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 对象列表
	DataList []BomDto `json:"data_list,omitempty" xml:"data_list>bom_dto,omitempty"`
	// 返回描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 请求唯一ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanBomUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DataList = m.DataList[:0]
	m.ResultMsg = ""
	m.TraceId = ""
	m.ResultCode = ""
}

var poolAlibabaTmallgenieScpPlanBomUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanBomUploadAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanBomUploadAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanBomUploadAPIResponse
func GetAlibabaTmallgenieScpPlanBomUploadAPIResponse() *AlibabaTmallgenieScpPlanBomUploadAPIResponse {
	return poolAlibabaTmallgenieScpPlanBomUploadAPIResponse.Get().(*AlibabaTmallgenieScpPlanBomUploadAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanBomUploadAPIResponse 将 AlibabaTmallgenieScpPlanBomUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanBomUploadAPIResponse(v *AlibabaTmallgenieScpPlanBomUploadAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanBomUploadAPIResponse.Put(v)
}
