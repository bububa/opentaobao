package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
快易通零售B2C API请求
alibaba.alihealth.drug.kyt.uploadb2cbill

快易通零售B2C单据上传
*/
type AlibabaAlihealthDrugKytUploadb2cbillAPIRequest struct {
    model.Params
    // 单据号【20位以内的唯一编码，可以使用16位UUID】
    _billCode   string
    // 单据时间【一般情况下写当前时间】
    _billTime   string
    // 企业ID
    _refUserId   string
    // 操作人
    _operIcCode   string
    // 主订单
    _masterOrder   string
    // lbx号
    _lbxOrder   string
    // 仓号
    _warehouseId   string
    // 药品ID
    _drugId   string
    // 追溯码[多个时用逗号分开]
    _traceCodes   []string
    // 订单来源
    _orderSource   string
}

// 初始化AlibabaAlihealthDrugKytUploadb2cbillAPIRequest对象
func NewAlibabaAlihealthDrugKytUploadb2cbillRequest() *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest{
    return &AlibabaAlihealthDrugKytUploadb2cbillAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.uploadb2cbill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据号【20位以内的唯一编码，可以使用16位UUID】
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetBillCode() string {
    return r._billCode
}
// BillTime Setter
// 单据时间【一般情况下写当前时间】
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetBillTime(_billTime string) error {
    r._billTime = _billTime
    r.Set("bill_time", _billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetBillTime() string {
    return r._billTime
}
// RefUserId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetRefUserId(_refUserId string) error {
    r._refUserId = _refUserId
    r.Set("ref_user_id", _refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetRefUserId() string {
    return r._refUserId
}
// OperIcCode Setter
// 操作人
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetOperIcCode(_operIcCode string) error {
    r._operIcCode = _operIcCode
    r.Set("oper_ic_code", _operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetOperIcCode() string {
    return r._operIcCode
}
// MasterOrder Setter
// 主订单
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetMasterOrder(_masterOrder string) error {
    r._masterOrder = _masterOrder
    r.Set("master_order", _masterOrder)
    return nil
}

// MasterOrder Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetMasterOrder() string {
    return r._masterOrder
}
// LbxOrder Setter
// lbx号
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetLbxOrder(_lbxOrder string) error {
    r._lbxOrder = _lbxOrder
    r.Set("lbx_order", _lbxOrder)
    return nil
}

// LbxOrder Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetLbxOrder() string {
    return r._lbxOrder
}
// WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetWarehouseId(_warehouseId string) error {
    r._warehouseId = _warehouseId
    r.Set("warehouse_id", _warehouseId)
    return nil
}

// WarehouseId Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetWarehouseId() string {
    return r._warehouseId
}
// DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetDrugId(_drugId string) error {
    r._drugId = _drugId
    r.Set("drug_id", _drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetDrugId() string {
    return r._drugId
}
// TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetTraceCodes(_traceCodes []string) error {
    r._traceCodes = _traceCodes
    r.Set("trace_codes", _traceCodes)
    return nil
}

// TraceCodes Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetTraceCodes() []string {
    return r._traceCodes
}
// OrderSource Setter
// 订单来源
func (r *AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) SetOrderSource(_orderSource string) error {
    r._orderSource = _orderSource
    r.Set("order_source", _orderSource)
    return nil
}

// OrderSource Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillAPIRequest) GetOrderSource() string {
    return r._orderSource
}
