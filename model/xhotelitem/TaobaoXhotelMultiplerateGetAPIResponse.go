package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMultiplerateGetAPIResponse 复杂房价查询接口 API返回值
// taobao.xhotel.multiplerate.get
//
// 查询复杂房价，支持通过入住人数，连住天数，商品信息，房价信息查询
type TaobaoXhotelMultiplerateGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelMultiplerateGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelMultiplerateGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelMultiplerateGetAPIResponseModel).Reset()
}

// TaobaoXhotelMultiplerateGetAPIResponseModel is 复杂房价查询接口 成功返回结果
type TaobaoXhotelMultiplerateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_multiplerate_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 复杂价格返回结果类
	Rates []MultipleRate `json:"rates,omitempty" xml:"rates>multiple_rate,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelMultiplerateGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Rates = m.Rates[:0]
}

var poolTaobaoXhotelMultiplerateGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelMultiplerateGetAPIResponse)
	},
}

// GetTaobaoXhotelMultiplerateGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelMultiplerateGetAPIResponse
func GetTaobaoXhotelMultiplerateGetAPIResponse() *TaobaoXhotelMultiplerateGetAPIResponse {
	return poolTaobaoXhotelMultiplerateGetAPIResponse.Get().(*TaobaoXhotelMultiplerateGetAPIResponse)
}

// ReleaseTaobaoXhotelMultiplerateGetAPIResponse 将 TaobaoXhotelMultiplerateGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelMultiplerateGetAPIResponse(v *TaobaoXhotelMultiplerateGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelMultiplerateGetAPIResponse.Put(v)
}
