package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出入库单据上传 APIRequest
alibaba.alihealth.drugtrace.top.yljg.uploadinoutbill

零售企业上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）, 
不包括对个人的零售出库，疫苗接种，领药出库。
*/
type AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest struct {
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

    // （协同平台数据合规）发货地址【必选】
    fromAddress   string 

    // （协同平台数据合规）收货地址【必选】
    toAddress   string 

    // （协同平台数据合规）发货单编号【必选】
    fromBillCode   string 

    // （协同平台数据合规）订货单编号【可选】
    orderCode   string 

    // （协同平台数据合规）发货人【必选】
    fromPerson   string 

    // （协同平台数据合规）收货人【必选】
    toPerson   string 

    // （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
    disRefEntId   string 

    // （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
    disEntId   string 

    // （协同平台数据合规）应收货总数量【必选】
    quReceivable   int64 

    // （协同平台数据合规）是否验证，0：未通过验证，1：已验证
    xtIsCheck   string 

    // （协同平台数据合规）未验证通过原因【验证未通过时填写】
    xtCheckCode   string 

    // （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
    xtCheckCodeDesc   string 

    // （协同平台数据合规）药品列表Json
    drugListJson   string 

    // （协同平台数据合规）单据委托企业refEntId【疫苗药品出库单填写】
    assRefEntId   string 

    // （协同平台数据合规）单据委托企业entId【疫苗药品出库单填写】
    assEntId   string 

}

func NewAlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest() *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest{
    return &AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.uploadinoutbill"
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetBillTime() string {
    return r.billTime
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetBillType() int64 {
    return r.billType
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetPhysicType() int64 {
    return r.physicType
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetRefUserId() string {
    return r.refUserId
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetAgentRefUserId(agentRefUserId string) error {
    r.agentRefUserId = agentRefUserId
    r.Set("agent_ref_user_id", agentRefUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetAgentRefUserId() string {
    return r.agentRefUserId
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetFromUserId() string {
    return r.fromUserId
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetToUserId(toUserId string) error {
    r.toUserId = toUserId
    r.Set("to_user_id", toUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetToUserId() string {
    return r.toUserId
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetDestUserId(destUserId string) error {
    r.destUserId = destUserId
    r.Set("dest_user_id", destUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetDestUserId() string {
    return r.destUserId
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetOperIcCode() string {
    return r.operIcCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetOperIcName() string {
    return r.operIcName
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetClientType() string {
    return r.clientType
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetReturnReasonCode(returnReasonCode string) error {
    r.returnReasonCode = returnReasonCode
    r.Set("return_reason_code", returnReasonCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetReturnReasonCode() string {
    return r.returnReasonCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetReturnReasonDes(returnReasonDes string) error {
    r.returnReasonDes = returnReasonDes
    r.Set("return_reason_des", returnReasonDes)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetReturnReasonDes() string {
    return r.returnReasonDes
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetCancelReasonCode(cancelReasonCode string) error {
    r.cancelReasonCode = cancelReasonCode
    r.Set("cancel_reason_code", cancelReasonCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetCancelReasonCode() string {
    return r.cancelReasonCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetCancelReasonDes(cancelReasonDes string) error {
    r.cancelReasonDes = cancelReasonDes
    r.Set("cancel_reason_des", cancelReasonDes)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetCancelReasonDes() string {
    return r.cancelReasonDes
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetExecuterName(executerName string) error {
    r.executerName = executerName
    r.Set("executer_name", executerName)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetExecuterName() string {
    return r.executerName
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetExecuterCode(executerCode string) error {
    r.executerCode = executerCode
    r.Set("executer_code", executerCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetExecuterCode() string {
    return r.executerCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetSuperviserName(superviserName string) error {
    r.superviserName = superviserName
    r.Set("superviser_name", superviserName)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetSuperviserName() string {
    return r.superviserName
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetSuperviserCode(superviserCode string) error {
    r.superviserCode = superviserCode
    r.Set("superviser_code", superviserCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetSuperviserCode() string {
    return r.superviserCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetWarehouseId(warehouseId string) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetWarehouseId() string {
    return r.warehouseId
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetDrugId() string {
    return r.drugId
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetTraceCodes(traceCodes []string) error {
    r.traceCodes = traceCodes
    r.Set("trace_codes", traceCodes)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetTraceCodes() []string {
    return r.traceCodes
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetFromAddress(fromAddress string) error {
    r.fromAddress = fromAddress
    r.Set("from_address", fromAddress)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetFromAddress() string {
    return r.fromAddress
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetToAddress(toAddress string) error {
    r.toAddress = toAddress
    r.Set("to_address", toAddress)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetToAddress() string {
    return r.toAddress
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetFromBillCode(fromBillCode string) error {
    r.fromBillCode = fromBillCode
    r.Set("from_bill_code", fromBillCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetFromBillCode() string {
    return r.fromBillCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetFromPerson(fromPerson string) error {
    r.fromPerson = fromPerson
    r.Set("from_person", fromPerson)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetFromPerson() string {
    return r.fromPerson
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetToPerson(toPerson string) error {
    r.toPerson = toPerson
    r.Set("to_person", toPerson)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetToPerson() string {
    return r.toPerson
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetDisRefEntId(disRefEntId string) error {
    r.disRefEntId = disRefEntId
    r.Set("dis_ref_ent_id", disRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetDisRefEntId() string {
    return r.disRefEntId
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetDisEntId(disEntId string) error {
    r.disEntId = disEntId
    r.Set("dis_ent_id", disEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetDisEntId() string {
    return r.disEntId
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetQuReceivable(quReceivable int64) error {
    r.quReceivable = quReceivable
    r.Set("qu_receivable", quReceivable)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetQuReceivable() int64 {
    return r.quReceivable
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetXtIsCheck(xtIsCheck string) error {
    r.xtIsCheck = xtIsCheck
    r.Set("xt_is_check", xtIsCheck)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetXtIsCheck() string {
    return r.xtIsCheck
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetXtCheckCode(xtCheckCode string) error {
    r.xtCheckCode = xtCheckCode
    r.Set("xt_check_code", xtCheckCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetXtCheckCode() string {
    return r.xtCheckCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetXtCheckCodeDesc(xtCheckCodeDesc string) error {
    r.xtCheckCodeDesc = xtCheckCodeDesc
    r.Set("xt_check_code_desc", xtCheckCodeDesc)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetXtCheckCodeDesc() string {
    return r.xtCheckCodeDesc
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetDrugListJson(drugListJson string) error {
    r.drugListJson = drugListJson
    r.Set("drug_list_json", drugListJson)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetDrugListJson() string {
    return r.drugListJson
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetAssRefEntId(assRefEntId string) error {
    r.assRefEntId = assRefEntId
    r.Set("ass_ref_ent_id", assRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetAssRefEntId() string {
    return r.assRefEntId
}

func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) SetAssEntId(assEntId string) error {
    r.assEntId = assEntId
    r.Set("ass_ent_id", assEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest) GetAssEntId() string {
    return r.assEntId
}

