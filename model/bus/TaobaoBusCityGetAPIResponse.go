package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusCityGetAPIResponse 城市接口 API返回值
// taobao.bus.city.get
//
// 汽车票出发城市获取接口，获取所有出发城市
type TaobaoBusCityGetAPIResponse struct {
	model.CommonResponse
	TaobaoBusCityGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusCityGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusCityGetAPIResponseModel).Reset()
}

// TaobaoBusCityGetAPIResponseModel is 城市接口 成功返回结果
type TaobaoBusCityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_city_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 城市返回结果
	Result *CitySearchRp `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusCityGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBusCityGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusCityGetAPIResponse)
	},
}

// GetTaobaoBusCityGetAPIResponse 从 sync.Pool 获取 TaobaoBusCityGetAPIResponse
func GetTaobaoBusCityGetAPIResponse() *TaobaoBusCityGetAPIResponse {
	return poolTaobaoBusCityGetAPIResponse.Get().(*TaobaoBusCityGetAPIResponse)
}

// ReleaseTaobaoBusCityGetAPIResponse 将 TaobaoBusCityGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusCityGetAPIResponse(v *TaobaoBusCityGetAPIResponse) {
	v.Reset()
	poolTaobaoBusCityGetAPIResponse.Put(v)
}
