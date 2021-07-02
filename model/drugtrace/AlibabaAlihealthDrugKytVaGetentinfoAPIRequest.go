package drugtrace

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytVaGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.va.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytVaGetentinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EntName Setter
// 公司名称(全称)
func (r *AlibabaAlihealthDrugKytVaGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// Get EntName Getter
func (r AlibabaAlihealthDrugKytVaGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
