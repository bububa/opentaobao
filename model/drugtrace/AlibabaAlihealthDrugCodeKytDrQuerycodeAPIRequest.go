package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytDrQuerycodeAPIRequest 多融根据码查询码信息 API请求
// alibaba.alihealth.drug.code.kyt.dr.querycode
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytDrQuerycodeAPIRequest struct {
	model.Params
	// 企业唯一标识（或appkey）
	_refEntId string
	// 码列表
	_codes []string
}

// NewAlibabaAlihealthDrugCodeKytDrQuerycodeRequest 初始化AlibabaAlihealthDrugCodeKytDrQuerycodeAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytDrQuerycodeRequest() *AlibabaAlihealthDrugCodeKytDrQuerycodeAPIRequest {
	return &AlibabaAlihealthDrugCodeKytDrQuerycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytDrQuerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.dr.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytDrQuerycodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业唯一标识（或appkey）
func (r *AlibabaAlihealthDrugCodeKytDrQuerycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytDrQuerycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeKytDrQuerycodeAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// Get Codes Getter
func (r AlibabaAlihealthDrugCodeKytDrQuerycodeAPIRequest) GetCodes() []string {
	return r._codes
}
