package rail

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailIrStationGetAPIResponse 国际火车票标准车站查询 API返回值
// alitrip.rail.ir.station.get
//
// 国际火车票提供给代理商用于查询标准车站信息，用于代理商对自己的车站与飞猪平台的车站做映射
type AlitripRailIrStationGetAPIResponse struct {
	model.CommonResponse
	AlitripRailIrStationGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripRailIrStationGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripRailIrStationGetAPIResponseModel).Reset()
}

// AlitripRailIrStationGetAPIResponseModel is 国际火车票标准车站查询 成功返回结果
type AlitripRailIrStationGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_rail_ir_station_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *RailResultList `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripRailIrStationGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripRailIrStationGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripRailIrStationGetAPIResponse)
	},
}

// GetAlitripRailIrStationGetAPIResponse 从 sync.Pool 获取 AlitripRailIrStationGetAPIResponse
func GetAlitripRailIrStationGetAPIResponse() *AlitripRailIrStationGetAPIResponse {
	return poolAlitripRailIrStationGetAPIResponse.Get().(*AlitripRailIrStationGetAPIResponse)
}

// ReleaseAlitripRailIrStationGetAPIResponse 将 AlitripRailIrStationGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripRailIrStationGetAPIResponse(v *AlitripRailIrStationGetAPIResponse) {
	v.Reset()
	poolAlitripRailIrStationGetAPIResponse.Put(v)
}
