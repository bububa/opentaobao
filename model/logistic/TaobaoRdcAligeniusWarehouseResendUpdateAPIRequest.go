package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniuswarehouseresendupdateAPIRequest 补发单状态回传 API请求
// taobao.rdc.aligenius.warehouse.resend.update
//
// 补发单状态回传接口
type TaobaordcaligeniuswarehouseresendupdateAPIRequest struct {
	model.Params
	// 参数
	_param0 *UpdateResendStatusDto
}

// NewTaobaordcaligeniuswarehouseresendupdateRequest 初始化TaobaordcaligeniuswarehouseresendupdateAPIRequest对象
func NewTaobaordcaligeniuswarehouseresendupdateRequest() *TaobaordcaligeniuswarehouseresendupdateAPIRequest {
	return &TaobaordcaligeniuswarehouseresendupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordcaligeniuswarehouseresendupdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.warehouse.resend.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordcaligeniuswarehouseresendupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordcaligeniuswarehouseresendupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *TaobaordcaligeniuswarehouseresendupdateAPIRequest) SetParam0(_param0 *UpdateResendStatusDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaordcaligeniuswarehouseresendupdateAPIRequest) GetParam0() *UpdateResendStatusDto {
	return r._param0
}
