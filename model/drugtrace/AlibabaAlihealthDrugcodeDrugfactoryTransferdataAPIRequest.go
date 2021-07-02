package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest 药厂传输数据 API请求
// alibaba.alihealth.drugcode.drugfactory.transferdata
//
// 药厂传输数据
type AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest struct {
	model.Params
	// 时间戳(毫秒级别)
	_timestampYl int64
	// 签名值
	_signValue string
	// 密文
	_cipherText string
	// 企业Id
	_refEntId string
}

// NewAlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest 初始化AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryTransferdataRequest() *AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest {
	return &AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.transferdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TimestampYl Setter
// 时间戳(毫秒级别)
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest) SetTimestampYl(_timestampYl int64) error {
	r._timestampYl = _timestampYl
	r.Set("timestamp_yl", _timestampYl)
	return nil
}

// Get TimestampYl Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest) GetTimestampYl() int64 {
	return r._timestampYl
}

// Set is SignValue Setter
// 签名值
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest) SetSignValue(_signValue string) error {
	r._signValue = _signValue
	r.Set("sign_value", _signValue)
	return nil
}

// Get SignValue Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest) GetSignValue() string {
	return r._signValue
}

// Set is CipherText Setter
// 密文
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest) SetCipherText(_cipherText string) error {
	r._cipherText = _cipherText
	r.Set("cipher_text", _cipherText)
	return nil
}

// Get CipherText Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest) GetCipherText() string {
	return r._cipherText
}

// Set is RefEntId Setter
// 企业Id
func (r *AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIRequest) GetRefEntId() string {
	return r._refEntId
}
