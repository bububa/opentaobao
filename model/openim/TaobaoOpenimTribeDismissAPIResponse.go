package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeDismissAPIResponse OPENIM群解散 API返回值
// taobao.openim.tribe.dismiss
//
// OPENIM群解散
type TaobaoOpenimTribeDismissAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeDismissAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeDismissAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribeDismissAPIResponseModel).Reset()
}

// TaobaoOpenimTribeDismissAPIResponseModel is OPENIM群解散 成功返回结果
type TaobaoOpenimTribeDismissAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_dismiss_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群服务code
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeDismissAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TribeCode = 0
}

var poolTaobaoOpenimTribeDismissAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribeDismissAPIResponse)
	},
}

// GetTaobaoOpenimTribeDismissAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribeDismissAPIResponse
func GetTaobaoOpenimTribeDismissAPIResponse() *TaobaoOpenimTribeDismissAPIResponse {
	return poolTaobaoOpenimTribeDismissAPIResponse.Get().(*TaobaoOpenimTribeDismissAPIResponse)
}

// ReleaseTaobaoOpenimTribeDismissAPIResponse 将 TaobaoOpenimTribeDismissAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribeDismissAPIResponse(v *TaobaoOpenimTribeDismissAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribeDismissAPIResponse.Put(v)
}
