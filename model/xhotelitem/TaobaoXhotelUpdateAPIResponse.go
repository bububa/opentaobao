package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelUpdateAPIResponse 酒店更新接口（ID不存在自动新增） API返回值
// taobao.xhotel.update
//
// 酒店更新接口
type TaobaoXhotelUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelUpdateAPIResponseModel is 酒店更新接口（ID不存在自动新增） 成功返回结果
type TaobaoXhotelUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店信息
	Xhotel *XHotel `json:"xhotel,omitempty" xml:"xhotel,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Xhotel = nil
}

var poolTaobaoXhotelUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelUpdateAPIResponse)
	},
}

// GetTaobaoXhotelUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelUpdateAPIResponse
func GetTaobaoXhotelUpdateAPIResponse() *TaobaoXhotelUpdateAPIResponse {
	return poolTaobaoXhotelUpdateAPIResponse.Get().(*TaobaoXhotelUpdateAPIResponse)
}

// ReleaseTaobaoXhotelUpdateAPIResponse 将 TaobaoXhotelUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelUpdateAPIResponse(v *TaobaoXhotelUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelUpdateAPIResponse.Put(v)
}
