package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码核查状态同步-医保 APIRequest
alibaba.alihealth.drug.code.code.check.medical.insurance

服务描述
核查平台在进行医保单据鉴证核查时，会记录单据中所有提交的追溯码信息；单据中的
追溯码包含所有正常和异常的数据；
此接口，针对正式鉴核的单据中提交的有效的、正常状态的追溯码，提供可由核查平台
发起，按单据鉴核时间顺序组织，向码上放心平台同步每笔单据核销的码状态信息；
入参采用数组方式提供，一次同步最多支持100条记录
*/
type AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest struct {
    model.Params

    // 行政区域
    bureauName   string 

    // 终端id
    terminalId   string 

    // 终端类型（1005100-零售药店；1005200-医疗机构）
    terminalType   string 

    // 核销类型(1012100：核销；1012900：退库)
    cType   string 

    // 码列表
    codes   []string 

    // 平台返回的终端id
    terminalEntId   string 

}

func NewAlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest() *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest{
    return &AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.code.check.medical.insurance"
}

func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) SetBureauName(bureauName string) error {
    r.bureauName = bureauName
    r.Set("bureau_name", bureauName)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) GetBureauName() string {
    return r.bureauName
}

func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) SetTerminalId(terminalId string) error {
    r.terminalId = terminalId
    r.Set("terminal_id", terminalId)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) GetTerminalId() string {
    return r.terminalId
}

func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) SetCType(cType string) error {
    r.cType = cType
    r.Set("c_type", cType)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) GetCType() string {
    return r.cType
}

func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) GetCodes() []string {
    return r.codes
}

func (r *AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) SetTerminalEntId(terminalEntId string) error {
    r.terminalEntId = terminalEntId
    r.Set("terminal_ent_id", terminalEntId)
    return nil
}

func (r AlibabaAlihealthDrugCodeCodeCheckMedicalInsuranceRequest) GetTerminalEntId() string {
    return r.terminalEntId
}

