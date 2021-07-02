package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest 根据码查询码信息 API请求
// alibaba.alihealth.drugtrace.top.yljg.query.codedetail
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest struct {
	model.Params
	// 企业唯一标识（或appkey）
	_refEntId string
	// 码列表
	_codes []string
}

// NewAlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest 初始化AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest() *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest {
	return &AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.query.codedetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识（或appkey）
func (r *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCodes is Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) GetCodes() []string {
	return r._codes
}
