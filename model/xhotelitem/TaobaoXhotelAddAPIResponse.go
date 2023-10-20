package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelAddAPIResponse 酒店新增接口（ID重复自动更新） API返回值
// taobao.xhotel.add
//
// 添加酒店或更新酒店
type TaobaoXhotelAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelAddAPIResponseModel).Reset()
}

// TaobaoXhotelAddAPIResponseModel is 酒店新增接口（ID重复自动更新） 成功返回结果
type TaobaoXhotelAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店信息
	Xhotel *XHotel `json:"xhotel,omitempty" xml:"xhotel,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Xhotel = nil
}

var poolTaobaoXhotelAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelAddAPIResponse)
	},
}

// GetTaobaoXhotelAddAPIResponse 从 sync.Pool 获取 TaobaoXhotelAddAPIResponse
func GetTaobaoXhotelAddAPIResponse() *TaobaoXhotelAddAPIResponse {
	return poolTaobaoXhotelAddAPIResponse.Get().(*TaobaoXhotelAddAPIResponse)
}

// ReleaseTaobaoXhotelAddAPIResponse 将 TaobaoXhotelAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelAddAPIResponse(v *TaobaoXhotelAddAPIResponse) {
	v.Reset()
	poolTaobaoXhotelAddAPIResponse.Put(v)
}
