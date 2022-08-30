package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasSendReportAPIResponse 淘宝客-推广者-查询红包发放个数 API返回值
// taobao.tbk.dg.vegas.send.report
//
// 查询账号下的红包发放个数。
type TaobaoTbkDgVegasSendReportAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgVegasSendReportAPIResponseModel
}

// TaobaoTbkDgVegasSendReportAPIResponseModel is 淘宝客-推广者-查询红包发放个数 成功返回结果
type TaobaoTbkDgVegasSendReportAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_vegas_send_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoTbkDgVegasSendReportResult `json:"result,omitempty" xml:"result,omitempty"`
}
