package jstinteractive

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveActivityUpdateAPIResponse 互动任务活动修改接口 API返回值
// taobao.jst.interactive.activity.update
//
// 互动任务活动修改接口
type TaobaoJstInteractiveActivityUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoJstInteractiveActivityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstInteractiveActivityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstInteractiveActivityUpdateAPIResponseModel).Reset()
}

// TaobaoJstInteractiveActivityUpdateAPIResponseModel is 互动任务活动修改接口 成功返回结果
type TaobaoJstInteractiveActivityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstInteractiveActivityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoJstInteractiveActivityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstInteractiveActivityUpdateAPIResponse)
	},
}

// GetTaobaoJstInteractiveActivityUpdateAPIResponse 从 sync.Pool 获取 TaobaoJstInteractiveActivityUpdateAPIResponse
func GetTaobaoJstInteractiveActivityUpdateAPIResponse() *TaobaoJstInteractiveActivityUpdateAPIResponse {
	return poolTaobaoJstInteractiveActivityUpdateAPIResponse.Get().(*TaobaoJstInteractiveActivityUpdateAPIResponse)
}

// ReleaseTaobaoJstInteractiveActivityUpdateAPIResponse 将 TaobaoJstInteractiveActivityUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstInteractiveActivityUpdateAPIResponse(v *TaobaoJstInteractiveActivityUpdateAPIResponse) {
	v.Reset()
	poolTaobaoJstInteractiveActivityUpdateAPIResponse.Put(v)
}
