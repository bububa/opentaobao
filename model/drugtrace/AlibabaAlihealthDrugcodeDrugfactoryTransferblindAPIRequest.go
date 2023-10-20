package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest 传输盲底文件 API请求
// alibaba.alihealth.drugcode.drugfactory.transferblind
//
// 临床用药试验-传输盲底文件
type AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest struct {
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

// NewAlibabaalihealthdrugcodedrugfactorytransferblindRequest 初始化AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest对象
func NewAlibabaalihealthdrugcodedrugfactorytransferblindRequest() *AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest {
	return &AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.transferblind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetSignValue is SignValue Setter
// 签名值
func (r *AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest) SetSignValue(_signValue string) error {
	r._signValue = _signValue
	r.Set("sign_value", _signValue)
	return nil
}

// GetSignValue SignValue Getter
func (r AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest) GetSignValue() string {
	return r._signValue
}

// SetCipherText is CipherText Setter
// 密文
func (r *AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest) SetCipherText(_cipherText string) error {
	r._cipherText = _cipherText
	r.Set("cipher_text", _cipherText)
	return nil
}

// GetCipherText CipherText Getter
func (r AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest) GetCipherText() string {
	return r._cipherText
}

// SetFileName is FileName Setter
// 文件名称
func (r *AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest) GetFileName() string {
	return r._fileName
}
