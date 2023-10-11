package paimai

import (
	"encoding/xml"

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

// TaobaoAuctionVehicleDetectReportUpdateAPIResponseModel is 检测服务-服务单报告信息更新 成功返回结果
type TaobaoAuctionVehicleDetectReportUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_vehicle_detect_report_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务返回结果
	Result *Result4Top `json:"result,omitempty" xml:"result,omitempty"`
}
