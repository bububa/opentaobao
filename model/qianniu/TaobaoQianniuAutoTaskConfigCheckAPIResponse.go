package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuAutoTaskConfigCheckAPIResponse 自动化任务设置校验 API返回值
// taobao.qianniu.auto.task.config.check
//
// 校验自动化任务配置
type TaobaoQianniuAutoTaskConfigCheckAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuAutoTaskConfigCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuAutoTaskConfigCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuAutoTaskConfigCheckAPIResponseModel).Reset()
}

// TaobaoQianniuAutoTaskConfigCheckAPIResponseModel is 自动化任务设置校验 成功返回结果
type TaobaoQianniuAutoTaskConfigCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_auto_task_config_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	RpcErrorMessage string `json:"rpc_error_message,omitempty" xml:"rpc_error_message,omitempty"`
	// 错误码
	RpcErrorCode int64 `json:"rpc_error_code,omitempty" xml:"rpc_error_code,omitempty"`
	// 是否成功
	RpcSuccess bool `json:"rpc_success,omitempty" xml:"rpc_success,omitempty"`
	// success
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuAutoTaskConfigCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RpcErrorMessage = ""
	m.RpcErrorCode = 0
	m.RpcSuccess = false
	m.Module = false
}

var poolTaobaoQianniuAutoTaskConfigCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuAutoTaskConfigCheckAPIResponse)
	},
}

// GetTaobaoQianniuAutoTaskConfigCheckAPIResponse 从 sync.Pool 获取 TaobaoQianniuAutoTaskConfigCheckAPIResponse
func GetTaobaoQianniuAutoTaskConfigCheckAPIResponse() *TaobaoQianniuAutoTaskConfigCheckAPIResponse {
	return poolTaobaoQianniuAutoTaskConfigCheckAPIResponse.Get().(*TaobaoQianniuAutoTaskConfigCheckAPIResponse)
}

// ReleaseTaobaoQianniuAutoTaskConfigCheckAPIResponse 将 TaobaoQianniuAutoTaskConfigCheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuAutoTaskConfigCheckAPIResponse(v *TaobaoQianniuAutoTaskConfigCheckAPIResponse) {
	v.Reset()
	poolTaobaoQianniuAutoTaskConfigCheckAPIResponse.Put(v)
}
