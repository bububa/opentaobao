package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopAuthGetAPIResponse 登陆 API返回值
// taobao.ailab.aicloud.top.auth.get
//
// 登陆
type TaobaoAilabAicloudTopAuthGetAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopAuthGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopAuthGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopAuthGetAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopAuthGetAPIResponseModel is 登陆 成功返回结果
type TaobaoAilabAicloudTopAuthGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_auth_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopAuthGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopAuthGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopAuthGetAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopAuthGetAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopAuthGetAPIResponse
func GetTaobaoAilabAicloudTopAuthGetAPIResponse() *TaobaoAilabAicloudTopAuthGetAPIResponse {
	return poolTaobaoAilabAicloudTopAuthGetAPIResponse.Get().(*TaobaoAilabAicloudTopAuthGetAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopAuthGetAPIResponse 将 TaobaoAilabAicloudTopAuthGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopAuthGetAPIResponse(v *TaobaoAilabAicloudTopAuthGetAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopAuthGetAPIResponse.Put(v)
}
