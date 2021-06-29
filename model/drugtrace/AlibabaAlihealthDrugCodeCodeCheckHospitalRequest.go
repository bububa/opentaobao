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
    _authRefEntId   string
    // 企业refEntId
    _refEntId   string
    // 城市名
    _bureauName   string
    // 终端名称
    _terminalName   string
    // 终端类型
    _terminalType   string
    // 核销类型
    _cType   string
    // 码列表
    _codes   []string
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
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetAuthRefEntId(_authRefEntId string) error {
    r._authRefEntId = _authRefEntId
    r.Set("auth_ref_ent_id", _authRefEntId)
    return nil
}

// AuthRefEntId Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetAuthRefEntId() string {
    return r._authRefEntId
}
// RefEntId Setter
// 企业refEntId
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetRefEntId() string {
    return r._refEntId
}
// BureauName Setter
// 城市名
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetBureauName(_bureauName string) error {
    r._bureauName = _bureauName
    r.Set("bureau_name", _bureauName)
    return nil
}

// BureauName Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetBureauName() string {
    return r._bureauName
}
// TerminalName Setter
// 终端名称
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetTerminalName(_terminalName string) error {
    r._terminalName = _terminalName
    r.Set("terminal_name", _terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetTerminalName() string {
    return r._terminalName
}
// TerminalType Setter
// 终端类型
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetTerminalType() string {
    return r._terminalType
}
// CType Setter
// 核销类型
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetCType(_cType string) error {
    r._cType = _cType
    r.Set("c_type", _cType)
    return nil
}

// CType Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetCType() string {
    return r._cType
}
// Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetCodes(_codes []string) error {
    r._codes = _codes
    r.Set("codes", _codes)
    return nil
}

// Codes Getter
func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetCodes() []string {
    return r._codes
}
