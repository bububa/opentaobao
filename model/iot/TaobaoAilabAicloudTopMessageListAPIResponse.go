package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessageListAPIResponse 获取留言列表 API返回值
// taobao.ailab.aicloud.top.message.list
//
// 根据指定参数获取留言列表
type TaobaoAilabAicloudTopMessageListAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMessageListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessageListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopMessageListAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopMessageListAPIResponseModel is 获取留言列表 成功返回结果
type TaobaoAilabAicloudTopMessageListAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_message_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessageListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopMessageListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMessageListAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopMessageListAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopMessageListAPIResponse
func GetTaobaoAilabAicloudTopMessageListAPIResponse() *TaobaoAilabAicloudTopMessageListAPIResponse {
	return poolTaobaoAilabAicloudTopMessageListAPIResponse.Get().(*TaobaoAilabAicloudTopMessageListAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopMessageListAPIResponse 将 TaobaoAilabAicloudTopMessageListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopMessageListAPIResponse(v *TaobaoAilabAicloudTopMessageListAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopMessageListAPIResponse.Put(v)
}
