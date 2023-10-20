package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimRelationsGetAPIResponse 获取openim账号的聊天关系 API返回值
// taobao.openim.relations.get
//
// 获取openim账号的聊天关系
type TaobaoOpenimRelationsGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimRelationsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimRelationsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimRelationsGetAPIResponseModel).Reset()
}

// TaobaoOpenimRelationsGetAPIResponseModel is 获取openim账号的聊天关系 成功返回结果
type TaobaoOpenimRelationsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_relations_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户信息列表
	Users []OpenImUser `json:"users,omitempty" xml:"users>open_im_user,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimRelationsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Users = m.Users[:0]
}

var poolTaobaoOpenimRelationsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimRelationsGetAPIResponse)
	},
}

// GetTaobaoOpenimRelationsGetAPIResponse 从 sync.Pool 获取 TaobaoOpenimRelationsGetAPIResponse
func GetTaobaoOpenimRelationsGetAPIResponse() *TaobaoOpenimRelationsGetAPIResponse {
	return poolTaobaoOpenimRelationsGetAPIResponse.Get().(*TaobaoOpenimRelationsGetAPIResponse)
}

// ReleaseTaobaoOpenimRelationsGetAPIResponse 将 TaobaoOpenimRelationsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimRelationsGetAPIResponse(v *TaobaoOpenimRelationsGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenimRelationsGetAPIResponse.Put(v)
}
