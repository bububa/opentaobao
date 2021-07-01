package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码核查状态同步-医保 API请求
alibaba.alihealth.drug.code.code.check.medical.insurance

服务描述
核查平台在进行医保单据鉴证核查时，会记录单据中所有提交的追溯码信息；单据中的
追溯码包含所有正常和异常的数据；
此接口，针对正式鉴核的单据中提交的有效的、正常状态的追溯码，提供可由核查平台
发起，按单据鉴核时间顺序组织，向码上放心平台同步每笔单据核销的码状态信息；
入参采用数组方式提供，一次同步最多支持100条记录
*/
type AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest struct {
    model.Params
    // 行政区域
    _bureauName   string
    // 终端id
    _terminalId   string
    // 终端类型（1005100-零售药店；1005200-医疗机构）
    _terminalType   string
    // 核销类型(1012100：核销；1012900：退库)
    _cType   string
    // 码列表
    _codes   []string
    // 平台返回的终端id
    _terminalEntId   string
}

// 初始化AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest对象
func NewAlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest() *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest{
    return &AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.code.check.medical.insurance"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BureauName Setter
// 行政区域
func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) SetBureauName(_bureauName string) error {
    r._bureauName = _bureauName
    r.Set("bureau_name", _bureauName)
    return nil
}

// BureauName Getter
func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) GetBureauName() string {
    return r._bureauName
}
// TerminalId Setter
// 终端id
func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) SetTerminalId(_terminalId string) error {
    r._terminalId = _terminalId
    r.Set("terminal_id", _terminalId)
    return nil
}

// TerminalId Getter
func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) GetTerminalId() string {
    return r._terminalId
}
// TerminalType Setter
// 终端类型（1005100-零售药店；1005200-医疗机构）
func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) GetTerminalType() string {
    return r._terminalType
}
// CType Setter
// 核销类型(1012100：核销；1012900：退库)
func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) SetCType(_cType string) error {
    r._cType = _cType
    r.Set("c_type", _cType)
    return nil
}

// CType Getter
func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) GetCType() string {
    return r._cType
}
// Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) SetCodes(_codes []string) error {
    r._codes = _codes
    r.Set("codes", _codes)
    return nil
}

// Codes Getter
func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) GetCodes() []string {
    return r._codes
}
// TerminalEntId Setter
// 平台返回的终端id
func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) SetTerminalEntId(_terminalEntId string) error {
    r._terminalEntId = _terminalEntId
    r.Set("terminal_ent_id", _terminalEntId)
    return nil
}

// TerminalEntId Getter
func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceAPIRequest) GetTerminalEntId() string {
    return r._terminalEntId
}
