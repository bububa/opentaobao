package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest 根据码获取码信息接口-医保 API请求
// alibaba.alihealth.drug.code.list.code.medical.insurance
//
// 服务描述
// 医保鉴证核查是基于在两定机构的药品管理（入库、出库或盘点）环节，增加扫码匹配
// 与验证鉴核流程；
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息；
// 若所传的监管码是非最小包装监管码，且存在药品混包的情况，则此接口不支持。这种
// 情况下，需要分多次调用该接口。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest struct {
	model.Params
	// 追溯码
	_codeList []string
	// ISV开放平台帐号标识
	_certIsvNo string
	// 调用方式：formal-正式、test-测试
	_invocation string
	// 终端类型 1005100-零售药店 ；10052-医疗机构
	_terminalType string
	// 调用零售药店名称
	_terminalName string
	// 门店所属的行政区域ID
	_bureauId string
	// 零售终端id
	_terminalEntId string
}

// NewAlibabaAlihealthDrugCodeListCodeMedicalInsuranceRequest 初始化AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest对象
func NewAlibabaAlihealthDrugCodeListCodeMedicalInsuranceRequest() *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest {
	return &AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.list.code.medical.insurance"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCodeList is CodeList Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// GetCodeList CodeList Getter
func (r AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) GetCodeList() []string {
	return r._codeList
}

// SetCertIsvNo is CertIsvNo Setter
// ISV开放平台帐号标识
func (r *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) SetCertIsvNo(_certIsvNo string) error {
	r._certIsvNo = _certIsvNo
	r.Set("cert_isv_no", _certIsvNo)
	return nil
}

// GetCertIsvNo CertIsvNo Getter
func (r AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) GetCertIsvNo() string {
	return r._certIsvNo
}

// SetInvocation is Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) SetInvocation(_invocation string) error {
	r._invocation = _invocation
	r.Set("invocation", _invocation)
	return nil
}

// GetInvocation Invocation Getter
func (r AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) GetInvocation() string {
	return r._invocation
}

// SetTerminalType is TerminalType Setter
// 终端类型 1005100-零售药店 ；10052-医疗机构
func (r *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetTerminalName is TerminalName Setter
// 调用零售药店名称
func (r *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// GetTerminalName TerminalName Getter
func (r AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// SetBureauId is BureauId Setter
// 门店所属的行政区域ID
func (r *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) SetBureauId(_bureauId string) error {
	r._bureauId = _bureauId
	r.Set("bureau_id", _bureauId)
	return nil
}

// GetBureauId BureauId Getter
func (r AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) GetBureauId() string {
	return r._bureauId
}

// SetTerminalEntId is TerminalEntId Setter
// 零售终端id
func (r *AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) SetTerminalEntId(_terminalEntId string) error {
	r._terminalEntId = _terminalEntId
	r.Set("terminal_ent_id", _terminalEntId)
	return nil
}

// GetTerminalEntId TerminalEntId Getter
func (r AlibabaAlihealthDrugCodeListCodeMedicalInsuranceAPIRequest) GetTerminalEntId() string {
	return r._terminalEntId
}
