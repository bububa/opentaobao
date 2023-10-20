package cainiaolocker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEndpointLockerTopStationAddorupdateAPIResponse 增加更新代收点 API返回值
// cainiao.endpoint.locker.top.station.addorupdate
//
// 新增或者修改代收点相关信息
type CainiaoEndpointLockerTopStationAddorupdateAPIResponse struct {
	model.CommonResponse
	CainiaoEndpointLockerTopStationAddorupdateAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoEndpointLockerTopStationAddorupdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoEndpointLockerTopStationAddorupdateAPIResponseModel).Reset()
}

// CainiaoEndpointLockerTopStationAddorupdateAPIResponseModel is 增加更新代收点 成功返回结果
type CainiaoEndpointLockerTopStationAddorupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_endpoint_locker_top_station_addorupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoEndpointLockerTopStationAddorupdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoEndpointLockerTopStationAddorupdateAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoEndpointLockerTopStationAddorupdateAPIResponse)
	},
}

// GetCainiaoEndpointLockerTopStationAddorupdateAPIResponse 从 sync.Pool 获取 CainiaoEndpointLockerTopStationAddorupdateAPIResponse
func GetCainiaoEndpointLockerTopStationAddorupdateAPIResponse() *CainiaoEndpointLockerTopStationAddorupdateAPIResponse {
	return poolCainiaoEndpointLockerTopStationAddorupdateAPIResponse.Get().(*CainiaoEndpointLockerTopStationAddorupdateAPIResponse)
}

// ReleaseCainiaoEndpointLockerTopStationAddorupdateAPIResponse 将 CainiaoEndpointLockerTopStationAddorupdateAPIResponse 保存到 sync.Pool
func ReleaseCainiaoEndpointLockerTopStationAddorupdateAPIResponse(v *CainiaoEndpointLockerTopStationAddorupdateAPIResponse) {
	v.Reset()
	poolCainiaoEndpointLockerTopStationAddorupdateAPIResponse.Put(v)
}
