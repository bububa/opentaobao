package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetdruglicenseAPIRequest 获取药品资质信息 API请求
// alibaba.alihealth.drug.kyt.getdruglicense
//
// 获取药品的资质信息。
type AlibabaAlihealthDrugKytGetdruglicenseAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 药品ID
	_drugId string
}

// NewAlibabaAlihealthDrugKytGetdruglicenseRequest 初始化AlibabaAlihealthDrugKytGetdruglicenseAPIRequest对象
func NewAlibabaAlihealthDrugKytGetdruglicenseRequest() *AlibabaAlihealthDrugKytGetdruglicenseAPIRequest {
	return &AlibabaAlihealthDrugKytGetdruglicenseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getdruglicense"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDrugId is DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// GetDrugId DrugId Getter
func (r AlibabaAlihealthDrugKytGetdruglicenseAPIRequest) GetDrugId() string {
	return r._drugId
}
