package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryReportAPIResponse 库存盘点通知接口 API返回值
// taobao.qimen.inventory.report
//
// WMS调用奇门的接口,将库存盘点情况回传ERP
type TaobaoQimenInventoryReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenInventoryReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenInventoryReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenInventoryReportAPIResponseModel).Reset()
}

// TaobaoQimenInventoryReportAPIResponseModel is 库存盘点通知接口 成功返回结果
type TaobaoQimenInventoryReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventory_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenInventoryReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenInventoryReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenInventoryReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventoryReportAPIResponse)
	},
}

// GetTaobaoQimenInventoryReportAPIResponse 从 sync.Pool 获取 TaobaoQimenInventoryReportAPIResponse
func GetTaobaoQimenInventoryReportAPIResponse() *TaobaoQimenInventoryReportAPIResponse {
	return poolTaobaoQimenInventoryReportAPIResponse.Get().(*TaobaoQimenInventoryReportAPIResponse)
}

// ReleaseTaobaoQimenInventoryReportAPIResponse 将 TaobaoQimenInventoryReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenInventoryReportAPIResponse(v *TaobaoQimenInventoryReportAPIResponse) {
	v.Reset()
	poolTaobaoQimenInventoryReportAPIResponse.Put(v)
}
