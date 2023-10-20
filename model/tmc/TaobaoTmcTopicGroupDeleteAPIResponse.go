package tmc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcTopicGroupDeleteAPIResponse 删除消息topic分组路由 API返回值
// taobao.tmc.topic.group.delete
//
// 删除根据topic名称路由消息到不同的分组关系
type TaobaoTmcTopicGroupDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoTmcTopicGroupDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmcTopicGroupDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmcTopicGroupDeleteAPIResponseModel).Reset()
}

// TaobaoTmcTopicGroupDeleteAPIResponseModel is 删除消息topic分组路由 成功返回结果
type TaobaoTmcTopicGroupDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_topic_group_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmcTopicGroupDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoTmcTopicGroupDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmcTopicGroupDeleteAPIResponse)
	},
}

// GetTaobaoTmcTopicGroupDeleteAPIResponse 从 sync.Pool 获取 TaobaoTmcTopicGroupDeleteAPIResponse
func GetTaobaoTmcTopicGroupDeleteAPIResponse() *TaobaoTmcTopicGroupDeleteAPIResponse {
	return poolTaobaoTmcTopicGroupDeleteAPIResponse.Get().(*TaobaoTmcTopicGroupDeleteAPIResponse)
}

// ReleaseTaobaoTmcTopicGroupDeleteAPIResponse 将 TaobaoTmcTopicGroupDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmcTopicGroupDeleteAPIResponse(v *TaobaoTmcTopicGroupDeleteAPIResponse) {
	v.Reset()
	poolTaobaoTmcTopicGroupDeleteAPIResponse.Put(v)
}
