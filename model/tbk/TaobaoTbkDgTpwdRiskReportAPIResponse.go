package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgTpwdRiskReportAPIResponse 淘宝客-推广者-淘口令预警及拦截查询 API返回值
// taobao.tbk.dg.tpwd.risk.report
//
// 淘宝客-推广者-淘口令预警及拦截查询
type TaobaoTbkDgTpwdRiskReportAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgTpwdRiskReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkDgTpwdRiskReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkDgTpwdRiskReportAPIResponseModel).Reset()
}

// TaobaoTbkDgTpwdRiskReportAPIResponseModel is 淘宝客-推广者-淘口令预警及拦截查询 成功返回结果
type TaobaoTbkDgTpwdRiskReportAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_tpwd_risk_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	Result *TaobaoTbkDgTpwdRiskReportResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkDgTpwdRiskReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTbkDgTpwdRiskReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgTpwdRiskReportAPIResponse)
	},
}

// GetTaobaoTbkDgTpwdRiskReportAPIResponse 从 sync.Pool 获取 TaobaoTbkDgTpwdRiskReportAPIResponse
func GetTaobaoTbkDgTpwdRiskReportAPIResponse() *TaobaoTbkDgTpwdRiskReportAPIResponse {
	return poolTaobaoTbkDgTpwdRiskReportAPIResponse.Get().(*TaobaoTbkDgTpwdRiskReportAPIResponse)
}

// ReleaseTaobaoTbkDgTpwdRiskReportAPIResponse 将 TaobaoTbkDgTpwdRiskReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkDgTpwdRiskReportAPIResponse(v *TaobaoTbkDgTpwdRiskReportAPIResponse) {
	v.Reset()
	poolTaobaoTbkDgTpwdRiskReportAPIResponse.Put(v)
}
