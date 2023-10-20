package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse 推广组下创意报表基础数据查询(汇总数据，不分类型) API返回值
// taobao.simba.rpt.adgroupcreativebase.get
//
// 推广组下创意报表基础数据查询(汇总数据，不分类型)
type TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptAdgroupcreativebaseGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRptAdgroupcreativebaseGetAPIResponseModel).Reset()
}

// TaobaoSimbaRptAdgroupcreativebaseGetAPIResponseModel is 推广组下创意报表基础数据查询(汇总数据，不分类型) 成功返回结果
type TaobaoSimbaRptAdgroupcreativebaseGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_adgroupcreativebase_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组下的创意基础数据列表
	RptAdgroupcreativeBaseList string `json:"rpt_adgroupcreative_base_list,omitempty" xml:"rpt_adgroupcreative_base_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRptAdgroupcreativebaseGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RptAdgroupcreativeBaseList = ""
}

var poolTaobaoSimbaRptAdgroupcreativebaseGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse)
	},
}

// GetTaobaoSimbaRptAdgroupcreativebaseGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse
func GetTaobaoSimbaRptAdgroupcreativebaseGetAPIResponse() *TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse {
	return poolTaobaoSimbaRptAdgroupcreativebaseGetAPIResponse.Get().(*TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse)
}

// ReleaseTaobaoSimbaRptAdgroupcreativebaseGetAPIResponse 将 TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRptAdgroupcreativebaseGetAPIResponse(v *TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRptAdgroupcreativebaseGetAPIResponse.Put(v)
}
