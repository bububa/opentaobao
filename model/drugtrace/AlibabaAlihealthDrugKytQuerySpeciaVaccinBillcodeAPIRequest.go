package drugtrace

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest) Reset() {
	r._billCode = ""
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.query.specia.vaccin.billcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeRequest()
	},
}

// GetAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest
func GetAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest() *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest {
	return poolAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest.Get().(*AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest 将 AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest(v *AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest.Put(v)
}
