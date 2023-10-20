package alihealthoutflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest 处方ca认证 API请求
// alibaba.alihealth.rx.ca.prescribe.signed.status.save
//
// 处方ca认证
type AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest struct {
	model.Params
	// 入参实体
	_param0 *PrescribeStatusDto
}

// NewAlibabaAlihealthRxCaPrescribeSignedStatusSaveRequest 初始化AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest对象
func NewAlibabaAlihealthRxCaPrescribeSignedStatusSaveRequest() *AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest {
	return &AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.rx.ca.prescribe.signed.status.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参实体
func (r *AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest) SetParam0(_param0 *PrescribeStatusDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest) GetParam0() *PrescribeStatusDto {
	return r._param0
}

var poolAlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthRxCaPrescribeSignedStatusSaveRequest()
	},
}

// GetAlibabaAlihealthRxCaPrescribeSignedStatusSaveRequest 从 sync.Pool 获取 AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest
func GetAlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest() *AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest {
	return poolAlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest.Get().(*AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest)
}

// ReleaseAlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest 将 AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest(v *AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest.Put(v)
}
