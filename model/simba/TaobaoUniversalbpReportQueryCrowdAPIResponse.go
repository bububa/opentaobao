package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportquerycrowdAPIResponse 人群报表查询 API返回值
// taobao.universalbp.report.query.crowd
//
// 人群报表查询
type TaobaouniversalbpreportquerycrowdAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpreportquerycrowdAPIResponseModel
}

// TaobaouniversalbpreportquerycrowdAPIResponseModel is 人群报表查询 成功返回结果
type TaobaouniversalbpreportquerycrowdAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_crowd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpreportquerycrowdTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
