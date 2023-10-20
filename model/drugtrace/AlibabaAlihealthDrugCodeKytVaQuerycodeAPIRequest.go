package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest 根据码查询码信息 API请求
// alibaba.alihealth.drug.code.kyt.va.querycode
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest struct {
	model.Params
	// 码列表
	_codes []string
	// 企业唯一标识（或appkey）
	_refEntId string
}

// NewAlibabaAlihealthDrugCodeKytVaQuerycodeRequest 初始化AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytVaQuerycodeRequest() *AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest {
	return &AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.va.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识（或appkey）
func (r *AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytVaQuerycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
