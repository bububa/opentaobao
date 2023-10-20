package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanChannelGetAPIResponse 5-IBP同步渠道接口 API返回值
// alibaba.tmallgenie.scp.plan.channel.get
//
// IBP同步渠道接口
type AlibabaTmallgenieScpPlanChannelGetAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanChannelGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanChannelGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanChannelGetAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanChannelGetAPIResponseModel is 5-IBP同步渠道接口 成功返回结果
type AlibabaTmallgenieScpPlanChannelGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_channel_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanChannelGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTmallgenieScpPlanChannelGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanChannelGetAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanChannelGetAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanChannelGetAPIResponse
func GetAlibabaTmallgenieScpPlanChannelGetAPIResponse() *AlibabaTmallgenieScpPlanChannelGetAPIResponse {
	return poolAlibabaTmallgenieScpPlanChannelGetAPIResponse.Get().(*AlibabaTmallgenieScpPlanChannelGetAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanChannelGetAPIResponse 将 AlibabaTmallgenieScpPlanChannelGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanChannelGetAPIResponse(v *AlibabaTmallgenieScpPlanChannelGetAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanChannelGetAPIResponse.Put(v)
}
