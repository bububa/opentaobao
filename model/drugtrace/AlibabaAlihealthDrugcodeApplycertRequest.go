package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
申请证书为对接方 APIRequest
alibaba.alihealth.drugcode.applycert

申请证书 为对接方(当前是药厂和中心化系统)
*/
type AlibabaAlihealthDrugcodeApplycertRequest struct {
    model.Params

    // 设备唯一标识编号
    serialNum   string 

    // 企业Id
    refEntId   string 

    // 企业名称
    entName   string 

    // 证书签名请求
    csr   string 

    // 证书丢失时的操作类型 (true：证书丢失)
    certLostFlag   bool 

}

func NewAlibabaAlihealthDrugcodeApplycertRequest() *AlibabaAlihealthDrugcodeApplycertRequest{
    return &AlibabaAlihealthDrugcodeApplycertRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugcodeApplycertRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.applycert"
}

func (r AlibabaAlihealthDrugcodeApplycertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugcodeApplycertRequest) SetSerialNum(serialNum string) error {
    r.serialNum = serialNum
    r.Set("serial_num", serialNum)
    return nil
}

func (r AlibabaAlihealthDrugcodeApplycertRequest) GetSerialNum() string {
    return r.serialNum
}

func (r *AlibabaAlihealthDrugcodeApplycertRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugcodeApplycertRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugcodeApplycertRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugcodeApplycertRequest) GetEntName() string {
    return r.entName
}

func (r *AlibabaAlihealthDrugcodeApplycertRequest) SetCsr(csr string) error {
    r.csr = csr
    r.Set("csr", csr)
    return nil
}

func (r AlibabaAlihealthDrugcodeApplycertRequest) GetCsr() string {
    return r.csr
}

func (r *AlibabaAlihealthDrugcodeApplycertRequest) SetCertLostFlag(certLostFlag bool) error {
    r.certLostFlag = certLostFlag
    r.Set("cert_lost_flag", certLostFlag)
    return nil
}

func (r AlibabaAlihealthDrugcodeApplycertRequest) GetCertLostFlag() bool {
    return r.certLostFlag
}

