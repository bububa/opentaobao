package drugtrace

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.smyx.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSmyxGetentinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
