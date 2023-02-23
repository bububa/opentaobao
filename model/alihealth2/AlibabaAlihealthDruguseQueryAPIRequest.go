package alihealth2

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
