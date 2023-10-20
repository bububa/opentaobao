package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportquerycreativeAPIResponse 创意报表查询 API返回值
// taobao.universalbp.report.query.creative
//
// 创意报表查询
type TaobaouniversalbpreportquerycreativeAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpreportquerycreativeAPIResponseModel
}

// TaobaouniversalbpreportquerycreativeAPIResponseModel is 创意报表查询 成功返回结果
type TaobaouniversalbpreportquerycreativeAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_creative_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpreportquerycreativeTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
