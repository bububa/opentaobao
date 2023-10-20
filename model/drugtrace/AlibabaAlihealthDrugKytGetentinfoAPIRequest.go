package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetentinfoAPIRequest 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 API请求
// alibaba.alihealth.drug.kyt.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaAlihealthDrugKytGetentinfoAPIRequest struct {
	model.Params
	// 企业名称
	_entName string
}

// NewAlibabaAlihealthDrugKytGetentinfoRequest 初始化AlibabaAlihealthDrugKytGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytGetentinfoRequest() *AlibabaAlihealthDrugKytGetentinfoAPIRequest {
	return &AlibabaAlihealthDrugKytGetentinfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytGetentinfoAPIRequest) Reset() {
	r._entName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytGetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}

var poolAlibabaAlihealthDrugKytGetentinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytGetentinfoRequest()
	},
}

// GetAlibabaAlihealthDrugKytGetentinfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytGetentinfoAPIRequest
func GetAlibabaAlihealthDrugKytGetentinfoAPIRequest() *AlibabaAlihealthDrugKytGetentinfoAPIRequest {
	return poolAlibabaAlihealthDrugKytGetentinfoAPIRequest.Get().(*AlibabaAlihealthDrugKytGetentinfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytGetentinfoAPIRequest 将 AlibabaAlihealthDrugKytGetentinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytGetentinfoAPIRequest(v *AlibabaAlihealthDrugKytGetentinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytGetentinfoAPIRequest.Put(v)
}
