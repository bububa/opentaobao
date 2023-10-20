package jstinteractive

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveTaskRegisterAPIResponse 互动任务开通接口 API返回值
// taobao.jst.interactive.task.register
//
// 调用互动任务开通接口为小程序开通互动任务
type TaobaoJstInteractiveTaskRegisterAPIResponse struct {
	model.CommonResponse
	TaobaoJstInteractiveTaskRegisterAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstInteractiveTaskRegisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstInteractiveTaskRegisterAPIResponseModel).Reset()
}

// TaobaoJstInteractiveTaskRegisterAPIResponseModel is 互动任务开通接口 成功返回结果
type TaobaoJstInteractiveTaskRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_task_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstInteractiveTaskRegisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoJstInteractiveTaskRegisterAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstInteractiveTaskRegisterAPIResponse)
	},
}

// GetTaobaoJstInteractiveTaskRegisterAPIResponse 从 sync.Pool 获取 TaobaoJstInteractiveTaskRegisterAPIResponse
func GetTaobaoJstInteractiveTaskRegisterAPIResponse() *TaobaoJstInteractiveTaskRegisterAPIResponse {
	return poolTaobaoJstInteractiveTaskRegisterAPIResponse.Get().(*TaobaoJstInteractiveTaskRegisterAPIResponse)
}

// ReleaseTaobaoJstInteractiveTaskRegisterAPIResponse 将 TaobaoJstInteractiveTaskRegisterAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstInteractiveTaskRegisterAPIResponse(v *TaobaoJstInteractiveTaskRegisterAPIResponse) {
	v.Reset()
	poolTaobaoJstInteractiveTaskRegisterAPIResponse.Put(v)
}
