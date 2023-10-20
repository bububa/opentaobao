package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptTargetingtagbaseGetAPIResponse 定向基础报表 API返回值
// taobao.simba.rpt.targetingtagbase.get
//
// 获取定向基础报表
type TaobaoSimbaRptTargetingtagbaseGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptTargetingtagbaseGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRptTargetingtagbaseGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRptTargetingtagbaseGetAPIResponseModel).Reset()
}

// TaobaoSimbaRptTargetingtagbaseGetAPIResponseModel is 定向基础报表 成功返回结果
type TaobaoSimbaRptTargetingtagbaseGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_targetingtagbase_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Results []RptBaseEntityDto `json:"results,omitempty" xml:"results>rpt_base_entity_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRptTargetingtagbaseGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoSimbaRptTargetingtagbaseGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRptTargetingtagbaseGetAPIResponse)
	},
}

// GetTaobaoSimbaRptTargetingtagbaseGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRptTargetingtagbaseGetAPIResponse
func GetTaobaoSimbaRptTargetingtagbaseGetAPIResponse() *TaobaoSimbaRptTargetingtagbaseGetAPIResponse {
	return poolTaobaoSimbaRptTargetingtagbaseGetAPIResponse.Get().(*TaobaoSimbaRptTargetingtagbaseGetAPIResponse)
}

// ReleaseTaobaoSimbaRptTargetingtagbaseGetAPIResponse 将 TaobaoSimbaRptTargetingtagbaseGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRptTargetingtagbaseGetAPIResponse(v *TaobaoSimbaRptTargetingtagbaseGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRptTargetingtagbaseGetAPIResponse.Put(v)
}
