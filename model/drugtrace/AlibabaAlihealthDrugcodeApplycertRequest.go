package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
申请证书为对接方 API请求
alibaba.alihealth.drugcode.applycert

申请证书 为对接方(当前是药厂和中心化系统)
*/
type AlibabaAlihealthDrugcodeApplycertRequest struct {
    model.Params
    // 设备唯一标识编号
    _serialNum   string
    // 企业Id
    _refEntId   string
    // 企业名称
    _entName   string
    // 证书签名请求
    _csr   string
    // 证书丢失时的操作类型 (true：证书丢失)
    _certLostFlag   bool
}

// 初始化AlibabaAlihealthDrugcodeApplycertRequest对象
func NewAlibabaAlihealthDrugcodeApplycertRequest() *AlibabaAlihealthDrugcodeApplycertRequest{
    return &AlibabaAlihealthDrugcodeApplycertRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeApplycertRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.applycert"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeApplycertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SerialNum Setter
// 设备唯一标识编号
func (r *AlibabaAlihealthDrugcodeApplycertRequest) SetSerialNum(_serialNum string) error {
    r._serialNum = _serialNum
    r.Set("serial_num", _serialNum)
    return nil
}

// SerialNum Getter
func (r AlibabaAlihealthDrugcodeApplycertRequest) GetSerialNum() string {
    return r._serialNum
}
// RefEntId Setter
// 企业Id
func (r *AlibabaAlihealthDrugcodeApplycertRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeApplycertRequest) GetRefEntId() string {
    return r._refEntId
}
// EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugcodeApplycertRequest) SetEntName(_entName string) error {
    r._entName = _entName
    r.Set("ent_name", _entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugcodeApplycertRequest) GetEntName() string {
    return r._entName
}
// Csr Setter
// 证书签名请求
func (r *AlibabaAlihealthDrugcodeApplycertRequest) SetCsr(_csr string) error {
    r._csr = _csr
    r.Set("csr", _csr)
    return nil
}

// Csr Getter
func (r AlibabaAlihealthDrugcodeApplycertRequest) GetCsr() string {
    return r._csr
}
// CertLostFlag Setter
// 证书丢失时的操作类型 (true：证书丢失)
func (r *AlibabaAlihealthDrugcodeApplycertRequest) SetCertLostFlag(_certLostFlag bool) error {
    r._certLostFlag = _certLostFlag
    r.Set("cert_lost_flag", _certLostFlag)
    return nil
}

// CertLostFlag Getter
func (r AlibabaAlihealthDrugcodeApplycertRequest) GetCertLostFlag() bool {
    return r._certLostFlag
}
