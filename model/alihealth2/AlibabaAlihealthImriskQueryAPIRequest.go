package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthImriskQueryAPIRequest 问诊质控接口 API请求
// alibaba.alihealth.imrisk.query
//
// 阿里健康的问诊质控接口
type AlibabaAlihealthImriskQueryAPIRequest struct {
	model.Params
	// 入参数
	_param0 *IMRiskCheckCommand
}

// NewAlibabaAlihealthImriskQueryRequest 初始化AlibabaAlihealthImriskQueryAPIRequest对象
func NewAlibabaAlihealthImriskQueryRequest() *AlibabaAlihealthImriskQueryAPIRequest {
	return &AlibabaAlihealthImriskQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthImriskQueryAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthImriskQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.imrisk.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthImriskQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthImriskQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参数
func (r *AlibabaAlihealthImriskQueryAPIRequest) SetParam0(_param0 *IMRiskCheckCommand) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaAlihealthImriskQueryAPIRequest) GetParam0() *IMRiskCheckCommand {
	return r._param0
}

var poolAlibabaAlihealthImriskQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthImriskQueryRequest()
	},
}

// GetAlibabaAlihealthImriskQueryRequest 从 sync.Pool 获取 AlibabaAlihealthImriskQueryAPIRequest
func GetAlibabaAlihealthImriskQueryAPIRequest() *AlibabaAlihealthImriskQueryAPIRequest {
	return poolAlibabaAlihealthImriskQueryAPIRequest.Get().(*AlibabaAlihealthImriskQueryAPIRequest)
}

// ReleaseAlibabaAlihealthImriskQueryAPIRequest 将 AlibabaAlihealthImriskQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthImriskQueryAPIRequest(v *AlibabaAlihealthImriskQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthImriskQueryAPIRequest.Put(v)
}
