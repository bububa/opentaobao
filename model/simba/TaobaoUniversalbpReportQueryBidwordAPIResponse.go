package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportquerybidwordAPIResponse 关键词报表查询 API返回值
// taobao.universalbp.report.query.bidword
//
// 关键词报表查询
type TaobaouniversalbpreportquerybidwordAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpreportquerybidwordAPIResponseModel
}

// TaobaouniversalbpreportquerybidwordAPIResponseModel is 关键词报表查询 成功返回结果
type TaobaouniversalbpreportquerybidwordAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_bidword_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpreportquerybidwordTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
