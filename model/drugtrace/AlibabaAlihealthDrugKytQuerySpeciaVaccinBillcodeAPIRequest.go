package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest 根据单据编号查询单据明细 API请求
// alibaba.alihealth.drug.kyt.query.specia.vaccin.billcode
//
// 根据单据编号查询单据明细
type AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest struct {
	model.Params
	// 单据号
	_billCode string
	// 企业id
	_refEntId string
}

// NewAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeRequest 初始化AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest对象
func NewAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeRequest() *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest {
	return &AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.query.specia.vaccin.billcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
