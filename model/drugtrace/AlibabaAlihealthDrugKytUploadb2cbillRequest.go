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
type AlibabaAlihealthDrugKytUploadb2cbillRequest struct {
    model.Params
    // 单据号【20位以内的唯一编码，可以使用16位UUID】
    billCode   string
    // 单据时间【一般情况下写当前时间】
    billTime   string
    // 企业ID
    refUserId   string
    // 操作人
    operIcCode   string
    // 主订单
    masterOrder   string
    // lbx号
    lbxOrder   string
    // 仓号
    warehouseId   string
    // 药品ID
    drugId   string
    // 追溯码[多个时用逗号分开]
    traceCodes   []string
    // 订单来源
    orderSource   string
}

// 初始化AlibabaAlihealthDrugKytUploadb2cbillRequest对象
func NewAlibabaAlihealthDrugKytUploadb2cbillRequest() *AlibabaAlihealthDrugKytUploadb2cbillRequest{
    return &AlibabaAlihealthDrugKytUploadb2cbillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.uploadb2cbill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据号【20位以内的唯一编码，可以使用16位UUID】
func (r *AlibabaAlihealthDrugKytUploadb2cbillRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetBillCode() string {
    return r.billCode
}
// BillTime Setter
// 单据时间【一般情况下写当前时间】
func (r *AlibabaAlihealthDrugKytUploadb2cbillRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetBillTime() string {
    return r.billTime
}
// RefUserId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytUploadb2cbillRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetRefUserId() string {
    return r.refUserId
}
// OperIcCode Setter
// 操作人
func (r *AlibabaAlihealthDrugKytUploadb2cbillRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetOperIcCode() string {
    return r.operIcCode
}
// MasterOrder Setter
// 主订单
func (r *AlibabaAlihealthDrugKytUploadb2cbillRequest) SetMasterOrder(masterOrder string) error {
    r.masterOrder = masterOrder
    r.Set("master_order", masterOrder)
    return nil
}

// MasterOrder Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetMasterOrder() string {
    return r.masterOrder
}
// LbxOrder Setter
// lbx号
func (r *AlibabaAlihealthDrugKytUploadb2cbillRequest) SetLbxOrder(lbxOrder string) error {
    r.lbxOrder = lbxOrder
    r.Set("lbx_order", lbxOrder)
    return nil
}

// LbxOrder Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetLbxOrder() string {
    return r.lbxOrder
}
// WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytUploadb2cbillRequest) SetWarehouseId(warehouseId string) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

// WarehouseId Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetWarehouseId() string {
    return r.warehouseId
}
// DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytUploadb2cbillRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetDrugId() string {
    return r.drugId
}
// TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytUploadb2cbillRequest) SetTraceCodes(traceCodes []string) error {
    r.traceCodes = traceCodes
    r.Set("trace_codes", traceCodes)
    return nil
}

// TraceCodes Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetTraceCodes() []string {
    return r.traceCodes
}
// OrderSource Setter
// 订单来源
func (r *AlibabaAlihealthDrugKytUploadb2cbillRequest) SetOrderSource(orderSource string) error {
    r.orderSource = orderSource
    r.Set("order_source", orderSource)
    return nil
}

// OrderSource Getter
func (r AlibabaAlihealthDrugKytUploadb2cbillRequest) GetOrderSource() string {
    return r.orderSource
}
