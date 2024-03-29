package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesGetentinfoAPIRequest 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 API请求
// alibaba.alihealth.drug.kyt.wes.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaAlihealthDrugKytWesGetentinfoAPIRequest struct {
	model.Params
	// 企业名称
	_entName string
}

// NewAlibabaAlihealthDrugKytWesGetentinfoRequest 初始化AlibabaAlihealthDrugKytWesGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytWesGetentinfoRequest() *AlibabaAlihealthDrugKytWesGetentinfoAPIRequest {
	return &AlibabaAlihealthDrugKytWesGetentinfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytWesGetentinfoAPIRequest) Reset() {
	r._entName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesGetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesGetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytWesGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytWesGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}

var poolAlibabaAlihealthDrugKytWesGetentinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytWesGetentinfoRequest()
	},
}

// GetAlibabaAlihealthDrugKytWesGetentinfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesGetentinfoAPIRequest
func GetAlibabaAlihealthDrugKytWesGetentinfoAPIRequest() *AlibabaAlihealthDrugKytWesGetentinfoAPIRequest {
	return poolAlibabaAlihealthDrugKytWesGetentinfoAPIRequest.Get().(*AlibabaAlihealthDrugKytWesGetentinfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytWesGetentinfoAPIRequest 将 AlibabaAlihealthDrugKytWesGetentinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesGetentinfoAPIRequest(v *AlibabaAlihealthDrugKytWesGetentinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesGetentinfoAPIRequest.Put(v)
}
