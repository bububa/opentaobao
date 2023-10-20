package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdrgetproteminfoAPIRequest 疫苗，获取生产企业的存储和运输温度 API请求
// alibaba.alihealth.drug.kyt.dr.getproteminfo
//
// 疫苗，获取生产企业的存储和运输温度
type AlibabaalihealthdrugkytdrgetproteminfoAPIRequest struct {
	model.Params
	// 调用企业ID
	_refEntId string
	// 药品ID
	_drugEntBaseInfoId string
	// 批次编号
	_batchNo string
	// 出库单号
	_billOutCode string
}

// NewAlibabaalihealthdrugkytdrgetproteminfoRequest 初始化AlibabaalihealthdrugkytdrgetproteminfoAPIRequest对象
func NewAlibabaalihealthdrugkytdrgetproteminfoRequest() *AlibabaalihealthdrugkytdrgetproteminfoAPIRequest {
	return &AlibabaalihealthdrugkytdrgetproteminfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytdrgetproteminfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.getproteminfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytdrgetproteminfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytdrgetproteminfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用企业ID
func (r *AlibabaalihealthdrugkytdrgetproteminfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytdrgetproteminfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDrugEntBaseInfoId is DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaalihealthdrugkytdrgetproteminfoAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
	r._drugEntBaseInfoId = _drugEntBaseInfoId
	r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
	return nil
}

// GetDrugEntBaseInfoId DrugEntBaseInfoId Getter
func (r AlibabaalihealthdrugkytdrgetproteminfoAPIRequest) GetDrugEntBaseInfoId() string {
	return r._drugEntBaseInfoId
}

// SetBatchNo is BatchNo Setter
// 批次编号
func (r *AlibabaalihealthdrugkytdrgetproteminfoAPIRequest) SetBatchNo(_batchNo string) error {
	r._batchNo = _batchNo
	r.Set("batch_no", _batchNo)
	return nil
}

// GetBatchNo BatchNo Getter
func (r AlibabaalihealthdrugkytdrgetproteminfoAPIRequest) GetBatchNo() string {
	return r._batchNo
}

// SetBillOutCode is BillOutCode Setter
// 出库单号
func (r *AlibabaalihealthdrugkytdrgetproteminfoAPIRequest) SetBillOutCode(_billOutCode string) error {
	r._billOutCode = _billOutCode
	r.Set("bill_out_code", _billOutCode)
	return nil
}

// GetBillOutCode BillOutCode Getter
func (r AlibabaalihealthdrugkytdrgetproteminfoAPIRequest) GetBillOutCode() string {
	return r._billOutCode
}
