package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodelistcodealkaliAPIRequest 根据码查询码信息-黄麻碱 API请求
// alibaba.alihealth.drug.code.list.code.alkali
//
// 服务描述
// 黄麻碱鉴证核查是基于在两定机构的药品管理（入库、出库或盘点）环节，增加扫码匹配
// 与验证鉴核流程；
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息；
// 若所传的监管码是非最小包装监管码，且存在药品混包的情况，则此接口不支持。这种
// 情况下，需要分多次调用该接口。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaalihealthdrugcodelistcodealkaliAPIRequest struct {
	model.Params
	// 追溯码
	_codeList []string
	// 证件编号
	_certIsvNo string
	// 调用方式：formal-正式、test-测试
	_invocation string
	// 终端类型 1：零售
	_terminalType string
	// 调用零售药店名称
	_terminalName string
	// 城市名称
	_bureauName string
	// 错误信息
	_errorMessage string
}

// NewAlibabaalihealthdrugcodelistcodealkaliRequest 初始化AlibabaalihealthdrugcodelistcodealkaliAPIRequest对象
func NewAlibabaalihealthdrugcodelistcodealkaliRequest() *AlibabaalihealthdrugcodelistcodealkaliAPIRequest {
	return &AlibabaalihealthdrugcodelistcodealkaliAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodelistcodealkaliAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.list.code.alkali"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodelistcodealkaliAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodelistcodealkaliAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodeList is CodeList Setter
// 追溯码
func (r *AlibabaalihealthdrugcodelistcodealkaliAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// GetCodeList CodeList Getter
func (r AlibabaalihealthdrugcodelistcodealkaliAPIRequest) GetCodeList() []string {
	return r._codeList
}

// SetCertIsvNo is CertIsvNo Setter
// 证件编号
func (r *AlibabaalihealthdrugcodelistcodealkaliAPIRequest) SetCertIsvNo(_certIsvNo string) error {
	r._certIsvNo = _certIsvNo
	r.Set("cert_isv_no", _certIsvNo)
	return nil
}

// GetCertIsvNo CertIsvNo Getter
func (r AlibabaalihealthdrugcodelistcodealkaliAPIRequest) GetCertIsvNo() string {
	return r._certIsvNo
}

// SetInvocation is Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaalihealthdrugcodelistcodealkaliAPIRequest) SetInvocation(_invocation string) error {
	r._invocation = _invocation
	r.Set("invocation", _invocation)
	return nil
}

// GetInvocation Invocation Getter
func (r AlibabaalihealthdrugcodelistcodealkaliAPIRequest) GetInvocation() string {
	return r._invocation
}

// SetTerminalType is TerminalType Setter
// 终端类型 1：零售
func (r *AlibabaalihealthdrugcodelistcodealkaliAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r AlibabaalihealthdrugcodelistcodealkaliAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetTerminalName is TerminalName Setter
// 调用零售药店名称
func (r *AlibabaalihealthdrugcodelistcodealkaliAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// GetTerminalName TerminalName Getter
func (r AlibabaalihealthdrugcodelistcodealkaliAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// SetBureauName is BureauName Setter
// 城市名称
func (r *AlibabaalihealthdrugcodelistcodealkaliAPIRequest) SetBureauName(_bureauName string) error {
	r._bureauName = _bureauName
	r.Set("bureau_name", _bureauName)
	return nil
}

// GetBureauName BureauName Getter
func (r AlibabaalihealthdrugcodelistcodealkaliAPIRequest) GetBureauName() string {
	return r._bureauName
}

// SetErrorMessage is ErrorMessage Setter
// 错误信息
func (r *AlibabaalihealthdrugcodelistcodealkaliAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// GetErrorMessage ErrorMessage Getter
func (r AlibabaalihealthdrugcodelistcodealkaliAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}
