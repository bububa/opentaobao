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
    _bureauName   string
    // 请求终端名称
    _terminalName   string
    // 终端类型：1005100-零售，1005200-医疗
    _terminalType   string
    // 调用方式：formal-正式、test-测试
    _invocation   string
    // 追溯码
    _code   string
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
func (r *AlibabaAlihealthDrugKytYbGetcoderelationRequest) SetBureauName(_bureauName string) error {
    r._bureauName = _bureauName
    r.Set("bureau_name", _bureauName)
    return nil
}

// BureauName Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetBureauName() string {
    return r._bureauName
}
// TerminalName Setter
// 请求终端名称
func (r *AlibabaAlihealthDrugKytYbGetcoderelationRequest) SetTerminalName(_terminalName string) error {
    r._terminalName = _terminalName
    r.Set("terminal_name", _terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetTerminalName() string {
    return r._terminalName
}
// TerminalType Setter
// 终端类型：1005100-零售，1005200-医疗
func (r *AlibabaAlihealthDrugKytYbGetcoderelationRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetTerminalType() string {
    return r._terminalType
}
// Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugKytYbGetcoderelationRequest) SetInvocation(_invocation string) error {
    r._invocation = _invocation
    r.Set("invocation", _invocation)
    return nil
}

// Invocation Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetInvocation() string {
    return r._invocation
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytYbGetcoderelationRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugKytYbGetcoderelationRequest) GetCode() string {
    return r._code
}
