package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleOnespuRegisterUpdateAPIRequest 闲鱼 ONESPU 挂载接口 API请求
// alibaba.idle.onespu.register.update
//
// 闲鱼 ONESPU 挂载接口
type AlibabaIdleOnespuRegisterUpdateAPIRequest struct {
	model.Params
	// 入参
	_oneSpuSpRegisterUpdateParam *OneSpuSpRegisterUpdateParam
}

// NewAlibabaIdleOnespuRegisterUpdateRequest 初始化AlibabaIdleOnespuRegisterUpdateAPIRequest对象
func NewAlibabaIdleOnespuRegisterUpdateRequest() *AlibabaIdleOnespuRegisterUpdateAPIRequest {
	return &AlibabaIdleOnespuRegisterUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleOnespuRegisterUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.onespu.register.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleOnespuRegisterUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleOnespuRegisterUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOneSpuSpRegisterUpdateParam is OneSpuSpRegisterUpdateParam Setter
// 入参
func (r *AlibabaIdleOnespuRegisterUpdateAPIRequest) SetOneSpuSpRegisterUpdateParam(_oneSpuSpRegisterUpdateParam *OneSpuSpRegisterUpdateParam) error {
	r._oneSpuSpRegisterUpdateParam = _oneSpuSpRegisterUpdateParam
	r.Set("one_spu_sp_register_update_param", _oneSpuSpRegisterUpdateParam)
	return nil
}

// GetOneSpuSpRegisterUpdateParam OneSpuSpRegisterUpdateParam Getter
func (r AlibabaIdleOnespuRegisterUpdateAPIRequest) GetOneSpuSpRegisterUpdateParam() *OneSpuSpRegisterUpdateParam {
	return r._oneSpuSpRegisterUpdateParam
}
