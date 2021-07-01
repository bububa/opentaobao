package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrGetentinfoAPIRequest
通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) API请求
alibaba.alihealth.drug.kyt.dr.getentinfo

根据企业名称查询ID */
type AlibabaAlihealthDrugKytDrGetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// NewAlibabaAlihealthDrugKytDrGetentinfoRequest 初始化AlibabaAlihealthDrugKytDrGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytDrGetentinfoRequest() *AlibabaAlihealthDrugKytDrGetentinfoAPIRequest {
	return &AlibabaAlihealthDrugKytDrGetentinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrGetentinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EntName Setter
// 公司名称(全称)
func (r *AlibabaAlihealthDrugKytDrGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// Get EntName Getter
func (r AlibabaAlihealthDrugKytDrGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
