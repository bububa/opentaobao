package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest 传输盲底文件 API请求
// alibaba.alihealth.drugcode.drugfactory.transferblind
//
// 临床用药试验-传输盲底文件
type AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 签名值
	_signValue string
	// 密文
	_cipherText string
	// 文件名称
	_fileName string
}

// NewAlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest 初始化AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest() *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest {
	return &AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.transferblind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is SignValue Setter
// 签名值
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetSignValue(_signValue string) error {
	r._signValue = _signValue
	r.Set("sign_value", _signValue)
	return nil
}

// Get SignValue Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetSignValue() string {
	return r._signValue
}

// Set is CipherText Setter
// 密文
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetCipherText(_cipherText string) error {
	r._cipherText = _cipherText
	r.Set("cipher_text", _cipherText)
	return nil
}

// Get CipherText Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetCipherText() string {
	return r._cipherText
}

// Set is FileName Setter
// 文件名称
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// Get FileName Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetFileName() string {
	return r._fileName
}
