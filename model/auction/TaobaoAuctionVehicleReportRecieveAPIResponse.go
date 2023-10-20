package auction

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionVehicleReportRecieveAPIResponse 机动车报告回调数据接收 API返回值
// taobao.auction.vehicle.report.recieve
//
// 机动车报告同步接收接口
type TaobaoAuctionVehicleReportRecieveAPIResponse struct {
	model.CommonResponse
	TaobaoAuctionVehicleReportRecieveAPIResponseModel
}

// TaobaoAuctionVehicleReportRecieveAPIResponseModel is 机动车报告回调数据接收 成功返回结果
type TaobaoAuctionVehicleReportRecieveAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_vehicle_report_recieve_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求返回对象
	Result *Result4top `json:"result,omitempty" xml:"result,omitempty"`
}
