package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest 零售修改出入库单追溯码 API请求
// alibaba.alihealth.drug.kyt.updatebillcode
//
// 零售修改出入库单追溯码
type AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest struct {
	model.Params
	// 追溯码
	_codeList []string
	// 企业ID
	_refEntId string
	// 操作人ID
	_icCode string
	// 单据ID
	_billId string
	// 单据类型
	_billType string
}

// NewAlibabaAlihealthDrugKytUpdatebillcodeRequest 初始化AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest对象
func NewAlibabaAlihealthDrugKytUpdatebillcodeRequest() *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest {
	return &AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) Reset() {
	r._codeList = r._codeList[:0]
	r._refEntId = ""
	r._icCode = ""
	r._billId = ""
	r._billType = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.updatebillcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodeList is CodeList Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) SetCodeList(_codeList []string) error {
	r._codeList = _codeList
	r.Set("code_list", _codeList)
	return nil
}

// GetCodeList CodeList Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetCodeList() []string {
	return r._codeList
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetIcCode is IcCode Setter
// 操作人ID
func (r *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) SetIcCode(_icCode string) error {
	r._icCode = _icCode
	r.Set("ic_code", _icCode)
	return nil
}

// GetIcCode IcCode Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetIcCode() string {
	return r._icCode
}

// SetBillId is BillId Setter
// 单据ID
func (r *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) SetBillId(_billId string) error {
	r._billId = _billId
	r.Set("bill_id", _billId)
	return nil
}

// GetBillId BillId Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetBillId() string {
	return r._billId
}

// SetBillType is BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) GetBillType() string {
	return r._billType
}

var poolAlibabaAlihealthDrugKytUpdatebillcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytUpdatebillcodeRequest()
	},
}

// GetAlibabaAlihealthDrugKytUpdatebillcodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest
func GetAlibabaAlihealthDrugKytUpdatebillcodeAPIRequest() *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest {
	return poolAlibabaAlihealthDrugKytUpdatebillcodeAPIRequest.Get().(*AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytUpdatebillcodeAPIRequest 将 AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytUpdatebillcodeAPIRequest(v *AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytUpdatebillcodeAPIRequest.Put(v)
}
