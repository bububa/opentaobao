package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest 根据单据号码查询码单据详情和码信息 API请求
// alibaba.alihealth.drug.kyt.query.code.relation.from.billcode
//
// 根据单据号码查询码单据详情和码信息
type AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest struct {
	model.Params
	// 单据号码
	_billCode string
	// 企业refEntId
	_refEntId string
}

// NewAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest 初始化AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest对象
func NewAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest() *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest {
	return &AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.query.code.relation.from.billcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is RefEntId Setter
// 企业refEntId
func (r *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
