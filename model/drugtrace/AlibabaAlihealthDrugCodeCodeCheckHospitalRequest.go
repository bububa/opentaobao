package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码核查状态同步-医院 APIRequest
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

func NewAlibabaAlihealthDrugCodeCodeCheckHospitalRequest() *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest{
    return &AlibabaAlihealthDrugCodeCodeCheckHospitalRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.code.check.hospital"
}

func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetAuthRefEntId(authRefEntId string) error {
    r.authRefEntId = authRefEntId
    r.Set("auth_ref_ent_id", authRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetAuthRefEntId() string {
    return r.authRefEntId
}

func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetBureauName(bureauName string) error {
    r.bureauName = bureauName
    r.Set("bureau_name", bureauName)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetBureauName() string {
    return r.bureauName
}

func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetTerminalName(terminalName string) error {
    r.terminalName = terminalName
    r.Set("terminal_name", terminalName)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetTerminalName() string {
    return r.terminalName
}

func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetCType(cType string) error {
    r.cType = cType
    r.Set("c_type", cType)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetCType() string {
    return r.cType
}

func (r *AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckHospitalRequest) GetCodes() []string {
    return r.codes
}

