package qt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTsPropertyGetAPIResponse 淘宝服务属性查询 API返回值
// taobao.ts.property.get
//
// 淘宝服务属性查询
type TaobaoTsPropertyGetAPIResponse struct {
	model.CommonResponse
	TaobaoTsPropertyGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTsPropertyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTsPropertyGetAPIResponseModel).Reset()
}

// TaobaoTsPropertyGetAPIResponseModel is 淘宝服务属性查询 成功返回结果
type TaobaoTsPropertyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ts_property_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务收费项相关属性对象
	ServiceItemProperty *ServiceItemProperty `json:"service_item_property,omitempty" xml:"service_item_property,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTsPropertyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceItemProperty = nil
}

var poolTaobaoTsPropertyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTsPropertyGetAPIResponse)
	},
}

// GetTaobaoTsPropertyGetAPIResponse 从 sync.Pool 获取 TaobaoTsPropertyGetAPIResponse
func GetTaobaoTsPropertyGetAPIResponse() *TaobaoTsPropertyGetAPIResponse {
	return poolTaobaoTsPropertyGetAPIResponse.Get().(*TaobaoTsPropertyGetAPIResponse)
}

// ReleaseTaobaoTsPropertyGetAPIResponse 将 TaobaoTsPropertyGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTsPropertyGetAPIResponse(v *TaobaoTsPropertyGetAPIResponse) {
	v.Reset()
	poolTaobaoTsPropertyGetAPIResponse.Put(v)
}
