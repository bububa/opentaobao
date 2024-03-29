package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest 查企业标识信息 API请求
// alibaba.alihealth.drug.kyt.smyx.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest struct {
	model.Params
	// 公司名称
	_entName string
}

// NewAlibabaAlihealthDrugKytSmyxGetentinfoRequest 初始化AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytSmyxGetentinfoRequest() *AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest {
	return &AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest) Reset() {
	r._entName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.smyx.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntName is EntName Setter
// 公司名称
func (r *AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}

var poolAlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytSmyxGetentinfoRequest()
	},
}

// GetAlibabaAlihealthDrugKytSmyxGetentinfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest
func GetAlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest() *AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest {
	return poolAlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest.Get().(*AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest 将 AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest(v *AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest.Put(v)
}
