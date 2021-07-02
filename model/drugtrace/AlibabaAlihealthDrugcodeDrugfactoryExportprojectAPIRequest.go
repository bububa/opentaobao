package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest 导出项目和药品目录 API请求
// alibaba.alihealth.drugcode.drugfactory.exportproject
//
// 导出临床项目及其药品目录
type AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
}

// NewAlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest 初始化AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest {
	return &AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.exportproject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest) GetRefEntId() string {
	return r._refEntId
}
