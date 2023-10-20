package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpreportqueryaccountAPIResponse 账户报表查询 API返回值
// taobao.universalbp.report.query.account
//
// 账户报表查询
type TaobaouniversalbpreportqueryaccountAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpreportqueryaccountAPIResponseModel
}

// TaobaouniversalbpreportqueryaccountAPIResponseModel is 账户报表查询 成功返回结果
type TaobaouniversalbpreportqueryaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpreportqueryaccountTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
