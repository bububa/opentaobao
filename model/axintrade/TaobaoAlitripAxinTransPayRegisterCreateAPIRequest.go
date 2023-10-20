package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintranspayregistercreateAPIRequest 提交支付服务开通 API请求
// taobao.alitrip.axin.trans.pay.register.create
//
// 阿信供销平台-提交支付服务开通接口
type TaobaoalitripaxintranspayregistercreateAPIRequest struct {
	model.Params
	// 提交支付服务开通接口入参
	_createDTO *AxinPayRegisterCreateDto
}

// NewTaobaoalitripaxintranspayregistercreateRequest 初始化TaobaoalitripaxintranspayregistercreateAPIRequest对象
func NewTaobaoalitripaxintranspayregistercreateRequest() *TaobaoalitripaxintranspayregistercreateAPIRequest {
	return &TaobaoalitripaxintranspayregistercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripaxintranspayregistercreateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.pay.register.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripaxintranspayregistercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripaxintranspayregistercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateDTO is CreateDTO Setter
// 提交支付服务开通接口入参
func (r *TaobaoalitripaxintranspayregistercreateAPIRequest) SetCreateDTO(_createDTO *AxinPayRegisterCreateDto) error {
	r._createDTO = _createDTO
	r.Set("create_d_t_o", _createDTO)
	return nil
}

// GetCreateDTO CreateDTO Getter
func (r TaobaoalitripaxintranspayregistercreateAPIRequest) GetCreateDTO() *AxinPayRegisterCreateDto {
	return r._createDTO
}
