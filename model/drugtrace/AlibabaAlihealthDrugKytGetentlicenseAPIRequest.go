package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetentlicenseAPIRequest 获取企业资质 API请求
// alibaba.alihealth.drug.kyt.getentlicense
//
// 获取企业的资质信息。
type AlibabaAlihealthDrugKytGetentlicenseAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
}

// NewAlibabaAlihealthDrugKytGetentlicenseRequest 初始化AlibabaAlihealthDrugKytGetentlicenseAPIRequest对象
func NewAlibabaAlihealthDrugKytGetentlicenseRequest() *AlibabaAlihealthDrugKytGetentlicenseAPIRequest {
	return &AlibabaAlihealthDrugKytGetentlicenseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetentlicenseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getentlicense"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetentlicenseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytGetentlicenseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytGetentlicenseAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytGetentlicenseAPIRequest) GetRefEntId() string {
	return r._refEntId
}
