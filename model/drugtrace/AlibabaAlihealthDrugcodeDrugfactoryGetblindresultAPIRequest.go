package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest 获取盲底文件处理结果 API请求
// alibaba.alihealth.drugcode.drugfactory.getblindresult
//
// 获取盲底文件处理结果
type AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 盲底文件名称
	_blindFileName string
}

// NewAlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest 初始化AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest() *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest {
	return &AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.getblindresult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is BlindFileName Setter
// 盲底文件名称
func (r *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) SetBlindFileName(_blindFileName string) error {
	r._blindFileName = _blindFileName
	r.Set("blind_file_name", _blindFileName)
	return nil
}

// Get BlindFileName Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) GetBlindFileName() string {
	return r._blindFileName
}
