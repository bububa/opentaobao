package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresscapacitytmsasyncAPIResponse 上门取退产能信息同步/更新 API返回值
// taobao.logistics.express.capacity.tms.async
//
// 上门取退产能信息同步/更新
type TaobaologisticsexpresscapacitytmsasyncAPIResponse struct {
	model.CommonResponse
	TaobaologisticsexpresscapacitytmsasyncAPIResponseModel
}

// TaobaologisticsexpresscapacitytmsasyncAPIResponseModel is 上门取退产能信息同步/更新 成功返回结果
type TaobaologisticsexpresscapacitytmsasyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_capacity_tms_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	CapacityResponse *CapacityResponse `json:"capacity_response,omitempty" xml:"capacity_response,omitempty"`
}
