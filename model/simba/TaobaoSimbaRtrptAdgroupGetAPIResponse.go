package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptAdgroupGetAPIResponse 获取推广组实时报表数据 API返回值
// taobao.simba.rtrpt.adgroup.get
//
// 获取推广组实时报表数据
type TaobaoSimbaRtrptAdgroupGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRtrptAdgroupGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptAdgroupGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRtrptAdgroupGetAPIResponseModel).Reset()
}

// TaobaoSimbaRtrptAdgroupGetAPIResponseModel is 获取推广组实时报表数据 成功返回结果
type TaobaoSimbaRtrptAdgroupGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rtrpt_adgroup_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1111
	Results []RtRptResultEntityDto `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptAdgroupGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoSimbaRtrptAdgroupGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRtrptAdgroupGetAPIResponse)
	},
}

// GetTaobaoSimbaRtrptAdgroupGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRtrptAdgroupGetAPIResponse
func GetTaobaoSimbaRtrptAdgroupGetAPIResponse() *TaobaoSimbaRtrptAdgroupGetAPIResponse {
	return poolTaobaoSimbaRtrptAdgroupGetAPIResponse.Get().(*TaobaoSimbaRtrptAdgroupGetAPIResponse)
}

// ReleaseTaobaoSimbaRtrptAdgroupGetAPIResponse 将 TaobaoSimbaRtrptAdgroupGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRtrptAdgroupGetAPIResponse(v *TaobaoSimbaRtrptAdgroupGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRtrptAdgroupGetAPIResponse.Put(v)
}
