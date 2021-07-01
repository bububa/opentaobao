package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest
零售修改出入库单追溯码 API请求
alibaba.alihealth.drug.kyt.updatebillcode

零售修改出入库单追溯码 */
type AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 操作人ID
	_icCode string
	// 单据ID
	_billId string
	// 单据类型
	_billType string
	// 追溯码
	_codeList []string
}

// NewAlibabaAlihealthDrugKytUpdatebillcodeRequest 初始化AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest对象
func NewAlibabaAlihealthDrugKytUpdatebillcodeRequest() *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest {
	return &AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.updatebillcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is IcCode Setter
// 操作人ID
func (r *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) SetIcCode(_icCode string) error {
	r._icCode = _icCode
	r.Set("ic_code", _icCode)
	return nil
}

// Get IcCode Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetIcCode() string {
	return r._icCode
}

// Set is BillId Setter
// 单据ID
func (r *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) SetBillId(_billId string) error {
	r._billId = _billId
	r.Set("bill_id", _billId)
	return nil
}

// Get BillId Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetBillId() string {
	return r._billId
}

// Set is BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// Get BillType Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetBillType() string {
	return r._billType
}

// Set is CodeList Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// Get CodeList Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetCodeList() []string {
	return r._codeList
}
