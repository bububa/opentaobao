package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
医院出入库信息上传 APIRequest
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

func NewAlibabaAlihealthDrugKytYyUploadinoutbillRequest() *AlibabaAlihealthDrugKytYyUploadinoutbillRequest{
    return &AlibabaAlihealthDrugKytYyUploadinoutbillRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yy.uploadinoutbill"
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetBillTime() string {
    return r.billTime
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetBillType() int64 {
    return r.billType
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetPhysicType() int64 {
    return r.physicType
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetRefUserId() string {
    return r.refUserId
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetAgentRefUserId(agentRefUserId string) error {
    r.agentRefUserId = agentRefUserId
    r.Set("agent_ref_user_id", agentRefUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetAgentRefUserId() string {
    return r.agentRefUserId
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetFromUserId() string {
    return r.fromUserId
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetToUserId(toUserId string) error {
    r.toUserId = toUserId
    r.Set("to_user_id", toUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetToUserId() string {
    return r.toUserId
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetDestUserId(destUserId string) error {
    r.destUserId = destUserId
    r.Set("dest_user_id", destUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetDestUserId() string {
    return r.destUserId
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetOperIcCode() string {
    return r.operIcCode
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetOperIcName() string {
    return r.operIcName
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetClientType() string {
    return r.clientType
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetReturnReasonCode(returnReasonCode string) error {
    r.returnReasonCode = returnReasonCode
    r.Set("return_reason_code", returnReasonCode)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetReturnReasonCode() string {
    return r.returnReasonCode
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetReturnReasonDes(returnReasonDes string) error {
    r.returnReasonDes = returnReasonDes
    r.Set("return_reason_des", returnReasonDes)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetReturnReasonDes() string {
    return r.returnReasonDes
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetCancelReasonCode(cancelReasonCode string) error {
    r.cancelReasonCode = cancelReasonCode
    r.Set("cancel_reason_code", cancelReasonCode)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetCancelReasonCode() string {
    return r.cancelReasonCode
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetCancelReasonDes(cancelReasonDes string) error {
    r.cancelReasonDes = cancelReasonDes
    r.Set("cancel_reason_des", cancelReasonDes)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetCancelReasonDes() string {
    return r.cancelReasonDes
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetExecuterName(executerName string) error {
    r.executerName = executerName
    r.Set("executer_name", executerName)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetExecuterName() string {
    return r.executerName
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetExecuterCode(executerCode string) error {
    r.executerCode = executerCode
    r.Set("executer_code", executerCode)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetExecuterCode() string {
    return r.executerCode
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetSuperviserName(superviserName string) error {
    r.superviserName = superviserName
    r.Set("superviser_name", superviserName)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetSuperviserName() string {
    return r.superviserName
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetSuperviserCode(superviserCode string) error {
    r.superviserCode = superviserCode
    r.Set("superviser_code", superviserCode)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetSuperviserCode() string {
    return r.superviserCode
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetWarehouseId(warehouseId string) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetWarehouseId() string {
    return r.warehouseId
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetDrugId() string {
    return r.drugId
}

func (r *AlibabaAlihealthDrugKytYyUploadinoutbillRequest) SetTraceCodes(traceCodes []string) error {
    r.traceCodes = traceCodes
    r.Set("trace_codes", traceCodes)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadinoutbillRequest) GetTraceCodes() []string {
    return r.traceCodes
}

