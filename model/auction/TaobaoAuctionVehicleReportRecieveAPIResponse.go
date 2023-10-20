package auction

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoAuctionVehicleReportRecieveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAuctionVehicleReportRecieveAPIResponseModel).Reset()
}

// TaobaoAuctionVehicleReportRecieveAPIResponseModel is 机动车报告回调数据接收 成功返回结果
type TaobaoAuctionVehicleReportRecieveAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_vehicle_report_recieve_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求返回对象
	Result *Result4Top `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAuctionVehicleReportRecieveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAuctionVehicleReportRecieveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAuctionVehicleReportRecieveAPIResponse)
	},
}

// GetTaobaoAuctionVehicleReportRecieveAPIResponse 从 sync.Pool 获取 TaobaoAuctionVehicleReportRecieveAPIResponse
func GetTaobaoAuctionVehicleReportRecieveAPIResponse() *TaobaoAuctionVehicleReportRecieveAPIResponse {
	return poolTaobaoAuctionVehicleReportRecieveAPIResponse.Get().(*TaobaoAuctionVehicleReportRecieveAPIResponse)
}

// ReleaseTaobaoAuctionVehicleReportRecieveAPIResponse 将 TaobaoAuctionVehicleReportRecieveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAuctionVehicleReportRecieveAPIResponse(v *TaobaoAuctionVehicleReportRecieveAPIResponse) {
	v.Reset()
	poolTaobaoAuctionVehicleReportRecieveAPIResponse.Put(v)
}
