package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytYyGetentinfoAPIRequest 得到企业信息 API请求
// alibaba.alihealth.drug.kyt.yy.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaAlihealthDrugKytYyGetentinfoAPIRequest struct {
	model.Params
	// 公司名称
	_entName string
}

// NewAlibabaAlihealthDrugKytYyGetentinfoRequest 初始化AlibabaAlihealthDrugKytYyGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytYyGetentinfoRequest() *AlibabaAlihealthDrugKytYyGetentinfoAPIRequest {
	return &AlibabaAlihealthDrugKytYyGetentinfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytYyGetentinfoAPIRequest) Reset() {
	r._entName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yy.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyGetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytYyGetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 公司名称
func (r *AlibabaAlihealthDrugKytYyGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytYyGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}

var poolAlibabaAlihealthDrugKytYyGetentinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytYyGetentinfoRequest()
	},
}

// GetAlibabaAlihealthDrugKytYyGetentinfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytYyGetentinfoAPIRequest
func GetAlibabaAlihealthDrugKytYyGetentinfoAPIRequest() *AlibabaAlihealthDrugKytYyGetentinfoAPIRequest {
	return poolAlibabaAlihealthDrugKytYyGetentinfoAPIRequest.Get().(*AlibabaAlihealthDrugKytYyGetentinfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytYyGetentinfoAPIRequest 将 AlibabaAlihealthDrugKytYyGetentinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytYyGetentinfoAPIRequest(v *AlibabaAlihealthDrugKytYyGetentinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytYyGetentinfoAPIRequest.Put(v)
}
