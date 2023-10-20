package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScVegasSendReportAPIResponse 淘宝客-服务商-查询红包发放个数 API返回值
// taobao.tbk.sc.vegas.send.report
//
// 服务商使用。可通过此接口查询对应推广者账号下的红包发放个数。
type TaobaoTbkScVegasSendReportAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScVegasSendReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScVegasSendReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScVegasSendReportAPIResponseModel).Reset()
}

// TaobaoTbkScVegasSendReportAPIResponseModel is 淘宝客-服务商-查询红包发放个数 成功返回结果
type TaobaoTbkScVegasSendReportAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_vegas_send_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoTbkScVegasSendReportResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScVegasSendReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTbkScVegasSendReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScVegasSendReportAPIResponse)
	},
}

// GetTaobaoTbkScVegasSendReportAPIResponse 从 sync.Pool 获取 TaobaoTbkScVegasSendReportAPIResponse
func GetTaobaoTbkScVegasSendReportAPIResponse() *TaobaoTbkScVegasSendReportAPIResponse {
	return poolTaobaoTbkScVegasSendReportAPIResponse.Get().(*TaobaoTbkScVegasSendReportAPIResponse)
}

// ReleaseTaobaoTbkScVegasSendReportAPIResponse 将 TaobaoTbkScVegasSendReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScVegasSendReportAPIResponse(v *TaobaoTbkScVegasSendReportAPIResponse) {
	v.Reset()
	poolTaobaoTbkScVegasSendReportAPIResponse.Put(v)
}
