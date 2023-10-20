package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrGetentinfoAPIRequest 通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) API请求
// alibaba.alihealth.drug.kyt.dr.getentinfo
//
// 根据企业名称查询ID
type AlibabaAlihealthDrugKytDrGetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// NewAlibabaAlihealthDrugKytDrGetentinfoRequest 初始化AlibabaAlihealthDrugKytDrGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytDrGetentinfoRequest() *AlibabaAlihealthDrugKytDrGetentinfoAPIRequest {
	return &AlibabaAlihealthDrugKytDrGetentinfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytDrGetentinfoAPIRequest) Reset() {
	r._entName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrGetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytDrGetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 公司名称(全称)
func (r *AlibabaAlihealthDrugKytDrGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytDrGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}

var poolAlibabaAlihealthDrugKytDrGetentinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytDrGetentinfoRequest()
	},
}

// GetAlibabaAlihealthDrugKytDrGetentinfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrGetentinfoAPIRequest
func GetAlibabaAlihealthDrugKytDrGetentinfoAPIRequest() *AlibabaAlihealthDrugKytDrGetentinfoAPIRequest {
	return poolAlibabaAlihealthDrugKytDrGetentinfoAPIRequest.Get().(*AlibabaAlihealthDrugKytDrGetentinfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytDrGetentinfoAPIRequest 将 AlibabaAlihealthDrugKytDrGetentinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrGetentinfoAPIRequest(v *AlibabaAlihealthDrugKytDrGetentinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrGetentinfoAPIRequest.Put(v)
}
