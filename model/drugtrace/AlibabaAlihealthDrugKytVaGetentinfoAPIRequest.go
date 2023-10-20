package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytVaGetentinfoAPIRequest 通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) API请求
// alibaba.alihealth.drug.kyt.va.getentinfo
//
// 根据企业名称查询企业唯一标识（ref_ent_id）及企业ID(ent_id)
type AlibabaAlihealthDrugKytVaGetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// NewAlibabaAlihealthDrugKytVaGetentinfoRequest 初始化AlibabaAlihealthDrugKytVaGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytVaGetentinfoRequest() *AlibabaAlihealthDrugKytVaGetentinfoAPIRequest {
	return &AlibabaAlihealthDrugKytVaGetentinfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytVaGetentinfoAPIRequest) Reset() {
	r._entName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytVaGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.va.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytVaGetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytVaGetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 公司名称(全称)
func (r *AlibabaAlihealthDrugKytVaGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytVaGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}

var poolAlibabaAlihealthDrugKytVaGetentinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytVaGetentinfoRequest()
	},
}

// GetAlibabaAlihealthDrugKytVaGetentinfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytVaGetentinfoAPIRequest
func GetAlibabaAlihealthDrugKytVaGetentinfoAPIRequest() *AlibabaAlihealthDrugKytVaGetentinfoAPIRequest {
	return poolAlibabaAlihealthDrugKytVaGetentinfoAPIRequest.Get().(*AlibabaAlihealthDrugKytVaGetentinfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytVaGetentinfoAPIRequest 将 AlibabaAlihealthDrugKytVaGetentinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytVaGetentinfoAPIRequest(v *AlibabaAlihealthDrugKytVaGetentinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytVaGetentinfoAPIRequest.Put(v)
}
