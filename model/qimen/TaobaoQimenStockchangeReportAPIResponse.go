package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStockchangeReportAPIResponse 库存异动通知接口 API返回值
// taobao.qimen.stockchange.report
//
// taobao.qimen.stockchange.report
type TaobaoQimenStockchangeReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStockchangeReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenStockchangeReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenStockchangeReportAPIResponseModel).Reset()
}

// TaobaoQimenStockchangeReportAPIResponseModel is 库存异动通知接口 成功返回结果
type TaobaoQimenStockchangeReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_stockchange_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenStockchangeReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenStockchangeReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenStockchangeReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockchangeReportAPIResponse)
	},
}

// GetTaobaoQimenStockchangeReportAPIResponse 从 sync.Pool 获取 TaobaoQimenStockchangeReportAPIResponse
func GetTaobaoQimenStockchangeReportAPIResponse() *TaobaoQimenStockchangeReportAPIResponse {
	return poolTaobaoQimenStockchangeReportAPIResponse.Get().(*TaobaoQimenStockchangeReportAPIResponse)
}

// ReleaseTaobaoQimenStockchangeReportAPIResponse 将 TaobaoQimenStockchangeReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenStockchangeReportAPIResponse(v *TaobaoQimenStockchangeReportAPIResponse) {
	v.Reset()
	poolTaobaoQimenStockchangeReportAPIResponse.Put(v)
}
