package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUploadinsignAPIRequest 仓库批量扫码回传接口 API请求
// alibaba.alihealth.drug.kyt.uploadinsign
//
// 连锁总部仓库在采购入库或者销售出库环节，批量采集追溯码之后回传到码上放心平台。
type AlibabaAlihealthDrugKytUploadinsignAPIRequest struct {
	model.Params
	// 单据编号（小于20位字符串，唯一）
	_billCode string
	// 单据生成时间
	_billTime string
	// 码上放心平台企业编码（仓库所有者）
	_refUserId string
	// 仓库名称（企业自定义）
	_warehouseId string
	// 药品ID
	_drugId string
	// 追溯码[多个时用逗号分开]
	_traceCodes []string
}

// NewAlibabaAlihealthDrugKytUploadinsignRequest 初始化AlibabaAlihealthDrugKytUploadinsignAPIRequest对象
func NewAlibabaAlihealthDrugKytUploadinsignRequest() *AlibabaAlihealthDrugKytUploadinsignAPIRequest {
	return &AlibabaAlihealthDrugKytUploadinsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadinsignAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.uploadinsign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadinsignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBillCode is BillCode Setter
// 单据编号（小于20位字符串，唯一）
func (r *AlibabaAlihealthDrugKytUploadinsignAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytUploadinsignAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据生成时间
func (r *AlibabaAlihealthDrugKytUploadinsignAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaAlihealthDrugKytUploadinsignAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetRefUserId is RefUserId Setter
// 码上放心平台企业编码（仓库所有者）
func (r *AlibabaAlihealthDrugKytUploadinsignAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaAlihealthDrugKytUploadinsignAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetWarehouseId is WarehouseId Setter
// 仓库名称（企业自定义）
func (r *AlibabaAlihealthDrugKytUploadinsignAPIRequest) SetWarehouseId(_warehouseId string) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r AlibabaAlihealthDrugKytUploadinsignAPIRequest) GetWarehouseId() string {
	return r._warehouseId
}

// SetDrugId is DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytUploadinsignAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// GetDrugId DrugId Getter
func (r AlibabaAlihealthDrugKytUploadinsignAPIRequest) GetDrugId() string {
	return r._drugId
}

// SetTraceCodes is TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytUploadinsignAPIRequest) SetTraceCodes(_traceCodes []string) error {
	r._traceCodes = _traceCodes
	r.Set("trace_codes", _traceCodes)
	return nil
}

// GetTraceCodes TraceCodes Getter
func (r AlibabaAlihealthDrugKytUploadinsignAPIRequest) GetTraceCodes() []string {
	return r._traceCodes
}
