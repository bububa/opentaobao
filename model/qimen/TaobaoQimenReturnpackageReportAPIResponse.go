package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReturnpackageReportAPIResponse 退货包裹状态通知接口 API返回值
// taobao.qimen.returnpackage.report
//
// 退货包裹状态通知接口
type TaobaoQimenReturnpackageReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenReturnpackageReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenReturnpackageReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenReturnpackageReportAPIResponseModel).Reset()
}

// TaobaoQimenReturnpackageReportAPIResponseModel is 退货包裹状态通知接口 成功返回结果
type TaobaoQimenReturnpackageReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_returnpackage_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenReturnpackageReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenReturnpackageReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenReturnpackageReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenReturnpackageReportAPIResponse)
	},
}

// GetTaobaoQimenReturnpackageReportAPIResponse 从 sync.Pool 获取 TaobaoQimenReturnpackageReportAPIResponse
func GetTaobaoQimenReturnpackageReportAPIResponse() *TaobaoQimenReturnpackageReportAPIResponse {
	return poolTaobaoQimenReturnpackageReportAPIResponse.Get().(*TaobaoQimenReturnpackageReportAPIResponse)
}

// ReleaseTaobaoQimenReturnpackageReportAPIResponse 将 TaobaoQimenReturnpackageReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenReturnpackageReportAPIResponse(v *TaobaoQimenReturnpackageReportAPIResponse) {
	v.Reset()
	poolTaobaoQimenReturnpackageReportAPIResponse.Put(v)
}
