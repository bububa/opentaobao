package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanTaopasswordCheckAPIRequest 淘口令检查 API请求
// alibaba.baichuan.taopassword.check
//
// 检查当前文本是否为淘口令
type AlibabaBaichuanTaopasswordCheckAPIRequest struct {
	model.Params
	// 参数DTO
	_paramDto *ParamDto
	// 系统自动生成
	_clientInfo *RichClientInfo
}

// NewAlibabaBaichuanTaopasswordCheckRequest 初始化AlibabaBaichuanTaopasswordCheckAPIRequest对象
func NewAlibabaBaichuanTaopasswordCheckRequest() *AlibabaBaichuanTaopasswordCheckAPIRequest {
	return &AlibabaBaichuanTaopasswordCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanTaopasswordCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.taopassword.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanTaopasswordCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamDto is ParamDto Setter
// 参数DTO
func (r *AlibabaBaichuanTaopasswordCheckAPIRequest) SetParamDto(_paramDto *ParamDto) error {
	r._paramDto = _paramDto
	r.Set("param_dto", _paramDto)
	return nil
}

// GetParamDto ParamDto Getter
func (r AlibabaBaichuanTaopasswordCheckAPIRequest) GetParamDto() *ParamDto {
	return r._paramDto
}

// SetClientInfo is ClientInfo Setter
// 系统自动生成
func (r *AlibabaBaichuanTaopasswordCheckAPIRequest) SetClientInfo(_clientInfo *RichClientInfo) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r AlibabaBaichuanTaopasswordCheckAPIRequest) GetClientInfo() *RichClientInfo {
	return r._clientInfo
}
