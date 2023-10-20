package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventorySynchronizeReportAPIResponse 库存状态同步确认接口 API返回值
// taobao.qimen.inventory.synchronize.report
//
// 库存状态同步确认接口
type TaobaoQimenInventorySynchronizeReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenInventorySynchronizeReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenInventorySynchronizeReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenInventorySynchronizeReportAPIResponseModel).Reset()
}

// TaobaoQimenInventorySynchronizeReportAPIResponseModel is 库存状态同步确认接口 成功返回结果
type TaobaoQimenInventorySynchronizeReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventory_synchronize_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenInventorySynchronizeReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenInventorySynchronizeReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenInventorySynchronizeReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventorySynchronizeReportAPIResponse)
	},
}

// GetTaobaoQimenInventorySynchronizeReportAPIResponse 从 sync.Pool 获取 TaobaoQimenInventorySynchronizeReportAPIResponse
func GetTaobaoQimenInventorySynchronizeReportAPIResponse() *TaobaoQimenInventorySynchronizeReportAPIResponse {
	return poolTaobaoQimenInventorySynchronizeReportAPIResponse.Get().(*TaobaoQimenInventorySynchronizeReportAPIResponse)
}

// ReleaseTaobaoQimenInventorySynchronizeReportAPIResponse 将 TaobaoQimenInventorySynchronizeReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenInventorySynchronizeReportAPIResponse(v *TaobaoQimenInventorySynchronizeReportAPIResponse) {
	v.Reset()
	poolTaobaoQimenInventorySynchronizeReportAPIResponse.Put(v)
}
