package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemlackReportAPIResponse 发货单缺货通知接口 API返回值
// taobao.qimen.itemlack.report
//
// WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
type TaobaoQimenItemlackReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenItemlackReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenItemlackReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenItemlackReportAPIResponseModel).Reset()
}

// TaobaoQimenItemlackReportAPIResponseModel is 发货单缺货通知接口 成功返回结果
type TaobaoQimenItemlackReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_itemlack_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenItemlackReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenItemlackReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenItemlackReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenItemlackReportAPIResponse)
	},
}

// GetTaobaoQimenItemlackReportAPIResponse 从 sync.Pool 获取 TaobaoQimenItemlackReportAPIResponse
func GetTaobaoQimenItemlackReportAPIResponse() *TaobaoQimenItemlackReportAPIResponse {
	return poolTaobaoQimenItemlackReportAPIResponse.Get().(*TaobaoQimenItemlackReportAPIResponse)
}

// ReleaseTaobaoQimenItemlackReportAPIResponse 将 TaobaoQimenItemlackReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenItemlackReportAPIResponse(v *TaobaoQimenItemlackReportAPIResponse) {
	v.Reset()
	poolTaobaoQimenItemlackReportAPIResponse.Put(v)
}
