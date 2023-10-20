package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdruguploadextinfoAPIRequest 中药饮片及器械对接 API请求
// alibaba.alihealth.drug.upload.extinfo
//
// 中药饮片及器械对接
type AlibabaalihealthdruguploadextinfoAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 药品ID
	_drugId string
	// 批次
	_batchNo string
	// 扩展信息
	_extInfoDto *ExtInfoDto
}

// NewAlibabaalihealthdruguploadextinfoRequest 初始化AlibabaalihealthdruguploadextinfoAPIRequest对象
func NewAlibabaalihealthdruguploadextinfoRequest() *AlibabaalihealthdruguploadextinfoAPIRequest {
	return &AlibabaalihealthdruguploadextinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdruguploadextinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.upload.extinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdruguploadextinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdruguploadextinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdruguploadextinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdruguploadextinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDrugId is DrugId Setter
// 药品ID
func (r *AlibabaalihealthdruguploadextinfoAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// GetDrugId DrugId Getter
func (r AlibabaalihealthdruguploadextinfoAPIRequest) GetDrugId() string {
	return r._drugId
}

// SetBatchNo is BatchNo Setter
// 批次
func (r *AlibabaalihealthdruguploadextinfoAPIRequest) SetBatchNo(_batchNo string) error {
	r._batchNo = _batchNo
	r.Set("batch_no", _batchNo)
	return nil
}

// GetBatchNo BatchNo Getter
func (r AlibabaalihealthdruguploadextinfoAPIRequest) GetBatchNo() string {
	return r._batchNo
}

// SetExtInfoDto is ExtInfoDto Setter
// 扩展信息
func (r *AlibabaalihealthdruguploadextinfoAPIRequest) SetExtInfoDto(_extInfoDto *ExtInfoDto) error {
	r._extInfoDto = _extInfoDto
	r.Set("ext_info_dto", _extInfoDto)
	return nil
}

// GetExtInfoDto ExtInfoDto Getter
func (r AlibabaalihealthdruguploadextinfoAPIRequest) GetExtInfoDto() *ExtInfoDto {
	return r._extInfoDto
}
