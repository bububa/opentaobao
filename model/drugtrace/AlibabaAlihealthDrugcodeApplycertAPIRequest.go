package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeApplycertAPIRequest 申请证书为对接方 API请求
// alibaba.alihealth.drugcode.applycert
//
// 申请证书 为对接方(当前是药厂和中心化系统)
type AlibabaAlihealthDrugcodeApplycertAPIRequest struct {
	model.Params
	// 设备唯一标识编号
	_serialNum string
	// 企业Id
	_refEntId string
	// 企业名称
	_entName string
	// 证书签名请求
	_csr string
	// 证书丢失时的操作类型 (true：证书丢失)
	_certLostFlag bool
}

// NewAlibabaAlihealthDrugcodeApplycertRequest 初始化AlibabaAlihealthDrugcodeApplycertAPIRequest对象
func NewAlibabaAlihealthDrugcodeApplycertRequest() *AlibabaAlihealthDrugcodeApplycertAPIRequest {
	return &AlibabaAlihealthDrugcodeApplycertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.applycert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSerialNum is SerialNum Setter
// 设备唯一标识编号
func (r *AlibabaAlihealthDrugcodeApplycertAPIRequest) SetSerialNum(_serialNum string) error {
	r._serialNum = _serialNum
	r.Set("serial_num", _serialNum)
	return nil
}

// GetSerialNum SerialNum Getter
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetSerialNum() string {
	return r._serialNum
}

// SetRefEntId is RefEntId Setter
// 企业Id
func (r *AlibabaAlihealthDrugcodeApplycertAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugcodeApplycertAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetEntName() string {
	return r._entName
}

// SetCsr is Csr Setter
// 证书签名请求
func (r *AlibabaAlihealthDrugcodeApplycertAPIRequest) SetCsr(_csr string) error {
	r._csr = _csr
	r.Set("csr", _csr)
	return nil
}

// GetCsr Csr Getter
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetCsr() string {
	return r._csr
}

// SetCertLostFlag is CertLostFlag Setter
// 证书丢失时的操作类型 (true：证书丢失)
func (r *AlibabaAlihealthDrugcodeApplycertAPIRequest) SetCertLostFlag(_certLostFlag bool) error {
	r._certLostFlag = _certLostFlag
	r.Set("cert_lost_flag", _certLostFlag)
	return nil
}

// GetCertLostFlag CertLostFlag Getter
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetCertLostFlag() bool {
	return r._certLostFlag
}
