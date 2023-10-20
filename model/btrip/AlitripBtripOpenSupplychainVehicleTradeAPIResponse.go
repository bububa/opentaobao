package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopensupplychainvehicletradeAPIResponse 商旅用车交易流水接口 API返回值
// alitrip.btrip.open.supplychain.vehicle.trade
//
// 商旅用车交易流水接口
type AlitripbtripopensupplychainvehicletradeAPIResponse struct {
	model.CommonResponse
	AlitripbtripopensupplychainvehicletradeAPIResponseModel
}

// AlitripbtripopensupplychainvehicletradeAPIResponseModel is 商旅用车交易流水接口 成功返回结果
type AlitripbtripopensupplychainvehicletradeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_open_supplychain_vehicle_trade_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
