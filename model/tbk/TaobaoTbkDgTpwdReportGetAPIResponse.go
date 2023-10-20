package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgTpwdReportGetAPIResponse 淘宝客-推广者-淘口令回流数据查询 API返回值
// taobao.tbk.dg.tpwd.report.get
//
// 淘宝客获取单个淘口令的回流PV、UV数据。
type TaobaoTbkDgTpwdReportGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgTpwdReportGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkDgTpwdReportGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkDgTpwdReportGetAPIResponseModel).Reset()
}

// TaobaoTbkDgTpwdReportGetAPIResponseModel is 淘宝客-推广者-淘口令回流数据查询 成功返回结果
type TaobaoTbkDgTpwdReportGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_tpwd_report_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *TaobaoTbkDgTpwdReportGetMapData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkDgTpwdReportGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkDgTpwdReportGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgTpwdReportGetAPIResponse)
	},
}

// GetTaobaoTbkDgTpwdReportGetAPIResponse 从 sync.Pool 获取 TaobaoTbkDgTpwdReportGetAPIResponse
func GetTaobaoTbkDgTpwdReportGetAPIResponse() *TaobaoTbkDgTpwdReportGetAPIResponse {
	return poolTaobaoTbkDgTpwdReportGetAPIResponse.Get().(*TaobaoTbkDgTpwdReportGetAPIResponse)
}

// ReleaseTaobaoTbkDgTpwdReportGetAPIResponse 将 TaobaoTbkDgTpwdReportGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkDgTpwdReportGetAPIResponse(v *TaobaoTbkDgTpwdReportGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkDgTpwdReportGetAPIResponse.Put(v)
}
