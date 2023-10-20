package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptTargetingtagGetAPIResponse 搜索人群离线报表 API返回值
// taobao.simba.rpt.targetingtag.get
//
// 获取搜搜人群实时报表
type TaobaoSimbaRptTargetingtagGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptTargetingtagGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRptTargetingtagGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRptTargetingtagGetAPIResponseModel).Reset()
}

// TaobaoSimbaRptTargetingtagGetAPIResponseModel is 搜索人群离线报表 成功返回结果
type TaobaoSimbaRptTargetingtagGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_targetingtag_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 111
	Results []RtRptResultEntityDto `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRptTargetingtagGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoSimbaRptTargetingtagGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRptTargetingtagGetAPIResponse)
	},
}

// GetTaobaoSimbaRptTargetingtagGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRptTargetingtagGetAPIResponse
func GetTaobaoSimbaRptTargetingtagGetAPIResponse() *TaobaoSimbaRptTargetingtagGetAPIResponse {
	return poolTaobaoSimbaRptTargetingtagGetAPIResponse.Get().(*TaobaoSimbaRptTargetingtagGetAPIResponse)
}

// ReleaseTaobaoSimbaRptTargetingtagGetAPIResponse 将 TaobaoSimbaRptTargetingtagGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRptTargetingtagGetAPIResponse(v *TaobaoSimbaRptTargetingtagGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRptTargetingtagGetAPIResponse.Put(v)
}
