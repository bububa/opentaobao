package rail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailIrCarrierGetAPIResponse 国际火车票铁路承运公司查询 API返回值
// alitrip.rail.ir.carrier.get
//
// 国际火车票提供给代理商用于查询标准铁路承运公司carrier信息，用于代理商自己的carrier与飞猪平台的carrier做映射
type AlitripRailIrCarrierGetAPIResponse struct {
	model.CommonResponse
	AlitripRailIrCarrierGetAPIResponseModel
}

// AlitripRailIrCarrierGetAPIResponseModel is 国际火车票铁路承运公司查询 成功返回结果
type AlitripRailIrCarrierGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_rail_ir_carrier_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *RailCarrierRs `json:"result,omitempty" xml:"result,omitempty"`
}
