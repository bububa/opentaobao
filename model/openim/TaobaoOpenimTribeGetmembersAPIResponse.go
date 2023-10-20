package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeGetmembersAPIResponse OPENIM群成员获取 API返回值
// taobao.openim.tribe.getmembers
//
// OPENIM群成员获取
type TaobaoOpenimTribeGetmembersAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeGetmembersAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeGetmembersAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribeGetmembersAPIResponseModel).Reset()
}

// TaobaoOpenimTribeGetmembersAPIResponseModel is OPENIM群成员获取 成功返回结果
type TaobaoOpenimTribeGetmembersAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_getmembers_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// OPENIM群成员列表
	TribeUserList []TribeUser `json:"tribe_user_list,omitempty" xml:"tribe_user_list>tribe_user,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeGetmembersAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TribeUserList = m.TribeUserList[:0]
}

var poolTaobaoOpenimTribeGetmembersAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribeGetmembersAPIResponse)
	},
}

// GetTaobaoOpenimTribeGetmembersAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribeGetmembersAPIResponse
func GetTaobaoOpenimTribeGetmembersAPIResponse() *TaobaoOpenimTribeGetmembersAPIResponse {
	return poolTaobaoOpenimTribeGetmembersAPIResponse.Get().(*TaobaoOpenimTribeGetmembersAPIResponse)
}

// ReleaseTaobaoOpenimTribeGetmembersAPIResponse 将 TaobaoOpenimTribeGetmembersAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribeGetmembersAPIResponse(v *TaobaoOpenimTribeGetmembersAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribeGetmembersAPIResponse.Put(v)
}
