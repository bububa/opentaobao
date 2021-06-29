package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
药厂传输数据 APIRequest
alibaba.alihealth.drugcode.drugfactory.transferdata

药厂传输数据
*/
type AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest struct {
    model.Params

    // 时间戳(毫秒级别)
    timestampYl   int64 

    // 签名值
    signValue   string 

    // 密文
    cipherText   string 

    // 企业Id
    refEntId   string 

}

func NewAlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest() *AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.transferdata"
}

func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) SetTimestampYl(timestampYl int64) error {
    r.timestampYl = timestampYl
    r.Set("timestamp_yl", timestampYl)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetTimestampYl() int64 {
    return r.timestampYl
}

func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) SetSignValue(signValue string) error {
    r.signValue = signValue
    r.Set("sign_value", signValue)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetSignValue() string {
    return r.signValue
}

func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) SetCipherText(cipherText string) error {
    r.cipherText = cipherText
    r.Set("cipher_text", cipherText)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetCipherText() string {
    return r.cipherText
}

func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetRefEntId() string {
    return r.refEntId
}

