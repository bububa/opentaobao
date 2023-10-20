package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptTargetingtageffectGetAPIResponse 获取定向效果报表数据 API返回值
// taobao.simba.rpt.targetingtageffect.get
//
// 获取定向效果报表数据
type TaobaoSimbaRptTargetingtageffectGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptTargetingtageffectGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRptTargetingtageffectGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRptTargetingtageffectGetAPIResponseModel).Reset()
}

// TaobaoSimbaRptTargetingtageffectGetAPIResponseModel is 获取定向效果报表数据 成功返回结果
type TaobaoSimbaRptTargetingtageffectGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_targetingtageffect_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 效果数据
	Results []RptEffectEntityDto `json:"results,omitempty" xml:"results>rpt_effect_entity_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRptTargetingtageffectGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoSimbaRptTargetingtageffectGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRptTargetingtageffectGetAPIResponse)
	},
}

// GetTaobaoSimbaRptTargetingtageffectGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRptTargetingtageffectGetAPIResponse
func GetTaobaoSimbaRptTargetingtageffectGetAPIResponse() *TaobaoSimbaRptTargetingtageffectGetAPIResponse {
	return poolTaobaoSimbaRptTargetingtageffectGetAPIResponse.Get().(*TaobaoSimbaRptTargetingtageffectGetAPIResponse)
}

// ReleaseTaobaoSimbaRptTargetingtageffectGetAPIResponse 将 TaobaoSimbaRptTargetingtageffectGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRptTargetingtageffectGetAPIResponse(v *TaobaoSimbaRptTargetingtageffectGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRptTargetingtageffectGetAPIResponse.Put(v)
}
