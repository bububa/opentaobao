package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderprocessReportAPIResponse 订单流水通知接口 API返回值
// taobao.qimen.orderprocess.report
//
// taobao.qimen.orderprocess.report
type TaobaoQimenOrderprocessReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenOrderprocessReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenOrderprocessReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenOrderprocessReportAPIResponseModel).Reset()
}

// TaobaoQimenOrderprocessReportAPIResponseModel is 订单流水通知接口 成功返回结果
type TaobaoQimenOrderprocessReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_orderprocess_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenOrderprocessReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenOrderprocessReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenOrderprocessReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderprocessReportAPIResponse)
	},
}

// GetTaobaoQimenOrderprocessReportAPIResponse 从 sync.Pool 获取 TaobaoQimenOrderprocessReportAPIResponse
func GetTaobaoQimenOrderprocessReportAPIResponse() *TaobaoQimenOrderprocessReportAPIResponse {
	return poolTaobaoQimenOrderprocessReportAPIResponse.Get().(*TaobaoQimenOrderprocessReportAPIResponse)
}

// ReleaseTaobaoQimenOrderprocessReportAPIResponse 将 TaobaoQimenOrderprocessReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenOrderprocessReportAPIResponse(v *TaobaoQimenOrderprocessReportAPIResponse) {
	v.Reset()
	poolTaobaoQimenOrderprocessReportAPIResponse.Put(v)
}
