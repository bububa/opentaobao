package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsinstanttracesearchAPIResponse 物流详情查询 API返回值
// taobao.logistics.instant.trace.search
//
// 物流详情查询
type TaobaologisticsinstanttracesearchAPIResponse struct {
	model.CommonResponse
	TaobaologisticsinstanttracesearchAPIResponseModel
}

// TaobaologisticsinstanttracesearchAPIResponseModel is 物流详情查询 成功返回结果
type TaobaologisticsinstanttracesearchAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_instant_trace_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaologisticsinstanttracesearchResult `json:"result,omitempty" xml:"result,omitempty"`
}
