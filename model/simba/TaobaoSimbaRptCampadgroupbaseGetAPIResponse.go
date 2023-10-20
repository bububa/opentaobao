package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptCampadgroupbaseGetAPIResponse 推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型) API返回值
// taobao.simba.rpt.campadgroupbase.get
//
// 推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型)
type TaobaoSimbaRptCampadgroupbaseGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptCampadgroupbaseGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRptCampadgroupbaseGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRptCampadgroupbaseGetAPIResponseModel).Reset()
}

// TaobaoSimbaRptCampadgroupbaseGetAPIResponseModel is 推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型) 成功返回结果
type TaobaoSimbaRptCampadgroupbaseGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_campadgroupbase_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广计划下推广组的基础数据列表
	RptCampadgroupBaseList string `json:"rpt_campadgroup_base_list,omitempty" xml:"rpt_campadgroup_base_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRptCampadgroupbaseGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RptCampadgroupBaseList = ""
}

var poolTaobaoSimbaRptCampadgroupbaseGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRptCampadgroupbaseGetAPIResponse)
	},
}

// GetTaobaoSimbaRptCampadgroupbaseGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRptCampadgroupbaseGetAPIResponse
func GetTaobaoSimbaRptCampadgroupbaseGetAPIResponse() *TaobaoSimbaRptCampadgroupbaseGetAPIResponse {
	return poolTaobaoSimbaRptCampadgroupbaseGetAPIResponse.Get().(*TaobaoSimbaRptCampadgroupbaseGetAPIResponse)
}

// ReleaseTaobaoSimbaRptCampadgroupbaseGetAPIResponse 将 TaobaoSimbaRptCampadgroupbaseGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRptCampadgroupbaseGetAPIResponse(v *TaobaoSimbaRptCampadgroupbaseGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRptCampadgroupbaseGetAPIResponse.Put(v)
}
