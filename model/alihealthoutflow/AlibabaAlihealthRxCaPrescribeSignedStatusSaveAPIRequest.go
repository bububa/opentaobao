package alihealthoutflow

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.rx.ca.prescribe.signed.status.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthRxCaPrescribeSignedStatusSaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
