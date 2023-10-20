package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrBillcheckAPIRequest 疫苗追溯验证 API请求
// alibaba.alihealth.drug.kyt.dr.billcheck
//
// 各级疾控在入库完成后，需要做追溯信息验证
type AlibabaAlihealthDrugKytDrBillcheckAPIRequest struct {
	model.Params
	// 调用企业ID
	_refEntId string
	// 单据编号
	_billCode string
	// 单据类型
	_billType string
	// 单据企业refEntId
	_owerRefEntId string
}

// NewAlibabaAlihealthDrugKytDrBillcheckRequest 初始化AlibabaAlihealthDrugKytDrBillcheckAPIRequest对象
func NewAlibabaAlihealthDrugKytDrBillcheckRequest() *AlibabaAlihealthDrugKytDrBillcheckAPIRequest {
	return &AlibabaAlihealthDrugKytDrBillcheckAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytDrBillcheckAPIRequest) Reset() {
	r._refEntId = ""
	r._billCode = ""
	r._billType = ""
	r._owerRefEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.billcheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用企业ID
func (r *AlibabaAlihealthDrugKytDrBillcheckAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBillCode is BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytDrBillcheckAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillType is BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytDrBillcheckAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetBillType() string {
	return r._billType
}

// SetOwerRefEntId is OwerRefEntId Setter
// 单据企业refEntId
func (r *AlibabaAlihealthDrugKytDrBillcheckAPIRequest) SetOwerRefEntId(_owerRefEntId string) error {
	r._owerRefEntId = _owerRefEntId
	r.Set("ower_ref_ent_id", _owerRefEntId)
	return nil
}

// GetOwerRefEntId OwerRefEntId Getter
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetOwerRefEntId() string {
	return r._owerRefEntId
}

var poolAlibabaAlihealthDrugKytDrBillcheckAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytDrBillcheckRequest()
	},
}

// GetAlibabaAlihealthDrugKytDrBillcheckRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrBillcheckAPIRequest
func GetAlibabaAlihealthDrugKytDrBillcheckAPIRequest() *AlibabaAlihealthDrugKytDrBillcheckAPIRequest {
	return poolAlibabaAlihealthDrugKytDrBillcheckAPIRequest.Get().(*AlibabaAlihealthDrugKytDrBillcheckAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytDrBillcheckAPIRequest 将 AlibabaAlihealthDrugKytDrBillcheckAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrBillcheckAPIRequest(v *AlibabaAlihealthDrugKytDrBillcheckAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrBillcheckAPIRequest.Put(v)
}
