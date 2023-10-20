package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse 推广组下的创意报表效果数据查询(汇总数据，不分类型) API返回值
// taobao.simba.rpt.adgroupcreativeeffect.get
//
// 推广组下的创意报表效果数据查询(汇总数据，不分类型)
type TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponseModel).Reset()
}

// TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponseModel is 推广组下的创意报表效果数据查询(汇总数据，不分类型) 成功返回结果
type TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_adgroupcreativeeffect_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组下的创意效果数据列表
	RptAdgroupcreativeEffectList string `json:"rpt_adgroupcreative_effect_list,omitempty" xml:"rpt_adgroupcreative_effect_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RptAdgroupcreativeEffectList = ""
}

var poolTaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse)
	},
}

// GetTaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse
func GetTaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse() *TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse {
	return poolTaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse.Get().(*TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse)
}

// ReleaseTaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse 将 TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse(v *TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse.Put(v)
}
