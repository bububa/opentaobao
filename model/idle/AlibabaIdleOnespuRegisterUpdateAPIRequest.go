package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleonespuregisterupdateAPIRequest 闲鱼 ONESPU 挂载接口 API请求
// alibaba.idle.onespu.register.update
//
// 闲鱼 ONESPU 挂载接口
type AlibabaidleonespuregisterupdateAPIRequest struct {
	model.Params
	// 入参
	_oneSpuSpRegisterUpdateParam *OneSpuSpRegisterUpdateParam
}

// NewAlibabaidleonespuregisterupdateRequest 初始化AlibabaidleonespuregisterupdateAPIRequest对象
func NewAlibabaidleonespuregisterupdateRequest() *AlibabaidleonespuregisterupdateAPIRequest {
	return &AlibabaidleonespuregisterupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleonespuregisterupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.onespu.register.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleonespuregisterupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleonespuregisterupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOneSpuSpRegisterUpdateParam is OneSpuSpRegisterUpdateParam Setter
// 入参
func (r *AlibabaidleonespuregisterupdateAPIRequest) SetOneSpuSpRegisterUpdateParam(_oneSpuSpRegisterUpdateParam *OneSpuSpRegisterUpdateParam) error {
	r._oneSpuSpRegisterUpdateParam = _oneSpuSpRegisterUpdateParam
	r.Set("one_spu_sp_register_update_param", _oneSpuSpRegisterUpdateParam)
	return nil
}

// GetOneSpuSpRegisterUpdateParam OneSpuSpRegisterUpdateParam Getter
func (r AlibabaidleonespuregisterupdateAPIRequest) GetOneSpuSpRegisterUpdateParam() *OneSpuSpRegisterUpdateParam {
	return r._oneSpuSpRegisterUpdateParam
}
