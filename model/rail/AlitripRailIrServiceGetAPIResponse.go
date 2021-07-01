package rail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripRailIrServiceGetAPIResponse
国际火车票仓位坐席查询 API返回值
alitrip.rail.ir.service.get

国际火车票标准仓位坐席查询 */
type AlitripRailIrServiceGetAPIResponse struct {
	model.CommonResponse
	AlitripRailIrServiceGetAPIResponseModel
}

// AlitripRailIrServiceGetAPIResponseModel is 国际火车票仓位坐席查询 成功返回结果
type AlitripRailIrServiceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_rail_ir_service_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripRailIrServiceGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
