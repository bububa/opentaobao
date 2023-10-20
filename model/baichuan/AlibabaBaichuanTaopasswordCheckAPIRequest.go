package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibababaichuantaopasswordcheckAPIRequest 淘口令检查 API请求
// alibaba.baichuan.taopassword.check
//
// 检查当前文本是否为淘口令
type AlibababaichuantaopasswordcheckAPIRequest struct {
	model.Params
	// 参数DTO
	_paramDto *ParamDto
	// 系统自动生成
	_clientInfo *RichClientInfo
}

// NewAlibababaichuantaopasswordcheckRequest 初始化AlibababaichuantaopasswordcheckAPIRequest对象
func NewAlibababaichuantaopasswordcheckRequest() *AlibababaichuantaopasswordcheckAPIRequest {
	return &AlibababaichuantaopasswordcheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibababaichuantaopasswordcheckAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.taopassword.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibababaichuantaopasswordcheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibababaichuantaopasswordcheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamDto is ParamDto Setter
// 参数DTO
func (r *AlibababaichuantaopasswordcheckAPIRequest) SetParamDto(_paramDto *ParamDto) error {
	r._paramDto = _paramDto
	r.Set("param_dto", _paramDto)
	return nil
}

// GetParamDto ParamDto Getter
func (r AlibababaichuantaopasswordcheckAPIRequest) GetParamDto() *ParamDto {
	return r._paramDto
}

// SetClientInfo is ClientInfo Setter
// 系统自动生成
func (r *AlibababaichuantaopasswordcheckAPIRequest) SetClientInfo(_clientInfo *RichClientInfo) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r AlibababaichuantaopasswordcheckAPIRequest) GetClientInfo() *RichClientInfo {
	return r._clientInfo
}
