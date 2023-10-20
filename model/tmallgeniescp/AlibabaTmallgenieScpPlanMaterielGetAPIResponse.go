package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanMaterielGetAPIResponse 1-IBP同步物料接口 API返回值
// alibaba.tmallgenie.scp.plan.materiel.get
//
// IBP同步物料接口
type AlibabaTmallgenieScpPlanMaterielGetAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanMaterielGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanMaterielGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanMaterielGetAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanMaterielGetAPIResponseModel is 1-IBP同步物料接口 成功返回结果
type AlibabaTmallgenieScpPlanMaterielGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_materiel_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanMaterielGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTmallgenieScpPlanMaterielGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanMaterielGetAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanMaterielGetAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanMaterielGetAPIResponse
func GetAlibabaTmallgenieScpPlanMaterielGetAPIResponse() *AlibabaTmallgenieScpPlanMaterielGetAPIResponse {
	return poolAlibabaTmallgenieScpPlanMaterielGetAPIResponse.Get().(*AlibabaTmallgenieScpPlanMaterielGetAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanMaterielGetAPIResponse 将 AlibabaTmallgenieScpPlanMaterielGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanMaterielGetAPIResponse(v *AlibabaTmallgenieScpPlanMaterielGetAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanMaterielGetAPIResponse.Put(v)
}
