package paimai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionZcVehicleDetectStatusProcessAPIResponse 检测服务-服务单状态流转 API返回值
// taobao.auction.zc.vehicle.detect.status.process
//
// 检测服务-服务单状态流转
type TaobaoAuctionZcVehicleDetectStatusProcessAPIResponse struct {
	model.CommonResponse
	TaobaoAuctionZcVehicleDetectStatusProcessAPIResponseModel
}

// TaobaoAuctionZcVehicleDetectStatusProcessAPIResponseModel is 检测服务-服务单状态流转 成功返回结果
type TaobaoAuctionZcVehicleDetectStatusProcessAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_zc_vehicle_detect_status_process_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务返回结果
	Result *Result4top `json:"result,omitempty" xml:"result,omitempty"`
}
