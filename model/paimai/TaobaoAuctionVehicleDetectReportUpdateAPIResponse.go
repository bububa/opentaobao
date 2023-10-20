package paimai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionVehicleDetectReportUpdateAPIResponse 检测服务-服务单报告信息更新 API返回值
// taobao.auction.vehicle.detect.report.update
//
// 检测服务-服务单报告信息更新
type TaobaoAuctionVehicleDetectReportUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAuctionVehicleDetectReportUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAuctionVehicleDetectReportUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAuctionVehicleDetectReportUpdateAPIResponseModel).Reset()
}

// TaobaoAuctionVehicleDetectReportUpdateAPIResponseModel is 检测服务-服务单报告信息更新 成功返回结果
type TaobaoAuctionVehicleDetectReportUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_vehicle_detect_report_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务返回结果
	Result *Result4Top `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAuctionVehicleDetectReportUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAuctionVehicleDetectReportUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAuctionVehicleDetectReportUpdateAPIResponse)
	},
}

// GetTaobaoAuctionVehicleDetectReportUpdateAPIResponse 从 sync.Pool 获取 TaobaoAuctionVehicleDetectReportUpdateAPIResponse
func GetTaobaoAuctionVehicleDetectReportUpdateAPIResponse() *TaobaoAuctionVehicleDetectReportUpdateAPIResponse {
	return poolTaobaoAuctionVehicleDetectReportUpdateAPIResponse.Get().(*TaobaoAuctionVehicleDetectReportUpdateAPIResponse)
}

// ReleaseTaobaoAuctionVehicleDetectReportUpdateAPIResponse 将 TaobaoAuctionVehicleDetectReportUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAuctionVehicleDetectReportUpdateAPIResponse(v *TaobaoAuctionVehicleDetectReportUpdateAPIResponse) {
	v.Reset()
	poolTaobaoAuctionVehicleDetectReportUpdateAPIResponse.Put(v)
}
