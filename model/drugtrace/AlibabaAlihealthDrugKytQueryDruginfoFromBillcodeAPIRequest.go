package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest 根据单据编号查询单据明细 API请求
// alibaba.alihealth.drug.kyt.query.druginfo.from.billcode
//
// 根据单据编号查询单据明细
type AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest struct {
	model.Params
	// 单据号
	_billCode string
	// 企业id
	_refEntId string
}

// NewAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest 初始化AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest对象
func NewAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest() *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest {
	return &AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) Reset() {
	r._billCode = ""
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.query.druginfo.from.billcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest()
	},
}

// GetAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest
func GetAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest() *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest {
	return poolAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest.Get().(*AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest 将 AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest(v *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest.Put(v)
}
