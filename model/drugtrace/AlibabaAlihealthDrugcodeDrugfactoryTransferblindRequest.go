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
type AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest struct {
    model.Params
    // 企业id
    _refEntId   string
    // 签名值
    _signValue   string
    // 密文
    _cipherText   string
    // 文件名称
    _fileName   string
}

// 初始化AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest() *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.transferblind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// SignValue Setter
// 签名值
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetSignValue(_signValue string) error {
    r._signValue = _signValue
    r.Set("sign_value", _signValue)
    return nil
}

// SignValue Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetSignValue() string {
    return r._signValue
}
// CipherText Setter
// 密文
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetCipherText(_cipherText string) error {
    r._cipherText = _cipherText
    r.Set("cipher_text", _cipherText)
    return nil
}

// CipherText Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetCipherText() string {
    return r._cipherText
}
// FileName Setter
// 文件名称
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetFileName(_fileName string) error {
    r._fileName = _fileName
    r.Set("file_name", _fileName)
    return nil
}

// FileName Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetFileName() string {
    return r._fileName
}
