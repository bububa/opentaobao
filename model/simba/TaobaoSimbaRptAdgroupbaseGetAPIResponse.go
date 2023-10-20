package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptAdgroupbaseGetAPIResponse 推广组基础报表数据对象 API返回值
// taobao.simba.rpt.adgroupbase.get
//
// 推广组基础报表数据对象
type TaobaoSimbaRptAdgroupbaseGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptAdgroupbaseGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRptAdgroupbaseGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRptAdgroupbaseGetAPIResponseModel).Reset()
}

// TaobaoSimbaRptAdgroupbaseGetAPIResponseModel is 推广组基础报表数据对象 成功返回结果
type TaobaoSimbaRptAdgroupbaseGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_adgroupbase_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 广告组基础数据对象
	RptAdgroupBaseList string `json:"rpt_adgroup_base_list,omitempty" xml:"rpt_adgroup_base_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRptAdgroupbaseGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RptAdgroupBaseList = ""
}

var poolTaobaoSimbaRptAdgroupbaseGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRptAdgroupbaseGetAPIResponse)
	},
}

// GetTaobaoSimbaRptAdgroupbaseGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRptAdgroupbaseGetAPIResponse
func GetTaobaoSimbaRptAdgroupbaseGetAPIResponse() *TaobaoSimbaRptAdgroupbaseGetAPIResponse {
	return poolTaobaoSimbaRptAdgroupbaseGetAPIResponse.Get().(*TaobaoSimbaRptAdgroupbaseGetAPIResponse)
}

// ReleaseTaobaoSimbaRptAdgroupbaseGetAPIResponse 将 TaobaoSimbaRptAdgroupbaseGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRptAdgroupbaseGetAPIResponse(v *TaobaoSimbaRptAdgroupbaseGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRptAdgroupbaseGetAPIResponse.Put(v)
}
