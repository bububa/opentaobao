package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusSeatpriceGetAPIResponse 汽车票余票接口 API返回值
// taobao.bus.seatprice.get
//
// 提供给商家，查询汽车票班次余票
type TaobaoBusSeatpriceGetAPIResponse struct {
	model.CommonResponse
	TaobaoBusSeatpriceGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusSeatpriceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusSeatpriceGetAPIResponseModel).Reset()
}

// TaobaoBusSeatpriceGetAPIResponseModel is 汽车票余票接口 成功返回结果
type TaobaoBusSeatpriceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_seatprice_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoBusSeatpriceGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusSeatpriceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBusSeatpriceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusSeatpriceGetAPIResponse)
	},
}

// GetTaobaoBusSeatpriceGetAPIResponse 从 sync.Pool 获取 TaobaoBusSeatpriceGetAPIResponse
func GetTaobaoBusSeatpriceGetAPIResponse() *TaobaoBusSeatpriceGetAPIResponse {
	return poolTaobaoBusSeatpriceGetAPIResponse.Get().(*TaobaoBusSeatpriceGetAPIResponse)
}

// ReleaseTaobaoBusSeatpriceGetAPIResponse 将 TaobaoBusSeatpriceGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusSeatpriceGetAPIResponse(v *TaobaoBusSeatpriceGetAPIResponse) {
	v.Reset()
	poolTaobaoBusSeatpriceGetAPIResponse.Put(v)
}
