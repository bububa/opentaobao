package jms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsGroupGetAPIResponse 查询ONS分组 API返回值
// taobao.jushita.jms.group.get
//
// 查询当前appkey在ONS中已有的分组
type TaobaoJushitaJmsGroupGetAPIResponse struct {
	model.CommonResponse
	TaobaoJushitaJmsGroupGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJushitaJmsGroupGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJushitaJmsGroupGetAPIResponseModel).Reset()
}

// TaobaoJushitaJmsGroupGetAPIResponseModel is 查询ONS分组 成功返回结果
type TaobaoJushitaJmsGroupGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jushita_jms_group_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分组信息
	Groups []MsgGroupDo `json:"groups,omitempty" xml:"groups>msg_group_do,omitempty"`
	// 返回的总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJushitaJmsGroupGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Groups = m.Groups[:0]
	m.TotalResults = 0
}

var poolTaobaoJushitaJmsGroupGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJushitaJmsGroupGetAPIResponse)
	},
}

// GetTaobaoJushitaJmsGroupGetAPIResponse 从 sync.Pool 获取 TaobaoJushitaJmsGroupGetAPIResponse
func GetTaobaoJushitaJmsGroupGetAPIResponse() *TaobaoJushitaJmsGroupGetAPIResponse {
	return poolTaobaoJushitaJmsGroupGetAPIResponse.Get().(*TaobaoJushitaJmsGroupGetAPIResponse)
}

// ReleaseTaobaoJushitaJmsGroupGetAPIResponse 将 TaobaoJushitaJmsGroupGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJushitaJmsGroupGetAPIResponse(v *TaobaoJushitaJmsGroupGetAPIResponse) {
	v.Reset()
	poolTaobaoJushitaJmsGroupGetAPIResponse.Put(v)
}
