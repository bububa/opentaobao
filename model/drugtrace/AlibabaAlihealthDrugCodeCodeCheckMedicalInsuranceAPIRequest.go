package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest 码核查状态同步-医保 API请求
// alibaba.alihealth.drug.code.code.check.medical.insurance
//
// 服务描述
// 核查平台在进行医保单据鉴证核查时，会记录单据中所有提交的追溯码信息；单据中的
// 追溯码包含所有正常和异常的数据；
// 此接口，针对正式鉴核的单据中提交的有效的、正常状态的追溯码，提供可由核查平台
// 发起，按单据鉴核时间顺序组织，向码上放心平台同步每笔单据核销的码状态信息；
// 入参采用数组方式提供，一次同步最多支持100条记录
type AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest struct {
	model.Params
	// 码列表
	_codes []string
	// 平台返回的终端id
	_terminalEntId string
	// 行政区域
	_bureauName string
	// 终端id
	_terminalId string
	// 终端类型（1005100-零售药店；1005200-医疗机构）
	_terminalType string
	// 核销类型(1012100：核销；1012900：退库)
	_cType string
}

// NewAlibabaalihealthdrugcodecodecheckmedicalinsuranceRequest 初始化AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest对象
func NewAlibabaalihealthdrugcodecodecheckmedicalinsuranceRequest() *AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest {
	return &AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.code.check.medical.insurance"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表
func (r *AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) GetCodes() []string {
	return r._codes
}

// SetTerminalEntId is TerminalEntId Setter
// 平台返回的终端id
func (r *AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) SetTerminalEntId(_terminalEntId string) error {
	r._terminalEntId = _terminalEntId
	r.Set("terminal_ent_id", _terminalEntId)
	return nil
}

// GetTerminalEntId TerminalEntId Getter
func (r AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) GetTerminalEntId() string {
	return r._terminalEntId
}

// SetBureauName is BureauName Setter
// 行政区域
func (r *AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) SetBureauName(_bureauName string) error {
	r._bureauName = _bureauName
	r.Set("bureau_name", _bureauName)
	return nil
}

// GetBureauName BureauName Getter
func (r AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) GetBureauName() string {
	return r._bureauName
}

// SetTerminalId is TerminalId Setter
// 终端id
func (r *AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) SetTerminalId(_terminalId string) error {
	r._terminalId = _terminalId
	r.Set("terminal_id", _terminalId)
	return nil
}

// GetTerminalId TerminalId Getter
func (r AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) GetTerminalId() string {
	return r._terminalId
}

// SetTerminalType is TerminalType Setter
// 终端类型（1005100-零售药店；1005200-医疗机构）
func (r *AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetCType is CType Setter
// 核销类型(1012100：核销；1012900：退库)
func (r *AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) SetCType(_cType string) error {
	r._cType = _cType
	r.Set("c_type", _cType)
	return nil
}

// GetCType CType Getter
func (r AlibabaalihealthdrugcodecodecheckmedicalinsuranceAPIRequest) GetCType() string {
	return r._cType
}
