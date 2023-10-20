package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeExpelAPIResponse OPENIM群踢出成员 API返回值
// taobao.openim.tribe.expel
//
// OPENIM群踢出成员
type TaobaoOpenimTribeExpelAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeExpelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeExpelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribeExpelAPIResponseModel).Reset()
}

// TaobaoOpenimTribeExpelAPIResponseModel is OPENIM群踢出成员 成功返回结果
type TaobaoOpenimTribeExpelAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_expel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群服务code
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeExpelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TribeCode = 0
}

var poolTaobaoOpenimTribeExpelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribeExpelAPIResponse)
	},
}

// GetTaobaoOpenimTribeExpelAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribeExpelAPIResponse
func GetTaobaoOpenimTribeExpelAPIResponse() *TaobaoOpenimTribeExpelAPIResponse {
	return poolTaobaoOpenimTribeExpelAPIResponse.Get().(*TaobaoOpenimTribeExpelAPIResponse)
}

// ReleaseTaobaoOpenimTribeExpelAPIResponse 将 TaobaoOpenimTribeExpelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribeExpelAPIResponse(v *TaobaoOpenimTribeExpelAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribeExpelAPIResponse.Put(v)
}
