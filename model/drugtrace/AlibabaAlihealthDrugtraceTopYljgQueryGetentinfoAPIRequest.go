package drugtrace

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest) Reset() {
	r._entName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.query.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 公司名称(全称)
func (r *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest
func GetAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest() *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest 将 AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest(v *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest.Put(v)
}
