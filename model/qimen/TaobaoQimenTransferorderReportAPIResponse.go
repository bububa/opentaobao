package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTransferorderReportAPIResponse 调拨单通知 API返回值
// taobao.qimen.transferorder.report
//
// 调拨单通知
type TaobaoQimenTransferorderReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenTransferorderReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenTransferorderReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenTransferorderReportAPIResponseModel).Reset()
}

// TaobaoQimenTransferorderReportAPIResponseModel is 调拨单通知 成功返回结果
type TaobaoQimenTransferorderReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_transferorder_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenTransferorderReportStruct `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenTransferorderReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenTransferorderReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenTransferorderReportAPIResponse)
	},
}

// GetTaobaoQimenTransferorderReportAPIResponse 从 sync.Pool 获取 TaobaoQimenTransferorderReportAPIResponse
func GetTaobaoQimenTransferorderReportAPIResponse() *TaobaoQimenTransferorderReportAPIResponse {
	return poolTaobaoQimenTransferorderReportAPIResponse.Get().(*TaobaoQimenTransferorderReportAPIResponse)
}

// ReleaseTaobaoQimenTransferorderReportAPIResponse 将 TaobaoQimenTransferorderReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenTransferorderReportAPIResponse(v *TaobaoQimenTransferorderReportAPIResponse) {
	v.Reset()
	poolTaobaoQimenTransferorderReportAPIResponse.Put(v)
}
