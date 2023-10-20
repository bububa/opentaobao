package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytStorebilldeleteAPIRequest 零售端单据删除 API请求
// alibaba.alihealth.drug.kyt.storebilldelete
//
// 零售端单据删除
type AlibabaAlihealthDrugKytStorebilldeleteAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 操作人编码
	_icCode string
	// 单据ID
	_billId string
	// 单据类型
	_billType string
}

// NewAlibabaAlihealthDrugKytStorebilldeleteRequest 初始化AlibabaAlihealthDrugKytStorebilldeleteAPIRequest对象
func NewAlibabaAlihealthDrugKytStorebilldeleteRequest() *AlibabaAlihealthDrugKytStorebilldeleteAPIRequest {
	return &AlibabaAlihealthDrugKytStorebilldeleteAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) Reset() {
	r._refEntId = ""
	r._icCode = ""
	r._billId = ""
	r._billType = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.storebilldelete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetIcCode is IcCode Setter
// 操作人编码
func (r *AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) SetIcCode(_icCode string) error {
	r._icCode = _icCode
	r.Set("ic_code", _icCode)
	return nil
}

// GetIcCode IcCode Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) GetIcCode() string {
	return r._icCode
}

// SetBillId is BillId Setter
// 单据ID
func (r *AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) SetBillId(_billId string) error {
	r._billId = _billId
	r.Set("bill_id", _billId)
	return nil
}

// GetBillId BillId Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) GetBillId() string {
	return r._billId
}

// SetBillType is BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) GetBillType() string {
	return r._billType
}

var poolAlibabaAlihealthDrugKytStorebilldeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytStorebilldeleteRequest()
	},
}

// GetAlibabaAlihealthDrugKytStorebilldeleteRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytStorebilldeleteAPIRequest
func GetAlibabaAlihealthDrugKytStorebilldeleteAPIRequest() *AlibabaAlihealthDrugKytStorebilldeleteAPIRequest {
	return poolAlibabaAlihealthDrugKytStorebilldeleteAPIRequest.Get().(*AlibabaAlihealthDrugKytStorebilldeleteAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytStorebilldeleteAPIRequest 将 AlibabaAlihealthDrugKytStorebilldeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytStorebilldeleteAPIRequest(v *AlibabaAlihealthDrugKytStorebilldeleteAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytStorebilldeleteAPIRequest.Put(v)
}
