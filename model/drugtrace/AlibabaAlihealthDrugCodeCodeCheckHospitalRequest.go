package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码核查状态同步-医院 API请求
alibaba.alihealth.drug.code.code.check.hospital

码核查状态同步-医院
*/
type AlibabaAlihealthDrugCodeCodeCheckHospitalRequest struct {
    model.Params
    // 认证企业refEntId
    authRefEntId   string
    // 企业refEntId
    refEntId   string
    // 城市名
    bureauName   string
    // 终端名称
    terminalName   string
    // 终端类型
    terminalType   string
    // 核销类型
    cType   string
    // 码列表
    codes   []string
}

// 初始化AlibabaAlihealthDrugCodeCodeCheckHospitalRequest对象
func NewAlibabaAlihealthDrugCodeCodeCheckHospitalRequest() *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest{
    return &AlibabaAlihealthDrugCodeCodeCheckHospitalRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.code.check.hospital"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AuthRefEntId Setter
// 认证企业refEntId
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetAuthRefEntId(authRefEntId string) error {
    r.authRefEntId = authRefEntId
    r.Set("auth_ref_ent_id", authRefEntId)
    return nil
}

// AuthRefEntId Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetAuthRefEntId() string {
    return r.authRefEntId
}
// RefEntId Setter
// 企业refEntId
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetRefEntId() string {
    return r.refEntId
}
// BureauName Setter
// 城市名
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetBureauName(bureauName string) error {
    r.bureauName = bureauName
    r.Set("bureau_name", bureauName)
    return nil
}

// BureauName Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetBureauName() string {
    return r.bureauName
}
// TerminalName Setter
// 终端名称
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetTerminalName(terminalName string) error {
    r.terminalName = terminalName
    r.Set("terminal_name", terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetTerminalName() string {
    return r.terminalName
}
// TerminalType Setter
// 终端类型
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetTerminalType() string {
    return r.terminalType
}
// CType Setter
// 核销类型
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetCType(cType string) error {
    r.cType = cType
    r.Set("c_type", cType)
    return nil
}

// CType Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetCType() string {
    return r.cType
}
// Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

// Codes Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetCodes() []string {
    return r.codes
}
