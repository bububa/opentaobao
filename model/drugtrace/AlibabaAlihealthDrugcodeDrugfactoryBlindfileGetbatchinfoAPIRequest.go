package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest 获取盲底文件中的批次信息 API请求
// alibaba.alihealth.drugcode.drugfactory.blindfile.getbatchinfo
//
// 获取盲底文件中的批次信息
type AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest struct {
	model.Params
	// 企业Id
	_refEntId string
	// 药品子类编码
	_subTypeNo string
	// 盲底文件上传开始时间（yyyy-MM-dd）
	_blindFileStartDate string
	// 盲底文件上传结束时间（yyyy-MM-dd）
	_blindFileEndDate string
}

// NewAlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoRequest 初始化AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoRequest() *AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest {
	return &AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.blindfile.getbatchinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefEntId is RefEntId Setter
// 企业Id
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetSubTypeNo is SubTypeNo Setter
// 药品子类编码
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest) SetSubTypeNo(_subTypeNo string) error {
	r._subTypeNo = _subTypeNo
	r.Set("sub_type_no", _subTypeNo)
	return nil
}

// GetSubTypeNo SubTypeNo Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest) GetSubTypeNo() string {
	return r._subTypeNo
}

// SetBlindFileStartDate is BlindFileStartDate Setter
// 盲底文件上传开始时间（yyyy-MM-dd）
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest) SetBlindFileStartDate(_blindFileStartDate string) error {
	r._blindFileStartDate = _blindFileStartDate
	r.Set("blind_file_start_date", _blindFileStartDate)
	return nil
}

// GetBlindFileStartDate BlindFileStartDate Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest) GetBlindFileStartDate() string {
	return r._blindFileStartDate
}

// SetBlindFileEndDate is BlindFileEndDate Setter
// 盲底文件上传结束时间（yyyy-MM-dd）
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest) SetBlindFileEndDate(_blindFileEndDate string) error {
	r._blindFileEndDate = _blindFileEndDate
	r.Set("blind_file_end_date", _blindFileEndDate)
	return nil
}

// GetBlindFileEndDate BlindFileEndDate Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest) GetBlindFileEndDate() string {
	return r._blindFileEndDate
}
