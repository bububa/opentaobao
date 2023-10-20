package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeJoinAPIResponse OPENIM群主动加入 API返回值
// taobao.openim.tribe.join
//
// OPENIM群主动加入
type TaobaoOpenimTribeJoinAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeJoinAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeJoinAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribeJoinAPIResponseModel).Reset()
}

// TaobaoOpenimTribeJoinAPIResponseModel is OPENIM群主动加入 成功返回结果
type TaobaoOpenimTribeJoinAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_join_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群服务code
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeJoinAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TribeCode = 0
}

var poolTaobaoOpenimTribeJoinAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribeJoinAPIResponse)
	},
}

// GetTaobaoOpenimTribeJoinAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribeJoinAPIResponse
func GetTaobaoOpenimTribeJoinAPIResponse() *TaobaoOpenimTribeJoinAPIResponse {
	return poolTaobaoOpenimTribeJoinAPIResponse.Get().(*TaobaoOpenimTribeJoinAPIResponse)
}

// ReleaseTaobaoOpenimTribeJoinAPIResponse 将 TaobaoOpenimTribeJoinAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribeJoinAPIResponse(v *TaobaoOpenimTribeJoinAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribeJoinAPIResponse.Put(v)
}
