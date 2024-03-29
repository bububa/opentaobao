package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest 通过企业名得到唯一标识ref_ent_id及企业ent_id API请求
// alibaba.alihealth.drugtrace.top.lsyd.query.getentinfo
//
// 根据企业名称查询ID
type AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// NewAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest 初始化AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest() *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest {
	return &AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest) Reset() {
	r._entName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.query.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 公司名称(全称)
func (r *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest
func GetAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest() *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest 将 AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest(v *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIRequest.Put(v)
}
