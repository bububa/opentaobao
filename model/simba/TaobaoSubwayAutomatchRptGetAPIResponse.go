package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayAutomatchRptGetAPIResponse 查询流量智选天级报告 API返回值
// taobao.subway.automatch.rpt.get
//
// 查询流量智选天级报告
type TaobaoSubwayAutomatchRptGetAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayAutomatchRptGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayAutomatchRptGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayAutomatchRptGetAPIResponseModel).Reset()
}

// TaobaoSubwayAutomatchRptGetAPIResponseModel is 查询流量智选天级报告 成功返回结果
type TaobaoSubwayAutomatchRptGetAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_automatch_rpt_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 流量智选天级别报表数据
	ResultList []ResultMap `json:"result_list,omitempty" xml:"result_list>result_map,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubwayAutomatchRptGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolTaobaoSubwayAutomatchRptGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayAutomatchRptGetAPIResponse)
	},
}

// GetTaobaoSubwayAutomatchRptGetAPIResponse 从 sync.Pool 获取 TaobaoSubwayAutomatchRptGetAPIResponse
func GetTaobaoSubwayAutomatchRptGetAPIResponse() *TaobaoSubwayAutomatchRptGetAPIResponse {
	return poolTaobaoSubwayAutomatchRptGetAPIResponse.Get().(*TaobaoSubwayAutomatchRptGetAPIResponse)
}

// ReleaseTaobaoSubwayAutomatchRptGetAPIResponse 将 TaobaoSubwayAutomatchRptGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayAutomatchRptGetAPIResponse(v *TaobaoSubwayAutomatchRptGetAPIResponse) {
	v.Reset()
	poolTaobaoSubwayAutomatchRptGetAPIResponse.Put(v)
}
