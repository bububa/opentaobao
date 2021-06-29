package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生产批发单据上传 API请求
alibaba.alihealth.drug.kyt.uploadcircubill

生产批发单据上传（非零售企业使用），包括101, "生产入库"；102, "采购入库"；103, "退货入库"；104, "调拨入库"；106, "零头入库"；107, "供应入库"；108, "召回入库"；110,"赠品入库"；111,"盘盈入库"；112,"报废入库"；113,"其他入库"
201, "销售出库"；202, "退货出库"；203, "调拨出库"；204, "返工出库"；205, "销毁出库"；206, "抽检出库"；207, "直调出库"；208, "生产出库"；209, "供应出库"；211, "召回出库"；212,"赠品出库"；214,"盘亏出库"；215,"损坏出库"；216,"报废出库"；217,"其他出库"；237, "直调退货"。
*/
type AlibabaAlihealthDrugKytUploadcircubillRequest struct {
    model.Params
    // 单据编号【同一个企业不能上传相同单据号】
    billCode   string
    // 单据时间（扫码时间）
    billTime   string
    // 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
    billType   int64
    // 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药
    physicType   int64
    // 货主企业（单据的所有者，上传人）【注意：该入参是ref_ent_id，不是ent_id】
    refUserId   string
    // 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
    agentRefUserId   string
    // 发货企业【注意：该入参是ent_id,并不是ref_ent_id】
    fromUserId   string
    // 收货企业【注意：该入参是ent_id,并不是ref_ent_id】
    toUserId   string
    // 直调企业【注意：该入参是ent_id,并不是ref_ent_id】
    destUserId   string
    // 操作人标识（appkey编号）
    operIcCode   string
    // 单据提交者姓名
    operIcName   string
    // 单据文件体【bas64字符串】，看对接文档中的代码示例，示例中有相应说明。
    fileContent   string
    // 单据名称
    uploadFileName   string
    // 客户端类型【暂定都写2】
    clientType   string
    // （协同平台数据合规）应收货总数量
    quReceivable   int64
    // （协同平台数据合规）是否验证，0：未通过验证，1：已验证
    xtIsCheck   string
    // （协同平台数据合规）未验证通过原因【验证未通过时填写】
    xtCheckCode   string
    // （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
    xtCheckCodeDesc   string
    // （协同平台数据合规）药品列表Json："codeCount":         药品数量 "commDrugId":     国家药品唯一标识 "exprieDate":         生产日期 "physicInfo":          药品信息 "pkgSpec":           包状规格 "prepnCount":       制剂数量 "produceBatchNo":生产批次 "produceDate":      生产日期
    drugListJson   string
    // （协同平台数据合规）单据委托企业refEntId
    assRefEntId   string
    // （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
    disEntId   string
    // （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
    disRefEntId   string
    // （协同平台数据合规）收货人【必选】
    toPerson   string
    // （协同平台数据合规）发货人【必选】
    fromPerson   string
    // （协同平台数据合规）订货单编号【可选】
    orderCode   string
    // （协同平台数据合规）发货单编号【必选】
    fromBillCode   string
    // （协同平台数据合规）收货地址【必选】
    toAddress   string
    // （协同平台数据合规）发货地址【必选】
    fromAddress   string
    // （协同平台数据合规）单据委托企业entId
    assEntId   string
}

// 初始化AlibabaAlihealthDrugKytUploadcircubillRequest对象
func NewAlibabaAlihealthDrugKytUploadcircubillRequest() *AlibabaAlihealthDrugKytUploadcircubillRequest{
    return &AlibabaAlihealthDrugKytUploadcircubillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.uploadcircubill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编号【同一个企业不能上传相同单据号】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetBillCode() string {
    return r.billCode
}
// BillTime Setter
// 单据时间（扫码时间）
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetBillTime() string {
    return r.billTime
}
// BillType Setter
// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetBillType() int64 {
    return r.billType
}
// PhysicType Setter
// 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetPhysicType() int64 {
    return r.physicType
}
// RefUserId Setter
// 货主企业（单据的所有者，上传人）【注意：该入参是ref_ent_id，不是ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetRefUserId() string {
    return r.refUserId
}
// AgentRefUserId Setter
// 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetAgentRefUserId(agentRefUserId string) error {
    r.agentRefUserId = agentRefUserId
    r.Set("agent_ref_user_id", agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetAgentRefUserId() string {
    return r.agentRefUserId
}
// FromUserId Setter
// 发货企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetFromUserId() string {
    return r.fromUserId
}
// ToUserId Setter
// 收货企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetToUserId(toUserId string) error {
    r.toUserId = toUserId
    r.Set("to_user_id", toUserId)
    return nil
}

// ToUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetToUserId() string {
    return r.toUserId
}
// DestUserId Setter
// 直调企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetDestUserId(destUserId string) error {
    r.destUserId = destUserId
    r.Set("dest_user_id", destUserId)
    return nil
}

// DestUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetDestUserId() string {
    return r.destUserId
}
// OperIcCode Setter
// 操作人标识（appkey编号）
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetOperIcCode() string {
    return r.operIcCode
}
// OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetOperIcName() string {
    return r.operIcName
}
// FileContent Setter
// 单据文件体【bas64字符串】，看对接文档中的代码示例，示例中有相应说明。
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetFileContent(fileContent string) error {
    r.fileContent = fileContent
    r.Set("file_content", fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetFileContent() string {
    return r.fileContent
}
// UploadFileName Setter
// 单据名称
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetUploadFileName(uploadFileName string) error {
    r.uploadFileName = uploadFileName
    r.Set("upload_file_name", uploadFileName)
    return nil
}

// UploadFileName Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetUploadFileName() string {
    return r.uploadFileName
}
// ClientType Setter
// 客户端类型【暂定都写2】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetClientType() string {
    return r.clientType
}
// QuReceivable Setter
// （协同平台数据合规）应收货总数量
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetQuReceivable(quReceivable int64) error {
    r.quReceivable = quReceivable
    r.Set("qu_receivable", quReceivable)
    return nil
}

// QuReceivable Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetQuReceivable() int64 {
    return r.quReceivable
}
// XtIsCheck Setter
// （协同平台数据合规）是否验证，0：未通过验证，1：已验证
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetXtIsCheck(xtIsCheck string) error {
    r.xtIsCheck = xtIsCheck
    r.Set("xt_is_check", xtIsCheck)
    return nil
}

// XtIsCheck Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetXtIsCheck() string {
    return r.xtIsCheck
}
// XtCheckCode Setter
// （协同平台数据合规）未验证通过原因【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetXtCheckCode(xtCheckCode string) error {
    r.xtCheckCode = xtCheckCode
    r.Set("xt_check_code", xtCheckCode)
    return nil
}

// XtCheckCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetXtCheckCode() string {
    return r.xtCheckCode
}
// XtCheckCodeDesc Setter
// （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetXtCheckCodeDesc(xtCheckCodeDesc string) error {
    r.xtCheckCodeDesc = xtCheckCodeDesc
    r.Set("xt_check_code_desc", xtCheckCodeDesc)
    return nil
}

// XtCheckCodeDesc Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetXtCheckCodeDesc() string {
    return r.xtCheckCodeDesc
}
// DrugListJson Setter
// （协同平台数据合规）药品列表Json："codeCount":         药品数量 "commDrugId":     国家药品唯一标识 "exprieDate":         生产日期 "physicInfo":          药品信息 "pkgSpec":           包状规格 "prepnCount":       制剂数量 "produceBatchNo":生产批次 "produceDate":      生产日期
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetDrugListJson(drugListJson string) error {
    r.drugListJson = drugListJson
    r.Set("drug_list_json", drugListJson)
    return nil
}

// DrugListJson Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetDrugListJson() string {
    return r.drugListJson
}
// AssRefEntId Setter
// （协同平台数据合规）单据委托企业refEntId
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetAssRefEntId(assRefEntId string) error {
    r.assRefEntId = assRefEntId
    r.Set("ass_ref_ent_id", assRefEntId)
    return nil
}

// AssRefEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetAssRefEntId() string {
    return r.assRefEntId
}
// DisEntId Setter
// （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetDisEntId(disEntId string) error {
    r.disEntId = disEntId
    r.Set("dis_ent_id", disEntId)
    return nil
}

// DisEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetDisEntId() string {
    return r.disEntId
}
// DisRefEntId Setter
// （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetDisRefEntId(disRefEntId string) error {
    r.disRefEntId = disRefEntId
    r.Set("dis_ref_ent_id", disRefEntId)
    return nil
}

// DisRefEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetDisRefEntId() string {
    return r.disRefEntId
}
// ToPerson Setter
// （协同平台数据合规）收货人【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetToPerson(toPerson string) error {
    r.toPerson = toPerson
    r.Set("to_person", toPerson)
    return nil
}

// ToPerson Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetToPerson() string {
    return r.toPerson
}
// FromPerson Setter
// （协同平台数据合规）发货人【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetFromPerson(fromPerson string) error {
    r.fromPerson = fromPerson
    r.Set("from_person", fromPerson)
    return nil
}

// FromPerson Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetFromPerson() string {
    return r.fromPerson
}
// OrderCode Setter
// （协同平台数据合规）订货单编号【可选】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetOrderCode() string {
    return r.orderCode
}
// FromBillCode Setter
// （协同平台数据合规）发货单编号【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetFromBillCode(fromBillCode string) error {
    r.fromBillCode = fromBillCode
    r.Set("from_bill_code", fromBillCode)
    return nil
}

// FromBillCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetFromBillCode() string {
    return r.fromBillCode
}
// ToAddress Setter
// （协同平台数据合规）收货地址【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetToAddress(toAddress string) error {
    r.toAddress = toAddress
    r.Set("to_address", toAddress)
    return nil
}

// ToAddress Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetToAddress() string {
    return r.toAddress
}
// FromAddress Setter
// （协同平台数据合规）发货地址【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetFromAddress(fromAddress string) error {
    r.fromAddress = fromAddress
    r.Set("from_address", fromAddress)
    return nil
}

// FromAddress Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetFromAddress() string {
    return r.fromAddress
}
// AssEntId Setter
// （协同平台数据合规）单据委托企业entId
func (r *AlibabaAlihealthDrugKytUploadcircubillRequest) SetAssEntId(assEntId string) error {
    r.assEntId = assEntId
    r.Set("ass_ent_id", assEntId)
    return nil
}

// AssEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillRequest) GetAssEntId() string {
    return r.assEntId
}
