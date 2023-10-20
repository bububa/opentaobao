package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse 推广组下的词基础报表数据查询(明细数据不分类型查询) API返回值
// taobao.simba.rpt.adgroupkeywordbase.get
//
// 推广组下的词基础报表数据查询(明细数据不分类型查询)
type TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponseModel).Reset()
}

// TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponseModel is 推广组下的词基础报表数据查询(明细数据不分类型查询) 成功返回结果
type TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_adgroupkeywordbase_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 词基础数据返回结果
	RptAdgroupkeywordBaseList string `json:"rpt_adgroupkeyword_base_list,omitempty" xml:"rpt_adgroupkeyword_base_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RptAdgroupkeywordBaseList = ""
}

var poolTaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse)
	},
}

// GetTaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse
func GetTaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse() *TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse {
	return poolTaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse.Get().(*TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse)
}

// ReleaseTaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse 将 TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse(v *TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse.Put(v)
}
