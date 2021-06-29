package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传出入库单据(传文件) API请求
alibaba.alihealth.drug.kyt.upinoutfile

上传出入库单据(传文件)
*/
type AlibabaAlihealthDrugKytUpinoutfileRequest struct {
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
    // 收货企业entId
    fromUserId   string
    // 发货企业entId
    toUserId   string
    // 直调企业标识
    destUserId   string
    // 单据提交者（key编号）
    operIcCode   string
    // 单据提交者姓名
    operIcName   string
    // 仓号
    warehouseId   string
    // 药品ID
    drugId   string
    // 文件内容
    fileContent   string
    // 文件名
    uploadFileName   string
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
}

// 初始化AlibabaAlihealthDrugKytUpinoutfileRequest对象
func NewAlibabaAlihealthDrugKytUpinoutfileRequest() *AlibabaAlihealthDrugKytUpinoutfileRequest{
    return &AlibabaAlihealthDrugKytUpinoutfileRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.upinoutfile"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetBillCode() string {
    return r.billCode
}
// BillTime Setter
// 单据时间
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetBillTime() string {
    return r.billTime
}
// BillType Setter
// 单据类型【102代表采购入库】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetBillType() int64 {
    return r.billType
}
// PhysicType Setter
// 药品类型【3普药2特药】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetPhysicType() int64 {
    return r.physicType
}
// RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetRefUserId() string {
    return r.refUserId
}
// AgentRefUserId Setter
// 代理企业REF标识
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetAgentRefUserId(agentRefUserId string) error {
    r.agentRefUserId = agentRefUserId
    r.Set("agent_ref_user_id", agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetAgentRefUserId() string {
    return r.agentRefUserId
}
// FromUserId Setter
// 收货企业entId
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetFromUserId() string {
    return r.fromUserId
}
// ToUserId Setter
// 发货企业entId
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetToUserId(toUserId string) error {
    r.toUserId = toUserId
    r.Set("to_user_id", toUserId)
    return nil
}

// ToUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetToUserId() string {
    return r.toUserId
}
// DestUserId Setter
// 直调企业标识
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetDestUserId(destUserId string) error {
    r.destUserId = destUserId
    r.Set("dest_user_id", destUserId)
    return nil
}

// DestUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetDestUserId() string {
    return r.destUserId
}
// OperIcCode Setter
// 单据提交者（key编号）
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetOperIcCode() string {
    return r.operIcCode
}
// OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetOperIcName() string {
    return r.operIcName
}
// WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetWarehouseId(warehouseId string) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

// WarehouseId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetWarehouseId() string {
    return r.warehouseId
}
// DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetDrugId() string {
    return r.drugId
}
// FileContent Setter
// 文件内容
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetFileContent(fileContent string) error {
    r.fileContent = fileContent
    r.Set("file_content", fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetFileContent() string {
    return r.fileContent
}
// UploadFileName Setter
// 文件名
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetUploadFileName(uploadFileName string) error {
    r.uploadFileName = uploadFileName
    r.Set("upload_file_name", uploadFileName)
    return nil
}

// UploadFileName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetUploadFileName() string {
    return r.uploadFileName
}
// ClientType Setter
// 客户端类型[必须填2]
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetClientType() string {
    return r.clientType
}
// ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetReturnReasonCode(returnReasonCode string) error {
    r.returnReasonCode = returnReasonCode
    r.Set("return_reason_code", returnReasonCode)
    return nil
}

// ReturnReasonCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetReturnReasonCode() string {
    return r.returnReasonCode
}
// ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetReturnReasonDes(returnReasonDes string) error {
    r.returnReasonDes = returnReasonDes
    r.Set("return_reason_des", returnReasonDes)
    return nil
}

// ReturnReasonDes Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetReturnReasonDes() string {
    return r.returnReasonDes
}
// CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetCancelReasonCode(cancelReasonCode string) error {
    r.cancelReasonCode = cancelReasonCode
    r.Set("cancel_reason_code", cancelReasonCode)
    return nil
}

// CancelReasonCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetCancelReasonCode() string {
    return r.cancelReasonCode
}
// CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetCancelReasonDes(cancelReasonDes string) error {
    r.cancelReasonDes = cancelReasonDes
    r.Set("cancel_reason_des", cancelReasonDes)
    return nil
}

// CancelReasonDes Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetCancelReasonDes() string {
    return r.cancelReasonDes
}
// ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetExecuterName(executerName string) error {
    r.executerName = executerName
    r.Set("executer_name", executerName)
    return nil
}

// ExecuterName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetExecuterName() string {
    return r.executerName
}
// ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetExecuterCode(executerCode string) error {
    r.executerCode = executerCode
    r.Set("executer_code", executerCode)
    return nil
}

// ExecuterCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetExecuterCode() string {
    return r.executerCode
}
// SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetSuperviserName(superviserName string) error {
    r.superviserName = superviserName
    r.Set("superviser_name", superviserName)
    return nil
}

// SuperviserName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetSuperviserName() string {
    return r.superviserName
}
// SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetSuperviserCode(superviserCode string) error {
    r.superviserCode = superviserCode
    r.Set("superviser_code", superviserCode)
    return nil
}

// SuperviserCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetSuperviserCode() string {
    return r.superviserCode
}
