package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetentinfoAPIRequest 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 API请求
// alibaba.alihealth.drug.kyt.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaAlihealthDrugKytGetentinfoAPIRequest struct {
	model.Params
	// 公司名称
	_entName string
}

// NewAlibabaAlihealthDrugKytGetentinfoRequest 初始化AlibabaAlihealthDrugKytGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytGetentinfoRequest() *AlibabaAlihealthDrugKytGetentinfoAPIRequest {
	return &AlibabaAlihealthDrugKytGetentinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetentinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEntName is EntName Setter
// 公司名称
func (r *AlibabaAlihealthDrugKytGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
