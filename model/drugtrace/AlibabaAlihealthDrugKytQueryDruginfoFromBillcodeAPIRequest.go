package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest
根据单据编号查询单据明细 API请求
alibaba.alihealth.drug.kyt.query.druginfo.from.billcode

根据单据编号查询单据明细 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.query.druginfo.from.billcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
