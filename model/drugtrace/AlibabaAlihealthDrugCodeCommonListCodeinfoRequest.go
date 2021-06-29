package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通用查询码接口 APIRequest
alibaba.alihealth.drug.code.common.list.codeinfo

通用查询码接口
*/
type AlibabaAlihealthDrugCodeCommonListCodeinfoRequest struct {
    model.Params

    // 企业refEntId
    refEntId   string 

    // 标示医院业务
    searchSource   string 

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

    // 验证权限企业id
    authRefEntId   string 

}

func NewAlibabaAlihealthDrugCodeCommonListCodeinfoRequest() *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest{
    return &AlibabaAlihealthDrugCodeCommonListCodeinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.common.list.codeinfo"
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetSearchSource(searchSource string) error {
    r.searchSource = searchSource
    r.Set("search_source", searchSource)
    return nil
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetSearchSource() string {
    return r.searchSource
}

func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetCodeList(codeList []string) error {
    r.codeList = codeList
    r.Set("code_list", codeList)
    return nil
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetCodeList() []string {
    return r.codeList
}

func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetCertIsvNo(certIsvNo string) error {
    r.certIsvNo = certIsvNo
    r.Set("cert_isv_no", certIsvNo)
    return nil
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetCertIsvNo() string {
    return r.certIsvNo
}

func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetInvocation(invocation string) error {
    r.invocation = invocation
    r.Set("invocation", invocation)
    return nil
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetInvocation() string {
    return r.invocation
}

func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetTerminalName(terminalName string) error {
    r.terminalName = terminalName
    r.Set("terminal_name", terminalName)
    return nil
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetTerminalName() string {
    return r.terminalName
}

func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetBureauName(bureauName string) error {
    r.bureauName = bureauName
    r.Set("bureau_name", bureauName)
    return nil
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetBureauName() string {
    return r.bureauName
}

func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetErrorMessage(errorMessage string) error {
    r.errorMessage = errorMessage
    r.Set("error_message", errorMessage)
    return nil
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetErrorMessage() string {
    return r.errorMessage
}

func (r *AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) SetAuthRefEntId(authRefEntId string) error {
    r.authRefEntId = authRefEntId
    r.Set("auth_ref_ent_id", authRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugCodeCommonListCodeinfoRequest) GetAuthRefEntId() string {
    return r.authRefEntId
}

