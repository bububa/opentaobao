package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传出入库单据(传文件) APIRequest
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

func NewAlibabaAlihealthDrugKytUpinoutfileRequest() *AlibabaAlihealthDrugKytUpinoutfileRequest{
    return &AlibabaAlihealthDrugKytUpinoutfileRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.upinoutfile"
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetBillTime() string {
    return r.billTime
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetBillType() int64 {
    return r.billType
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetPhysicType() int64 {
    return r.physicType
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetRefUserId() string {
    return r.refUserId
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetAgentRefUserId(agentRefUserId string) error {
    r.agentRefUserId = agentRefUserId
    r.Set("agent_ref_user_id", agentRefUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetAgentRefUserId() string {
    return r.agentRefUserId
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetFromUserId() string {
    return r.fromUserId
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetToUserId(toUserId string) error {
    r.toUserId = toUserId
    r.Set("to_user_id", toUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetToUserId() string {
    return r.toUserId
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetDestUserId(destUserId string) error {
    r.destUserId = destUserId
    r.Set("dest_user_id", destUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetDestUserId() string {
    return r.destUserId
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetOperIcCode() string {
    return r.operIcCode
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetOperIcName() string {
    return r.operIcName
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetWarehouseId(warehouseId string) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetWarehouseId() string {
    return r.warehouseId
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetDrugId() string {
    return r.drugId
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetFileContent(fileContent string) error {
    r.fileContent = fileContent
    r.Set("file_content", fileContent)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetFileContent() string {
    return r.fileContent
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetUploadFileName(uploadFileName string) error {
    r.uploadFileName = uploadFileName
    r.Set("upload_file_name", uploadFileName)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetUploadFileName() string {
    return r.uploadFileName
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetClientType() string {
    return r.clientType
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetReturnReasonCode(returnReasonCode string) error {
    r.returnReasonCode = returnReasonCode
    r.Set("return_reason_code", returnReasonCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetReturnReasonCode() string {
    return r.returnReasonCode
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetReturnReasonDes(returnReasonDes string) error {
    r.returnReasonDes = returnReasonDes
    r.Set("return_reason_des", returnReasonDes)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetReturnReasonDes() string {
    return r.returnReasonDes
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetCancelReasonCode(cancelReasonCode string) error {
    r.cancelReasonCode = cancelReasonCode
    r.Set("cancel_reason_code", cancelReasonCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetCancelReasonCode() string {
    return r.cancelReasonCode
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetCancelReasonDes(cancelReasonDes string) error {
    r.cancelReasonDes = cancelReasonDes
    r.Set("cancel_reason_des", cancelReasonDes)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetCancelReasonDes() string {
    return r.cancelReasonDes
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetExecuterName(executerName string) error {
    r.executerName = executerName
    r.Set("executer_name", executerName)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetExecuterName() string {
    return r.executerName
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetExecuterCode(executerCode string) error {
    r.executerCode = executerCode
    r.Set("executer_code", executerCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetExecuterCode() string {
    return r.executerCode
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetSuperviserName(superviserName string) error {
    r.superviserName = superviserName
    r.Set("superviser_name", superviserName)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetSuperviserName() string {
    return r.superviserName
}

func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetSuperviserCode(superviserCode string) error {
    r.superviserCode = superviserCode
    r.Set("superviser_code", superviserCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetSuperviserCode() string {
    return r.superviserCode
}

