package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码查询码信息-黄麻碱 API请求
alibaba.alihealth.drug.code.list.code.alkali

服务描述
黄麻碱鉴证核查是基于在两定机构的药品管理（入库、出库或盘点）环节，增加扫码匹配
与验证鉴核流程；
此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息；
若所传的监管码是非最小包装监管码，且存在药品混包的情况，则此接口不支持。这种
情况下，需要分多次调用该接口。
核查平台优先过滤非8开头的，长度非20位数字的码信息。
*/
type AlibabaAlihealthDrugCodeListCodeAlkaliRequest struct {
    model.Params
    // 追溯码
    codeList   []string
    // 证件编号
    certIsvNo   string
    // 调用方式：formal-正式、test-测试
    invocation   string
    // 终端类型 1：零售
    terminalType   string
    // 调用零售药店名称
    terminalName   string
    // 城市名称
    bureauName   string
    // 错误信息
    errorMessage   string
}

// 初始化AlibabaAlihealthDrugCodeListCodeAlkaliRequest对象
func NewAlibabaAlihealthDrugCodeListCodeAlkaliRequest() *AlibabaAlihealthDrugCodeListCodeAlkaliRequest{
    return &AlibabaAlihealthDrugCodeListCodeAlkaliRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeListCodeAlkaliRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.list.code.alkali"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeListCodeAlkaliRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CodeList Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeListCodeAlkaliRequest) SetCodeList(codeList []string) error {
    r.codeList = codeList
    r.Set("code_list", codeList)
    return nil
}

// CodeList Getter
func (r AlibabaAlihealthDrugCodeListCodeAlkaliRequest) GetCodeList() []string {
    return r.codeList
}
// CertIsvNo Setter
// 证件编号
func (r *AlibabaAlihealthDrugCodeListCodeAlkaliRequest) SetCertIsvNo(certIsvNo string) error {
    r.certIsvNo = certIsvNo
    r.Set("cert_isv_no", certIsvNo)
    return nil
}

// CertIsvNo Getter
func (r AlibabaAlihealthDrugCodeListCodeAlkaliRequest) GetCertIsvNo() string {
    return r.certIsvNo
}
// Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugCodeListCodeAlkaliRequest) SetInvocation(invocation string) error {
    r.invocation = invocation
    r.Set("invocation", invocation)
    return nil
}

// Invocation Getter
func (r AlibabaAlihealthDrugCodeListCodeAlkaliRequest) GetInvocation() string {
    return r.invocation
}
// TerminalType Setter
// 终端类型 1：零售
func (r *AlibabaAlihealthDrugCodeListCodeAlkaliRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r AlibabaAlihealthDrugCodeListCodeAlkaliRequest) GetTerminalType() string {
    return r.terminalType
}
// TerminalName Setter
// 调用零售药店名称
func (r *AlibabaAlihealthDrugCodeListCodeAlkaliRequest) SetTerminalName(terminalName string) error {
    r.terminalName = terminalName
    r.Set("terminal_name", terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugCodeListCodeAlkaliRequest) GetTerminalName() string {
    return r.terminalName
}
// BureauName Setter
// 城市名称
func (r *AlibabaAlihealthDrugCodeListCodeAlkaliRequest) SetBureauName(bureauName string) error {
    r.bureauName = bureauName
    r.Set("bureau_name", bureauName)
    return nil
}

// BureauName Getter
func (r AlibabaAlihealthDrugCodeListCodeAlkaliRequest) GetBureauName() string {
    return r.bureauName
}
// ErrorMessage Setter
// 错误信息
func (r *AlibabaAlihealthDrugCodeListCodeAlkaliRequest) SetErrorMessage(errorMessage string) error {
    r.errorMessage = errorMessage
    r.Set("error_message", errorMessage)
    return nil
}

// ErrorMessage Getter
func (r AlibabaAlihealthDrugCodeListCodeAlkaliRequest) GetErrorMessage() string {
    return r.errorMessage
}
