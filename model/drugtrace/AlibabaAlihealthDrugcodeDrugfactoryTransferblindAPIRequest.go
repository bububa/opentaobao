package drugtrace

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) Reset() {
	r._refEntId = ""
	r._signValue = ""
	r._cipherText = ""
	r._fileName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.transferblind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetSignValue is SignValue Setter
// 签名值
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetSignValue(_signValue string) error {
	r._signValue = _signValue
	r.Set("sign_value", _signValue)
	return nil
}

// GetSignValue SignValue Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetSignValue() string {
	return r._signValue
}

// SetCipherText is CipherText Setter
// 密文
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetCipherText(_cipherText string) error {
	r._cipherText = _cipherText
	r.Set("cipher_text", _cipherText)
	return nil
}

// GetCipherText CipherText Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetCipherText() string {
	return r._cipherText
}

// SetFileName is FileName Setter
// 文件名称
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) GetFileName() string {
	return r._fileName
}

var poolAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest()
	},
}

// GetAlibabaAlihealthDrugcodeDrugfactoryTransferblindRequest 从 sync.Pool 获取 AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest
func GetAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest() *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest {
	return poolAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest.Get().(*AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest)
}

// ReleaseAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest 将 AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest(v *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest.Put(v)
}
