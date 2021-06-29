package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
医院出入库信息上传 API请求
alibaba.alihealth.drug.kyt.yy.uploadinoutbill

医院上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）, 
不包括对个人的零售出库，疫苗接种，领药出库。
*/
type AlibabaAlihealthDrugKytYyUploadinoutbillRequest struct {
    model.Params
    // 单据编码
    billCode   string
    // 单据时间
    billTime   string
    // 单据类型【102代表采购入库】
    billType   int64
    // 药品类型【3普药2特药】
    physicType   int64
    // 上传企业的单位编码
    refUserId   string
    // 代理企业REF标识
    agentRefUserId   string
    // 发货企业entId
    fromUserId   string
    // 收货企业entId
    toUserId   string
    // 直调企业标识
    destUserId   string
    // 单据提交者（appkey编号）
    operIcCode   string
    // 单据提交者姓名
    operIcName   string
    // 客户端类型[必须填2]
    clientType   string
    // 退货原因代码[退货入出库时填写]
    returnReasonCode   string
    // 退货原因描述[退货入出库时填写]
    returnReasonDes   string
    // 注销原因代码【销毁出库时填写】
    cancelReasonCode   string
    // 注销原因描述【销毁出库时填写】
    cancelReasonDes   string
    // 执行人姓名【销毁出库时填写】
    executerName   string
    // 执行人证件号【销毁出库时填写】
    executerCode   string
    // 监督人姓名【销毁出库时填写】
    superviserName   string
    // 监督人证件号【销毁出库时填写】
    superviserCode   string
    // 仓号
    warehouseId   string
    // 药品ID[企业自已系统的药品ID]
    drugId   string
    // 追溯码[多个时用逗号分开]
    traceCodes   []string
}

// 初始化AlibabaAlihealthDrugKytYyUploadinoutbillRequest对象
func NewAlibabaAlihealthDrugKytYyUploadinoutbillRequest() *AlibabaAlihealthDrugKytYyUploadinoutbillRequest{
    return &AlibabaAlihealthDrugKytYyUploadinoutbillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yy.uploadinoutbill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetBillCode() string {
    return r.billCode
}
// BillTime Setter
// 单据时间
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetBillTime() string {
    return r.billTime
}
// BillType Setter
// 单据类型【102代表采购入库】
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetBillType() int64 {
    return r.billType
}
// PhysicType Setter
// 药品类型【3普药2特药】
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetPhysicType() int64 {
    return r.physicType
}
// RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetRefUserId() string {
    return r.refUserId
}
// AgentRefUserId Setter
// 代理企业REF标识
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetAgentRefUserId(agentRefUserId string) error {
    r.agentRefUserId = agentRefUserId
    r.Set("agent_ref_user_id", agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetAgentRefUserId() string {
    return r.agentRefUserId
}
// FromUserId Setter
// 发货企业entId
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetFromUserId() string {
    return r.fromUserId
}
// ToUserId Setter
// 收货企业entId
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetToUserId(toUserId string) error {
    r.toUserId = toUserId
    r.Set("to_user_id", toUserId)
    return nil
}

// ToUserId Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetToUserId() string {
    return r.toUserId
}
// DestUserId Setter
// 直调企业标识
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetDestUserId(destUserId string) error {
    r.destUserId = destUserId
    r.Set("dest_user_id", destUserId)
    return nil
}

// DestUserId Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetDestUserId() string {
    return r.destUserId
}
// OperIcCode Setter
// 单据提交者（appkey编号）
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetOperIcCode() string {
    return r.operIcCode
}
// OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetOperIcName() string {
    return r.operIcName
}
// ClientType Setter
// 客户端类型[必须填2]
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetClientType() string {
    return r.clientType
}
// ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetReturnReasonCode(returnReasonCode string) error {
    r.returnReasonCode = returnReasonCode
    r.Set("return_reason_code", returnReasonCode)
    return nil
}

// ReturnReasonCode Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetReturnReasonCode() string {
    return r.returnReasonCode
}
// ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetReturnReasonDes(returnReasonDes string) error {
    r.returnReasonDes = returnReasonDes
    r.Set("return_reason_des", returnReasonDes)
    return nil
}

// ReturnReasonDes Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetReturnReasonDes() string {
    return r.returnReasonDes
}
// CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetCancelReasonCode(cancelReasonCode string) error {
    r.cancelReasonCode = cancelReasonCode
    r.Set("cancel_reason_code", cancelReasonCode)
    return nil
}

// CancelReasonCode Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetCancelReasonCode() string {
    return r.cancelReasonCode
}
// CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetCancelReasonDes(cancelReasonDes string) error {
    r.cancelReasonDes = cancelReasonDes
    r.Set("cancel_reason_des", cancelReasonDes)
    return nil
}

// CancelReasonDes Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetCancelReasonDes() string {
    return r.cancelReasonDes
}
// ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetExecuterName(executerName string) error {
    r.executerName = executerName
    r.Set("executer_name", executerName)
    return nil
}

// ExecuterName Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetExecuterName() string {
    return r.executerName
}
// ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetExecuterCode(executerCode string) error {
    r.executerCode = executerCode
    r.Set("executer_code", executerCode)
    return nil
}

// ExecuterCode Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetExecuterCode() string {
    return r.executerCode
}
// SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetSuperviserName(superviserName string) error {
    r.superviserName = superviserName
    r.Set("superviser_name", superviserName)
    return nil
}

// SuperviserName Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetSuperviserName() string {
    return r.superviserName
}
// SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetSuperviserCode(superviserCode string) error {
    r.superviserCode = superviserCode
    r.Set("superviser_code", superviserCode)
    return nil
}

// SuperviserCode Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetSuperviserCode() string {
    return r.superviserCode
}
// WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetWarehouseId(warehouseId string) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

// WarehouseId Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetWarehouseId() string {
    return r.warehouseId
}
// DrugId Setter
// 药品ID[企业自已系统的药品ID]
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetDrugId() string {
    return r.drugId
}
// TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetTraceCodes(traceCodes []string) error {
    r.traceCodes = traceCodes
    r.Set("trace_codes", traceCodes)
    return nil
}

// TraceCodes Getter
func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetTraceCodes() []string {
    return r.traceCodes
}
