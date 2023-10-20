package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeCreateAPIResponse 创建群 API返回值
// taobao.openim.tribe.create
//
// 创建一个openim的群
type TaobaoOpenimTribeCreateAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribeCreateAPIResponseModel).Reset()
}

// TaobaoOpenimTribeCreateAPIResponseModel is 创建群 成功返回结果
type TaobaoOpenimTribeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建群的信息
	TribeInfo *TribeInfo `json:"tribe_info,omitempty" xml:"tribe_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TribeInfo = nil
}

var poolTaobaoOpenimTribeCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribeCreateAPIResponse)
	},
}

// GetTaobaoOpenimTribeCreateAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribeCreateAPIResponse
func GetTaobaoOpenimTribeCreateAPIResponse() *TaobaoOpenimTribeCreateAPIResponse {
	return poolTaobaoOpenimTribeCreateAPIResponse.Get().(*TaobaoOpenimTribeCreateAPIResponse)
}

// ReleaseTaobaoOpenimTribeCreateAPIResponse 将 TaobaoOpenimTribeCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribeCreateAPIResponse(v *TaobaoOpenimTribeCreateAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribeCreateAPIResponse.Put(v)
}
