package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
传输盲底文件 API请求
alibaba.alihealth.drugcode.drugfactory.transferblind

临床用药试验-传输盲底文件
*/
type AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest struct {
    model.Params
    // 企业id
    refEntId   string
    // 签名值
    signValue   string
    // 密文
    cipherText   string
    // 文件名称
    fileName   string
}

// 初始化AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest() *AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.transferblind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest) GetRefEntId() string {
    return r.refEntId
}
// SignValue Setter
// 签名值
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest) SetSignValue(signValue string) error {
    r.signValue = signValue
    r.Set("sign_value", signValue)
    return nil
}

// SignValue Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest) GetSignValue() string {
    return r.signValue
}
// CipherText Setter
// 密文
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest) SetCipherText(cipherText string) error {
    r.cipherText = cipherText
    r.Set("cipher_text", cipherText)
    return nil
}

// CipherText Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest) GetCipherText() string {
    return r.cipherText
}
// FileName Setter
// 文件名称
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

// FileName Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest) GetFileName() string {
    return r.fileName
}
