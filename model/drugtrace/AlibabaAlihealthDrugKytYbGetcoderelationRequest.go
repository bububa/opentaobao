package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
医保-查询码的所有子码 API请求
alibaba.alihealth.drug.kyt.yb.getcoderelation

应用于药店或医院入库环节，通过扫码获取下级码进行入库；
通过码查询所有子码以及包装比例
*/
type AlibabaAlihealthDrugKytYbGetcoderelationRequest struct {
    model.Params
    // 社保局(所属地市名称)
    bureauName   string
    // 请求终端名称
    terminalName   string
    // 终端类型：1005100-零售，1005200-医疗
    terminalType   string
    // 调用方式：formal-正式、test-测试
    invocation   string
    // 追溯码
    code   string
}

// 初始化AlibabaAlihealthDrugKytYbGetcoderelationRequest对象
func NewAlibabaAlihealthDrugKytYbGetcoderelationRequest() *AlibabaAlihealthDrugKytYbGetcoderelationRequest{
    return &AlibabaAlihealthDrugKytYbGetcoderelationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yb.getcoderelation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BureauName Setter
// 社保局(所属地市名称)
func (r *AlibabaAlihealthDrugKytYbGetcoderelationRequest) SetBureauName(bureauName string) error {
    r.bureauName = bureauName
    r.Set("bureau_name", bureauName)
    return nil
}

// BureauName Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetBureauName() string {
    return r.bureauName
}
// TerminalName Setter
// 请求终端名称
func (r *AlibabaAlihealthDrugKytYbGetcoderelationRequest) SetTerminalName(terminalName string) error {
    r.terminalName = terminalName
    r.Set("terminal_name", terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetTerminalName() string {
    return r.terminalName
}
// TerminalType Setter
// 终端类型：1005100-零售，1005200-医疗
func (r *AlibabaAlihealthDrugKytYbGetcoderelationRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetTerminalType() string {
    return r.terminalType
}
// Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugKytYbGetcoderelationRequest) SetInvocation(invocation string) error {
    r.invocation = invocation
    r.Set("invocation", invocation)
    return nil
}

// Invocation Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetInvocation() string {
    return r.invocation
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytYbGetcoderelationRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetCode() string {
    return r.code
}
