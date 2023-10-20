package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUploadb2cbillAPIRequest 快易通零售B2C API请求
// alibaba.alihealth.drug.kyt.uploadb2cbill
//
// 快易通零售B2C单据上传
type AlibabaAlihealthDrugKytUploadb2cbillAPIRequest struct {
	model.Params
	// 追溯码[多个时用逗号分开]
	_traceCodes []string
	// 单据号【20位以内的唯一编码，可以使用16位UUID】
	_billCode string
	// 单据时间【一般情况下写当前时间】
	_billTime string
	// 企业ID
	_refUserId string
	// 操作人
	_operIcCode string
	// 主订单
	_masterOrder string
	// lbx号
	_lbxOrder string
	// 仓号
	_warehouseId string
	// 药品ID
	_drugId string
	// 订单来源
	_orderSource string
}

// NewAlibabaAlihealthDrugKytUploadb2cbillRequest 初始化AlibabaAlihealthDrugKytUploadb2cbillAPIRequest对象
func NewAlibabaAlihealthDrugKytUploadb2cbillRequest() *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest {
	return &AlibabaAlihealthDrugKytUploadb2cbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.uploadb2cbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceCodes is TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetTraceCodes(_traceCodes []string) error {
	r._traceCodes = _traceCodes
	r.Set("trace_codes", _traceCodes)
	return nil
}

// GetTraceCodes TraceCodes Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetTraceCodes() []string {
	return r._traceCodes
}

// SetBillCode is BillCode Setter
// 单据号【20位以内的唯一编码，可以使用16位UUID】
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据时间【一般情况下写当前时间】
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetRefUserId is RefUserId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetOperIcCode is OperIcCode Setter
// 操作人
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// GetOperIcCode OperIcCode Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// SetMasterOrder is MasterOrder Setter
// 主订单
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetMasterOrder(_masterOrder string) error {
	r._masterOrder = _masterOrder
	r.Set("master_order", _masterOrder)
	return nil
}

// GetMasterOrder MasterOrder Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetMasterOrder() string {
	return r._masterOrder
}

// SetLbxOrder is LbxOrder Setter
// lbx号
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetLbxOrder(_lbxOrder string) error {
	r._lbxOrder = _lbxOrder
	r.Set("lbx_order", _lbxOrder)
	return nil
}

// GetLbxOrder LbxOrder Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetLbxOrder() string {
	return r._lbxOrder
}

// SetWarehouseId is WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetWarehouseId(_warehouseId string) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetWarehouseId() string {
	return r._warehouseId
}

// SetDrugId is DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// GetDrugId DrugId Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetDrugId() string {
	return r._drugId
}

// SetOrderSource is OrderSource Setter
// 订单来源
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetOrderSource(_orderSource string) error {
	r._orderSource = _orderSource
	r.Set("order_source", _orderSource)
	return nil
}

// GetOrderSource OrderSource Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetOrderSource() string {
	return r._orderSource
}
