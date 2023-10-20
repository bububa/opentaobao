package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytuploadinsignAPIRequest 仓库批量扫码回传接口 API请求
// alibaba.alihealth.drug.kyt.uploadinsign
//
// 连锁总部仓库在采购入库或者销售出库环节，批量采集追溯码之后回传到码上放心平台。
type AlibabaalihealthdrugkytuploadinsignAPIRequest struct {
	model.Params
	// 追溯码[多个时用逗号分开]
	_traceCodes []string
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
}

// NewAlibabaalihealthdrugkytuploadinsignRequest 初始化AlibabaalihealthdrugkytuploadinsignAPIRequest对象
func NewAlibabaalihealthdrugkytuploadinsignRequest() *AlibabaalihealthdrugkytuploadinsignAPIRequest {
	return &AlibabaalihealthdrugkytuploadinsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytuploadinsignAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.uploadinsign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytuploadinsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytuploadinsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceCodes is TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaalihealthdrugkytuploadinsignAPIRequest) SetTraceCodes(_traceCodes []string) error {
	r._traceCodes = _traceCodes
	r.Set("trace_codes", _traceCodes)
	return nil
}

// GetTraceCodes TraceCodes Getter
func (r AlibabaalihealthdrugkytuploadinsignAPIRequest) GetTraceCodes() []string {
	return r._traceCodes
}

// SetBillCode is BillCode Setter
// 单据编号（小于20位字符串，唯一）
func (r *AlibabaalihealthdrugkytuploadinsignAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytuploadinsignAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据生成时间
func (r *AlibabaalihealthdrugkytuploadinsignAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaalihealthdrugkytuploadinsignAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetRefUserId is RefUserId Setter
// 码上放心平台企业编码（仓库所有者）
func (r *AlibabaalihealthdrugkytuploadinsignAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaalihealthdrugkytuploadinsignAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetWarehouseId is WarehouseId Setter
// 仓库名称（企业自定义）
func (r *AlibabaalihealthdrugkytuploadinsignAPIRequest) SetWarehouseId(_warehouseId string) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r AlibabaalihealthdrugkytuploadinsignAPIRequest) GetWarehouseId() string {
	return r._warehouseId
}

// SetDrugId is DrugId Setter
// 药品ID
func (r *AlibabaalihealthdrugkytuploadinsignAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// GetDrugId DrugId Getter
func (r AlibabaalihealthdrugkytuploadinsignAPIRequest) GetDrugId() string {
	return r._drugId
}
