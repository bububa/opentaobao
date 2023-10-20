package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse 下载飞猪国际城市结果 API返回值
// taobao.xhotel.city.coordinates.batch.download
//
// 给国际酒店供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程
type TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelCityCoordinatesBatchDownloadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelCityCoordinatesBatchDownloadAPIResponseModel).Reset()
}

// TaobaoXhotelCityCoordinatesBatchDownloadAPIResponseModel is 下载飞猪国际城市结果 成功返回结果
type TaobaoXhotelCityCoordinatesBatchDownloadAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_city_coordinates_batch_download_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 经纬度计算结果列表
	CoordinateList []Coordinate `json:"coordinate_list,omitempty" xml:"coordinate_list>coordinate,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelCityCoordinatesBatchDownloadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CoordinateList = m.CoordinateList[:0]
}

var poolTaobaoXhotelCityCoordinatesBatchDownloadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse)
	},
}

// GetTaobaoXhotelCityCoordinatesBatchDownloadAPIResponse 从 sync.Pool 获取 TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse
func GetTaobaoXhotelCityCoordinatesBatchDownloadAPIResponse() *TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse {
	return poolTaobaoXhotelCityCoordinatesBatchDownloadAPIResponse.Get().(*TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse)
}

// ReleaseTaobaoXhotelCityCoordinatesBatchDownloadAPIResponse 将 TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelCityCoordinatesBatchDownloadAPIResponse(v *TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse) {
	v.Reset()
	poolTaobaoXhotelCityCoordinatesBatchDownloadAPIResponse.Put(v)
}
