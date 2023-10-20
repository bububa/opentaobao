package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDruguseQueryAPIRequest 合理用药规则查询 API请求
// alibaba.alihealth.druguse.query
//
// 查询用户购买的药品命中的风险规则
type AlibabaAlihealthDruguseQueryAPIRequest struct {
	model.Params
	// 入参
	_command *SafeMedicationReqCommand
}

// NewAlibabaAlihealthDruguseQueryRequest 初始化AlibabaAlihealthDruguseQueryAPIRequest对象
func NewAlibabaAlihealthDruguseQueryRequest() *AlibabaAlihealthDruguseQueryAPIRequest {
	return &AlibabaAlihealthDruguseQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDruguseQueryAPIRequest) Reset() {
	r._command = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDruguseQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.druguse.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDruguseQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDruguseQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommand is Command Setter
// 入参
func (r *AlibabaAlihealthDruguseQueryAPIRequest) SetCommand(_command *SafeMedicationReqCommand) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r AlibabaAlihealthDruguseQueryAPIRequest) GetCommand() *SafeMedicationReqCommand {
	return r._command
}

var poolAlibabaAlihealthDruguseQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDruguseQueryRequest()
	},
}

// GetAlibabaAlihealthDruguseQueryRequest 从 sync.Pool 获取 AlibabaAlihealthDruguseQueryAPIRequest
func GetAlibabaAlihealthDruguseQueryAPIRequest() *AlibabaAlihealthDruguseQueryAPIRequest {
	return poolAlibabaAlihealthDruguseQueryAPIRequest.Get().(*AlibabaAlihealthDruguseQueryAPIRequest)
}

// ReleaseAlibabaAlihealthDruguseQueryAPIRequest 将 AlibabaAlihealthDruguseQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDruguseQueryAPIRequest(v *AlibabaAlihealthDruguseQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDruguseQueryAPIRequest.Put(v)
}
