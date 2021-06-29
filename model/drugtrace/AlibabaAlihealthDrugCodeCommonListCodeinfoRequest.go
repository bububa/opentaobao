package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通用查询码接口 API请求
alibaba.alihealth.drug.code.common.list.codeinfo

通用查询码接口
*/
type AlibabaAlihealthDrugCodeCommonListCodeinfoRequest struct {
    model.Params
    // 企业refEntId
    _refEntId   string
    // 标示医院业务
    _searchSource   string
    // 追溯码
    _codeList   []string
    // 证件编号
    _certIsvNo   string
    // 调用方式：formal-正式、test-测试
    _invocation   string
    // 终端类型 1：零售
    _terminalType   string
    // 调用零售药店名称
    _terminalName   string
    // 城市名称
    _bureauName   string
    // 错误信息
    _errorMessage   string
    // 验证权限企业id
    _authRefEntId   string
}

// 初始化AlibabaAlihealthDrugCodeCommonListCodeinfoRequest对象
func NewAlibabaAlihealthDrugCodeCommonListCodeinfoRequest() *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest{
    return &AlibabaAlihealthDrugCodeCommonListCodeinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.common.list.codeinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业refEntId
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetRefEntId() string {
    return r._refEntId
}
// SearchSource Setter
// 标示医院业务
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetSearchSource(_searchSource string) error {
    r._searchSource = _searchSource
    r.Set("search_source", _searchSource)
    return nil
}

// SearchSource Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetSearchSource() string {
    return r._searchSource
}
// CodeList Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetCodeList(_codeList []string) error {
    r._codeList = _codeList
    r.Set("code_list", _codeList)
    return nil
}

// CodeList Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetCodeList() []string {
    return r._codeList
}
// CertIsvNo Setter
// 证件编号
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetCertIsvNo(_certIsvNo string) error {
    r._certIsvNo = _certIsvNo
    r.Set("cert_isv_no", _certIsvNo)
    return nil
}

// CertIsvNo Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetCertIsvNo() string {
    return r._certIsvNo
}
// Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetInvocation(_invocation string) error {
    r._invocation = _invocation
    r.Set("invocation", _invocation)
    return nil
}

// Invocation Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetInvocation() string {
    return r._invocation
}
// TerminalType Setter
// 终端类型 1：零售
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetTerminalType() string {
    return r._terminalType
}
// TerminalName Setter
// 调用零售药店名称
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetTerminalName(_terminalName string) error {
    r._terminalName = _terminalName
    r.Set("terminal_name", _terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetTerminalName() string {
    return r._terminalName
}
// BureauName Setter
// 城市名称
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetBureauName(_bureauName string) error {
    r._bureauName = _bureauName
    r.Set("bureau_name", _bureauName)
    return nil
}

// BureauName Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetBureauName() string {
    return r._bureauName
}
// ErrorMessage Setter
// 错误信息
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetErrorMessage(_errorMessage string) error {
    r._errorMessage = _errorMessage
    r.Set("error_message", _errorMessage)
    return nil
}

// ErrorMessage Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetErrorMessage() string {
    return r._errorMessage
}
// AuthRefEntId Setter
// 验证权限企业id
func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetAuthRefEntId(_authRefEntId string) error {
    r._authRefEntId = _authRefEntId
    r.Set("auth_ref_ent_id", _authRefEntId)
    return nil
}

// AuthRefEntId Getter
func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetAuthRefEntId() string {
    return r._authRefEntId
}
