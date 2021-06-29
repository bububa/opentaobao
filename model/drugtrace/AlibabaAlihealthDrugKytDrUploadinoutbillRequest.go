package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗企业出入库上传 APIRequest
alibaba.alihealth.drug.kyt.dr.uploadinoutbill

零售企业上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）, 
不包括对个人的零售出库，疫苗接种，领药出库。
*/
type AlibabaAlihealthDrugKytDrUploadinoutbillRequest struct {
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

func NewAlibabaAlihealthDrugKytDrUploadinoutbillRequest() *AlibabaAlihealthDrugKytDrUploadinoutbillRequest{
    return &AlibabaAlihealthDrugKytDrUploadinoutbillRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.uploadinoutbill"
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetBillTime() string {
    return r.billTime
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetBillType() int64 {
    return r.billType
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetPhysicType() int64 {
    return r.physicType
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetRefUserId() string {
    return r.refUserId
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetAgentRefUserId(agentRefUserId string) error {
    r.agentRefUserId = agentRefUserId
    r.Set("agent_ref_user_id", agentRefUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetAgentRefUserId() string {
    return r.agentRefUserId
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetFromUserId() string {
    return r.fromUserId
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetToUserId(toUserId string) error {
    r.toUserId = toUserId
    r.Set("to_user_id", toUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetToUserId() string {
    return r.toUserId
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetDestUserId(destUserId string) error {
    r.destUserId = destUserId
    r.Set("dest_user_id", destUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetDestUserId() string {
    return r.destUserId
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetOperIcCode() string {
    return r.operIcCode
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetOperIcName() string {
    return r.operIcName
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetClientType() string {
    return r.clientType
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetReturnReasonCode(returnReasonCode string) error {
    r.returnReasonCode = returnReasonCode
    r.Set("return_reason_code", returnReasonCode)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetReturnReasonCode() string {
    return r.returnReasonCode
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetReturnReasonDes(returnReasonDes string) error {
    r.returnReasonDes = returnReasonDes
    r.Set("return_reason_des", returnReasonDes)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetReturnReasonDes() string {
    return r.returnReasonDes
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetCancelReasonCode(cancelReasonCode string) error {
    r.cancelReasonCode = cancelReasonCode
    r.Set("cancel_reason_code", cancelReasonCode)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetCancelReasonCode() string {
    return r.cancelReasonCode
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetCancelReasonDes(cancelReasonDes string) error {
    r.cancelReasonDes = cancelReasonDes
    r.Set("cancel_reason_des", cancelReasonDes)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetCancelReasonDes() string {
    return r.cancelReasonDes
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetExecuterName(executerName string) error {
    r.executerName = executerName
    r.Set("executer_name", executerName)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetExecuterName() string {
    return r.executerName
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetExecuterCode(executerCode string) error {
    r.executerCode = executerCode
    r.Set("executer_code", executerCode)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetExecuterCode() string {
    return r.executerCode
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetSuperviserName(superviserName string) error {
    r.superviserName = superviserName
    r.Set("superviser_name", superviserName)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetSuperviserName() string {
    return r.superviserName
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetSuperviserCode(superviserCode string) error {
    r.superviserCode = superviserCode
    r.Set("superviser_code", superviserCode)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetSuperviserCode() string {
    return r.superviserCode
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetWarehouseId(warehouseId string) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetWarehouseId() string {
    return r.warehouseId
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetDrugId() string {
    return r.drugId
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetTraceCodes(traceCodes []string) error {
    r.traceCodes = traceCodes
    r.Set("trace_codes", traceCodes)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetTraceCodes() []string {
    return r.traceCodes
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetFromAddress(fromAddress string) error {
    r.fromAddress = fromAddress
    r.Set("from_address", fromAddress)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetFromAddress() string {
    return r.fromAddress
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetToAddress(toAddress string) error {
    r.toAddress = toAddress
    r.Set("to_address", toAddress)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetToAddress() string {
    return r.toAddress
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetFromBillCode(fromBillCode string) error {
    r.fromBillCode = fromBillCode
    r.Set("from_bill_code", fromBillCode)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetFromBillCode() string {
    return r.fromBillCode
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetFromPerson(fromPerson string) error {
    r.fromPerson = fromPerson
    r.Set("from_person", fromPerson)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetFromPerson() string {
    return r.fromPerson
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetToPerson(toPerson string) error {
    r.toPerson = toPerson
    r.Set("to_person", toPerson)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetToPerson() string {
    return r.toPerson
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetDisRefEntId(disRefEntId string) error {
    r.disRefEntId = disRefEntId
    r.Set("dis_ref_ent_id", disRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetDisRefEntId() string {
    return r.disRefEntId
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetDisEntId(disEntId string) error {
    r.disEntId = disEntId
    r.Set("dis_ent_id", disEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetDisEntId() string {
    return r.disEntId
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetQuReceivable(quReceivable int64) error {
    r.quReceivable = quReceivable
    r.Set("qu_receivable", quReceivable)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetQuReceivable() int64 {
    return r.quReceivable
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetXtIsCheck(xtIsCheck string) error {
    r.xtIsCheck = xtIsCheck
    r.Set("xt_is_check", xtIsCheck)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetXtIsCheck() string {
    return r.xtIsCheck
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetXtCheckCode(xtCheckCode string) error {
    r.xtCheckCode = xtCheckCode
    r.Set("xt_check_code", xtCheckCode)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetXtCheckCode() string {
    return r.xtCheckCode
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetXtCheckCodeDesc(xtCheckCodeDesc string) error {
    r.xtCheckCodeDesc = xtCheckCodeDesc
    r.Set("xt_check_code_desc", xtCheckCodeDesc)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetXtCheckCodeDesc() string {
    return r.xtCheckCodeDesc
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetDrugListJson(drugListJson string) error {
    r.drugListJson = drugListJson
    r.Set("drug_list_json", drugListJson)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetDrugListJson() string {
    return r.drugListJson
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetAssRefEntId(assRefEntId string) error {
    r.assRefEntId = assRefEntId
    r.Set("ass_ref_ent_id", assRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetAssRefEntId() string {
    return r.assRefEntId
}

func (r *AlibabaAlihealthDrugKytDrUploadinoutbillRequest) SetAssEntId(assEntId string) error {
    r.assEntId = assEntId
    r.Set("ass_ent_id", assEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrUploadinoutbillRequest) GetAssEntId() string {
    return r.assEntId
}

