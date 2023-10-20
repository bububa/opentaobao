package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodecodecheckhospitalAPIRequest 码核查状态同步-医院 API请求
// alibaba.alihealth.drug.code.code.check.hospital
//
// 码核查状态同步-医院
type AlibabaalihealthdrugcodecodecheckhospitalAPIRequest struct {
	model.Params
	// 码列表
	_codes []string
	// 认证企业refEntId
	_authRefEntId string
	// 企业refEntId
	_refEntId string
	// 城市名
	_bureauName string
	// 终端名称
	_terminalName string
	// 终端类型
	_terminalType string
	// 核销类型
	_cType string
}

// NewAlibabaalihealthdrugcodecodecheckhospitalRequest 初始化AlibabaalihealthdrugcodecodecheckhospitalAPIRequest对象
func NewAlibabaalihealthdrugcodecodecheckhospitalRequest() *AlibabaalihealthdrugcodecodecheckhospitalAPIRequest {
	return &AlibabaalihealthdrugcodecodecheckhospitalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.code.check.hospital"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表
func (r *AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) GetCodes() []string {
	return r._codes
}

// SetAuthRefEntId is AuthRefEntId Setter
// 认证企业refEntId
func (r *AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) SetAuthRefEntId(_authRefEntId string) error {
	r._authRefEntId = _authRefEntId
	r.Set("auth_ref_ent_id", _authRefEntId)
	return nil
}

// GetAuthRefEntId AuthRefEntId Getter
func (r AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) GetAuthRefEntId() string {
	return r._authRefEntId
}

// SetRefEntId is RefEntId Setter
// 企业refEntId
func (r *AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBureauName is BureauName Setter
// 城市名
func (r *AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) SetBureauName(_bureauName string) error {
	r._bureauName = _bureauName
	r.Set("bureau_name", _bureauName)
	return nil
}

// GetBureauName BureauName Getter
func (r AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) GetBureauName() string {
	return r._bureauName
}

// SetTerminalName is TerminalName Setter
// 终端名称
func (r *AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// GetTerminalName TerminalName Getter
func (r AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// SetTerminalType is TerminalType Setter
// 终端类型
func (r *AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetCType is CType Setter
// 核销类型
func (r *AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) SetCType(_cType string) error {
	r._cType = _cType
	r.Set("c_type", _cType)
	return nil
}

// GetCType CType Getter
func (r AlibabaalihealthdrugcodecodecheckhospitalAPIRequest) GetCType() string {
	return r._cType
}
