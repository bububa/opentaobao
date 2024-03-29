package drugtrace

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) Reset() {
	r._billCode = ""
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.query.code.relation.from.billcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillCode is BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetRefEntId is RefEntId Setter
// 企业refEntId
func (r *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest()
	},
}

// GetAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest
func GetAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest() *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest {
	return poolAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest.Get().(*AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest 将 AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest(v *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest.Put(v)
}
