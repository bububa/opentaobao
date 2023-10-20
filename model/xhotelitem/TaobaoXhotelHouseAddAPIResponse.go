package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelHouseAddAPIResponse 非标准民宿房源添加 API返回值
// taobao.xhotel.house.add
//
// 添加酒店或更新酒店
type TaobaoXhotelHouseAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelHouseAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelHouseAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelHouseAddAPIResponseModel).Reset()
}

// TaobaoXhotelHouseAddAPIResponseModel is 非标准民宿房源添加 成功返回结果
type TaobaoXhotelHouseAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_house_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店信息
	Xhotel *XHotel `json:"xhotel,omitempty" xml:"xhotel,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelHouseAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Xhotel = nil
}

var poolTaobaoXhotelHouseAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelHouseAddAPIResponse)
	},
}

// GetTaobaoXhotelHouseAddAPIResponse 从 sync.Pool 获取 TaobaoXhotelHouseAddAPIResponse
func GetTaobaoXhotelHouseAddAPIResponse() *TaobaoXhotelHouseAddAPIResponse {
	return poolTaobaoXhotelHouseAddAPIResponse.Get().(*TaobaoXhotelHouseAddAPIResponse)
}

// ReleaseTaobaoXhotelHouseAddAPIResponse 将 TaobaoXhotelHouseAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelHouseAddAPIResponse(v *TaobaoXhotelHouseAddAPIResponse) {
	v.Reset()
	poolTaobaoXhotelHouseAddAPIResponse.Put(v)
}
