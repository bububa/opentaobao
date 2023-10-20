package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdrugdetailAPIRequest 查询药品详细信息 API请求
// alibaba.alihealth.drug.kyt.drugdetail
//
// 查询药品详细信息
type AlibabaalihealthdrugkytdrugdetailAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 药品ID
	_drugEntBaseInfoId string
}

// NewAlibabaalihealthdrugkytdrugdetailRequest 初始化AlibabaalihealthdrugkytdrugdetailAPIRequest对象
func NewAlibabaalihealthdrugkytdrugdetailRequest() *AlibabaalihealthdrugkytdrugdetailAPIRequest {
	return &AlibabaalihealthdrugkytdrugdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytdrugdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.drugdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytdrugdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytdrugdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytdrugdetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytdrugdetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDrugEntBaseInfoId is DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaalihealthdrugkytdrugdetailAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
	r._drugEntBaseInfoId = _drugEntBaseInfoId
	r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
	return nil
}

// GetDrugEntBaseInfoId DrugEntBaseInfoId Getter
func (r AlibabaalihealthdrugkytdrugdetailAPIRequest) GetDrugEntBaseInfoId() string {
	return r._drugEntBaseInfoId
}
