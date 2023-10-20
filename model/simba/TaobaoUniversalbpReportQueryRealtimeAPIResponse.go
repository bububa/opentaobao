package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportqueryrealtimeAPIResponse 实时报表查询 API返回值
// taobao.universalbp.report.query.realtime
//
// 实时报表查询
type TaobaouniversalbpreportqueryrealtimeAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpreportqueryrealtimeAPIResponseModel
}

// TaobaouniversalbpreportqueryrealtimeAPIResponseModel is 实时报表查询 成功返回结果
type TaobaouniversalbpreportqueryrealtimeAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_realtime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpreportqueryrealtimeTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
