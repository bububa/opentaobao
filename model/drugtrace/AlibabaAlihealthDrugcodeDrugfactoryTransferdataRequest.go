package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
药厂传输数据 API请求
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

// 初始化AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest() *AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.transferdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TimestampYl Setter
// 时间戳(毫秒级别)
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) SetTimestampYl(timestampYl int64) error {
    r.timestampYl = timestampYl
    r.Set("timestamp_yl", timestampYl)
    return nil
}

// TimestampYl Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetTimestampYl() int64 {
    return r.timestampYl
}
// SignValue Setter
// 签名值
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) SetSignValue(signValue string) error {
    r.signValue = signValue
    r.Set("sign_value", signValue)
    return nil
}

// SignValue Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetSignValue() string {
    return r.signValue
}
// CipherText Setter
// 密文
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) SetCipherText(cipherText string) error {
    r.cipherText = cipherText
    r.Set("cipher_text", cipherText)
    return nil
}

// CipherText Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetCipherText() string {
    return r.cipherText
}
// RefEntId Setter
// 企业Id
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest) GetRefEntId() string {
    return r.refEntId
}
