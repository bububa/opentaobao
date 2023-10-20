package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusBusnumberGetAPIResponse 汽车票车次查询 API返回值
// taobao.bus.busnumber.get
//
// 提供汽车票车次查询服务
type TaobaoBusBusnumberGetAPIResponse struct {
	model.CommonResponse
	TaobaoBusBusnumberGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusBusnumberGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusBusnumberGetAPIResponseModel).Reset()
}

// TaobaoBusBusnumberGetAPIResponseModel is 汽车票车次查询 成功返回结果
type TaobaoBusBusnumberGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_busnumber_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoBusBusnumberGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusBusnumberGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBusBusnumberGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusBusnumberGetAPIResponse)
	},
}

// GetTaobaoBusBusnumberGetAPIResponse 从 sync.Pool 获取 TaobaoBusBusnumberGetAPIResponse
func GetTaobaoBusBusnumberGetAPIResponse() *TaobaoBusBusnumberGetAPIResponse {
	return poolTaobaoBusBusnumberGetAPIResponse.Get().(*TaobaoBusBusnumberGetAPIResponse)
}

// ReleaseTaobaoBusBusnumberGetAPIResponse 将 TaobaoBusBusnumberGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusBusnumberGetAPIResponse(v *TaobaoBusBusnumberGetAPIResponse) {
	v.Reset()
	poolTaobaoBusBusnumberGetAPIResponse.Put(v)
}
