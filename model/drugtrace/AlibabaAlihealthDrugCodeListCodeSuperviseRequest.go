package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码获取码信息-监管 API请求
alibaba.alihealth.drug.code.list.code.supervise

服务描述
医保鉴证核查是基于在两定机构的药品管理（入库、出库或盘点）环节，增加扫码匹配
与验证鉴核流程；
此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息；
若所传的监管码是非最小包装监管码，且存在药品混包的情况，则此接口不支持。这种
情况下，需要分多次调用该接口。
核查平台优先过滤非8开头的，长度非20位数字的码信息。
*/
type AlibabaAlihealthDrugCodeListCodeSuperviseRequest struct {
    model.Params
    // 追溯码
    _codeList   []string
    // ISV开放平台帐号标识
    _certIsvNo   string
    // 调用方式：formal-正式、test-测试
    _invocation   string
    // 终端类型 1005100-零售药店 ；10052-医疗机构    10053-建行智能pos机，10054-建行裕农通，10055-建行手机银行，10056-建行龙易行，10057-建行APP，10058-建行自助设备，10059-建行扫码设备
    _terminalType   string
    // 调用零售药店名称
    _terminalName   string
    // 门店所属的行政区域ID
    _bureauId   string
    // 零售终端id
    _terminalEntId   string
}

// 初始化AlibabaAlihealthDrugCodeListCodeSuperviseRequest对象
func NewAlibabaAlihealthDrugCodeListCodeSuperviseRequest() *AlibabaAlihealthDrugCodeListCodeSuperviseRequest{
    return &AlibabaAlihealthDrugCodeListCodeSuperviseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeListCodeSuperviseRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.list.code.supervise"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeListCodeSuperviseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CodeList Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseRequest) SetCodeList(_codeList []string) error {
    r._codeList = _codeList
    r.Set("code_list", _codeList)
    return nil
}

// CodeList Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseRequest) GetCodeList() []string {
    return r._codeList
}
// CertIsvNo Setter
// ISV开放平台帐号标识
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseRequest) SetCertIsvNo(_certIsvNo string) error {
    r._certIsvNo = _certIsvNo
    r.Set("cert_isv_no", _certIsvNo)
    return nil
}

// CertIsvNo Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseRequest) GetCertIsvNo() string {
    return r._certIsvNo
}
// Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseRequest) SetInvocation(_invocation string) error {
    r._invocation = _invocation
    r.Set("invocation", _invocation)
    return nil
}

// Invocation Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseRequest) GetInvocation() string {
    return r._invocation
}
// TerminalType Setter
// 终端类型 1005100-零售药店 ；10052-医疗机构    10053-建行智能pos机，10054-建行裕农通，10055-建行手机银行，10056-建行龙易行，10057-建行APP，10058-建行自助设备，10059-建行扫码设备
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseRequest) GetTerminalType() string {
    return r._terminalType
}
// TerminalName Setter
// 调用零售药店名称
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseRequest) SetTerminalName(_terminalName string) error {
    r._terminalName = _terminalName
    r.Set("terminal_name", _terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseRequest) GetTerminalName() string {
    return r._terminalName
}
// BureauId Setter
// 门店所属的行政区域ID
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseRequest) SetBureauId(_bureauId string) error {
    r._bureauId = _bureauId
    r.Set("bureau_id", _bureauId)
    return nil
}

// BureauId Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseRequest) GetBureauId() string {
    return r._bureauId
}
// TerminalEntId Setter
// 零售终端id
func (r *AlibabaAlihealthDrugCodeListCodeSuperviseRequest) SetTerminalEntId(_terminalEntId string) error {
    r._terminalEntId = _terminalEntId
    r.Set("terminal_ent_id", _terminalEntId)
    return nil
}

// TerminalEntId Getter
func (r AlibabaAlihealthDrugCodeListCodeSuperviseRequest) GetTerminalEntId() string {
    return r._terminalEntId
}
