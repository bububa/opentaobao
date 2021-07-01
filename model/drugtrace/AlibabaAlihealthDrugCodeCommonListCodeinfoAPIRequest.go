package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest
通用查询码接口 API请求
alibaba.alihealth.drug.code.common.list.codeinfo

通用查询码接口 */
type AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest struct {
	model.Params
	// 企业refEntId
	_refEntId string
	// 标示医院业务
	_searchSource string
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
	// 验证权限企业id
	_authRefEntId string
}

// NewAlibabaAlihealthDrugCodeCommonListCodeinfoRequest 初始化AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest对象
func NewAlibabaAlihealthDrugCodeCommonListCodeinfoRequest() *AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest {
	return &AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.common.list.codeinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业refEntId
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is SearchSource Setter
// 标示医院业务
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) SetSearchSource(_searchSource string) error {
	r._searchSource = _searchSource
	r.Set("search_source", _searchSource)
	return nil
}

// Get SearchSource Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetSearchSource() string {
	return r._searchSource
}

// Set is CodeList Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// Get CodeList Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetCodeList() []string {
	return r._codeList
}

// Set is CertIsvNo Setter
// 证件编号
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) SetCertIsvNo(_certIsvNo string) error {
	r._certIsvNo = _certIsvNo
	r.Set("cert_isv_no", _certIsvNo)
	return nil
}

// Get CertIsvNo Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetCertIsvNo() string {
	return r._certIsvNo
}

// Set is Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) SetInvocation(_invocation string) error {
	r._invocation = _invocation
	r.Set("invocation", _invocation)
	return nil
}

// Get Invocation Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetInvocation() string {
	return r._invocation
}

// Set is TerminalType Setter
// 终端类型 1：零售
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// Get TerminalType Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// Set is TerminalName Setter
// 调用零售药店名称
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// Get TerminalName Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// Set is BureauName Setter
// 城市名称
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) SetBureauName(_bureauName string) error {
	r._bureauName = _bureauName
	r.Set("bureau_name", _bureauName)
	return nil
}

// Get BureauName Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetBureauName() string {
	return r._bureauName
}

// Set is ErrorMessage Setter
// 错误信息
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// Get ErrorMessage Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// Set is AuthRefEntId Setter
// 验证权限企业id
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) SetAuthRefEntId(_authRefEntId string) error {
	r._authRefEntId = _authRefEntId
	r.Set("auth_ref_ent_id", _authRefEntId)
	return nil
}

// Get AuthRefEntId Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest) GetAuthRefEntId() string {
	return r._authRefEntId
}
