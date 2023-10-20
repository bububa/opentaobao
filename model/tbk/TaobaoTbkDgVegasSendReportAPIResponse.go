package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgvegassendreportAPIResponse 淘宝客-推广者-查询红包发放个数 API返回值
// taobao.tbk.dg.vegas.send.report
//
// 查询账号下的红包发放个数。
type TaobaotbkdgvegassendreportAPIResponse struct {
	model.CommonResponse
	TaobaotbkdgvegassendreportAPIResponseModel
}

// TaobaotbkdgvegassendreportAPIResponseModel is 淘宝客-推广者-查询红包发放个数 成功返回结果
type TaobaotbkdgvegassendreportAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_vegas_send_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaotbkdgvegassendreportResult `json:"result,omitempty" xml:"result,omitempty"`
}
