package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse 11-同步本周的po单（从W-1周到W+4周） API返回值
// alibaba.tmallgenie.scp.plan.current.po.get
//
// 11-同步本周的po单（从W-1周到W+4周）
type AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanCurrentPoGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanCurrentPoGetAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanCurrentPoGetAPIResponseModel is 11-同步本周的po单（从W-1周到W+4周） 成功返回结果
type AlibabaTmallgenieScpPlanCurrentPoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_current_po_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanCurrentPoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTmallgenieScpPlanCurrentPoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanCurrentPoGetAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse
func GetAlibabaTmallgenieScpPlanCurrentPoGetAPIResponse() *AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse {
	return poolAlibabaTmallgenieScpPlanCurrentPoGetAPIResponse.Get().(*AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanCurrentPoGetAPIResponse 将 AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanCurrentPoGetAPIResponse(v *AlibabaTmallgenieScpPlanCurrentPoGetAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanCurrentPoGetAPIResponse.Put(v)
}
