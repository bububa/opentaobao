package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest 通过企业名得到唯一标识ref_ent_id及企业ent_id API请求
// alibaba.alihealth.drugtrace.top.yljg.query.getentinfo
//
// 根据企业名称查询ID
type AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// NewAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest 初始化AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest() *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest {
	return &AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.query.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EntName Setter
// 公司名称(全称)
func (r *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// Get EntName Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
