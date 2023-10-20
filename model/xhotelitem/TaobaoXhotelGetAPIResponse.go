package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelGetAPIResponse 酒店查询接口 API返回值
// taobao.xhotel.get
//
// 酒店查询接口
type TaobaoXhotelGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelGetAPIResponseModel).Reset()
}

// TaobaoXhotelGetAPIResponseModel is 酒店查询接口 成功返回结果
type TaobaoXhotelGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询得到的hotel
	Xhotel *FirstResult `json:"xhotel,omitempty" xml:"xhotel,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Xhotel = nil
}

var poolTaobaoXhotelGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelGetAPIResponse)
	},
}

// GetTaobaoXhotelGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelGetAPIResponse
func GetTaobaoXhotelGetAPIResponse() *TaobaoXhotelGetAPIResponse {
	return poolTaobaoXhotelGetAPIResponse.Get().(*TaobaoXhotelGetAPIResponse)
}

// ReleaseTaobaoXhotelGetAPIResponse 将 TaobaoXhotelGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelGetAPIResponse(v *TaobaoXhotelGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelGetAPIResponse.Put(v)
}
