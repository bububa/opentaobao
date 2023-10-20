package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCrowdofflinenewFindAPIResponse 获取人群离线多日汇总报表 API返回值
// taobao.subway.crowdofflinenew.find
//
// 获取人群离线报表
type TaobaoSubwayCrowdofflinenewFindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayCrowdofflinenewFindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayCrowdofflinenewFindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayCrowdofflinenewFindAPIResponseModel).Reset()
}

// TaobaoSubwayCrowdofflinenewFindAPIResponseModel is 获取人群离线多日汇总报表 成功返回结果
type TaobaoSubwayCrowdofflinenewFindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_crowdofflinenew_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result []ReportResultTopDto `json:"result,omitempty" xml:"result>report_result_top_dto,omitempty"`
	// 提示
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubwayCrowdofflinenewFindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.Message = ""
	m.TotalCount = 0
}

var poolTaobaoSubwayCrowdofflinenewFindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayCrowdofflinenewFindAPIResponse)
	},
}

// GetTaobaoSubwayCrowdofflinenewFindAPIResponse 从 sync.Pool 获取 TaobaoSubwayCrowdofflinenewFindAPIResponse
func GetTaobaoSubwayCrowdofflinenewFindAPIResponse() *TaobaoSubwayCrowdofflinenewFindAPIResponse {
	return poolTaobaoSubwayCrowdofflinenewFindAPIResponse.Get().(*TaobaoSubwayCrowdofflinenewFindAPIResponse)
}

// ReleaseTaobaoSubwayCrowdofflinenewFindAPIResponse 将 TaobaoSubwayCrowdofflinenewFindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayCrowdofflinenewFindAPIResponse(v *TaobaoSubwayCrowdofflinenewFindAPIResponse) {
	v.Reset()
	poolTaobaoSubwayCrowdofflinenewFindAPIResponse.Put(v)
}
