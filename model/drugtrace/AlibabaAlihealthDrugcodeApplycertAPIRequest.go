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
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SerialNum Setter
// 设备唯一标识编号
func (r *AlibabaAlihealthDrugcodeApplycertAPIRequest) SetSerialNum(_serialNum string) error {
	r._serialNum = _serialNum
	r.Set("serial_num", _serialNum)
	return nil
}

// Get SerialNum Getter
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetSerialNum() string {
	return r._serialNum
}

// Set is RefEntId Setter
// 企业Id
func (r *AlibabaAlihealthDrugcodeApplycertAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugcodeApplycertAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// Get EntName Getter
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetEntName() string {
	return r._entName
}

// Set is Csr Setter
// 证书签名请求
func (r *AlibabaAlihealthDrugcodeApplycertAPIRequest) SetCsr(_csr string) error {
	r._csr = _csr
	r.Set("csr", _csr)
	return nil
}

// Get Csr Getter
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetCsr() string {
	return r._csr
}

// Set is CertLostFlag Setter
// 证书丢失时的操作类型 (true：证书丢失)
func (r *AlibabaAlihealthDrugcodeApplycertAPIRequest) SetCertLostFlag(_certLostFlag bool) error {
	r._certLostFlag = _certLostFlag
	r.Set("cert_lost_flag", _certLostFlag)
	return nil
}

// Get CertLostFlag Getter
func (r AlibabaAlihealthDrugcodeApplycertAPIRequest) GetCertLostFlag() bool {
	return r._certLostFlag
}
