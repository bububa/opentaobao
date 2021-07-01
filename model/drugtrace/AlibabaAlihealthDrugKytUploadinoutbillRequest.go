package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
企业上传出入库信息 API请求
alibaba.alihealth.drug.kyt.uploadinoutbill

企业上传出入库信息，包括101, "生产入库"；102, "采购入库"；103, "退货入库"；104, "调拨入库"；106, "零头入库"；107, "供应入库"；108, "召回入库"；110,"赠品入库"；111,"盘盈入库"；112,"报废入库"；113,"其他入库"
201, "销售出库"；202, "退货出库"；203, "调拨出库"；204, "返工出库"；205, "销毁出库"；206, "抽检出库"；207, "直调出库"；208, "生产出库"；209, "供应出库"；211, "召回出库"；212,"赠品出库"；214,"盘亏出库"；215,"损坏出库"；216,"报废出库"；217,"其他出库"；237, "直调退货"。
不包括对个人的零售出库，疫苗接种，领药出库。
本接口与uploadcircubill接口的主要区别的，本接口入参中直接上传追溯码（多个码时用逗号分隔）。uploadcircubill接口入参中，需要上传码的单据文件（用扫码枪生成的xml文件），一般情况下使用uploadcircubill接口上传单据文件。
*/
type AlibabaAlihealthDrugKytUploadinoutbillAPIRequest struct {
    model.Params
    // 单据编号【同一个企业不能上传相同单据号】
    _billCode   string
    // 单据时间（扫码时间）
    _billTime   string
    // 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
    _billType   int64
    // 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药
    _physicType   int64
    // 货主（单据的所有者，上传人），上传企业的单位维一编码【注意：该入参是ref_ent_id，不是ent_id】
    _refUserId   string
    // 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
    _agentRefUserId   string
    // 发货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
    _fromUserId   string
    // 收货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
    _toUserId   string
    // 直调企业标识【注意：该入参是ent_id,并不是ref_ent_id】
    _destUserId   string
    // 单据提交者（调用接口时的appkey编号）
    _operIcCode   string
    // 单据提交者姓名（出入库单上传人的名子）
    _operIcName   string
    // 调用方类型[必须填2]
    _clientType   string
    // 退货原因代码[退货入出库时填写]
    _returnReasonCode   string
    // 退货原因描述[退货入出库时填写]
    _returnReasonDes   string
    // 注销原因代码【销毁出库时填写】
    _cancelReasonCode   string
    // 注销原因描述【销毁出库时填写】
    _cancelReasonDes   string
    // 执行人姓名【销毁出库时填写】
    _executerName   string
    // 执行人证件号【销毁出库时填写】
    _executerCode   string
    // 监督人姓名【销毁出库时填写】
    _superviserName   string
    // 监督人证件号【销毁出库时填写】
    _superviserCode   string
    // 仓号
    _warehouseId   string
    // 药品ID[企业自已系统的药品ID]
    _drugId   string
    // 追溯码[多个时用逗号分开]
    _traceCodes   []string
    // （协同平台数据合规）应收货总数量
    _quReceivable   int64
    // （协同平台数据合规）是否验证，0：未通过验证，1：已验证
    _xtIsCheck   string
    // （协同平台数据合规）未验证通过原因【验证未通过时填写】
    _xtCheckCode   string
    // （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
    _xtCheckCodeDesc   string
    // （协同平台数据合规）药品列表Json："codeCount": 药品数量 "commDrugId": 国家药品唯一标识 "exprieDate": 生产日期 "physicInfo": 药品信息 "pkgSpec": 包状规格 "prepnCount": 制剂数量 "produceBatchNo":生产批次 "produceDate": 生产日期
    _drugListJson   string
    // （协同平台数据合规）单据委托企业refEntId
    _assRefEntId   string
    // （协同平台数据合规）药品配送企业entId出库单，收货方为医疗机构时填写】
    _disEntId   string
    // （协同平台数据合规）药品配送企业出库单，收货方为医疗机构时填写】
    _disRefEntId   string
    // （协同平台数据合规）收货人【必选】
    _toPerson   string
    // （协同平台数据合规）发货人【必选】
    _fromPerson   string
    // （协同平台数据合规）订货单编号【可选】
    _orderCode   string
    // （协同平台数据合规）发货单编号【必选】
    _fromBillCode   string
    // （协同平台数据合规）收货地址【必选】
    _toAddress   string
    // （协同平台数据合规）发货地址【必选】
    _fromAddress   string
    // （协同平台数据合规）单据委托企业entId
    _assEntId   string
}

// 初始化AlibabaAlihealthDrugKytUploadinoutbillAPIRequest对象
func NewAlibabaAlihealthDrugKytUploadinoutbillRequest() *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest{
    return &AlibabaAlihealthDrugKytUploadinoutbillAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.uploadinoutbill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编号【同一个企业不能上传相同单据号】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetBillCode() string {
    return r._billCode
}
// BillTime Setter
// 单据时间（扫码时间）
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetBillTime(_billTime string) error {
    r._billTime = _billTime
    r.Set("bill_time", _billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetBillTime() string {
    return r._billTime
}
// BillType Setter
// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetBillType(_billType int64) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetBillType() int64 {
    return r._billType
}
// PhysicType Setter
// 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetPhysicType(_physicType int64) error {
    r._physicType = _physicType
    r.Set("physic_type", _physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetPhysicType() int64 {
    return r._physicType
}
// RefUserId Setter
// 货主（单据的所有者，上传人），上传企业的单位维一编码【注意：该入参是ref_ent_id，不是ent_id】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetRefUserId(_refUserId string) error {
    r._refUserId = _refUserId
    r.Set("ref_user_id", _refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetRefUserId() string {
    return r._refUserId
}
// AgentRefUserId Setter
// 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
    r._agentRefUserId = _agentRefUserId
    r.Set("agent_ref_user_id", _agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetAgentRefUserId() string {
    return r._agentRefUserId
}
// FromUserId Setter
// 发货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetFromUserId(_fromUserId string) error {
    r._fromUserId = _fromUserId
    r.Set("from_user_id", _fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetFromUserId() string {
    return r._fromUserId
}
// ToUserId Setter
// 收货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetToUserId(_toUserId string) error {
    r._toUserId = _toUserId
    r.Set("to_user_id", _toUserId)
    return nil
}

// ToUserId Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetToUserId() string {
    return r._toUserId
}
// DestUserId Setter
// 直调企业标识【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetDestUserId(_destUserId string) error {
    r._destUserId = _destUserId
    r.Set("dest_user_id", _destUserId)
    return nil
}

// DestUserId Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetDestUserId() string {
    return r._destUserId
}
// OperIcCode Setter
// 单据提交者（调用接口时的appkey编号）
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetOperIcCode(_operIcCode string) error {
    r._operIcCode = _operIcCode
    r.Set("oper_ic_code", _operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetOperIcCode() string {
    return r._operIcCode
}
// OperIcName Setter
// 单据提交者姓名（出入库单上传人的名子）
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetOperIcName(_operIcName string) error {
    r._operIcName = _operIcName
    r.Set("oper_ic_name", _operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetOperIcName() string {
    return r._operIcName
}
// ClientType Setter
// 调用方类型[必须填2]
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetClientType() string {
    return r._clientType
}
// ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetReturnReasonCode(_returnReasonCode string) error {
    r._returnReasonCode = _returnReasonCode
    r.Set("return_reason_code", _returnReasonCode)
    return nil
}

// ReturnReasonCode Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetReturnReasonCode() string {
    return r._returnReasonCode
}
// ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetReturnReasonDes(_returnReasonDes string) error {
    r._returnReasonDes = _returnReasonDes
    r.Set("return_reason_des", _returnReasonDes)
    return nil
}

// ReturnReasonDes Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetReturnReasonDes() string {
    return r._returnReasonDes
}
// CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetCancelReasonCode(_cancelReasonCode string) error {
    r._cancelReasonCode = _cancelReasonCode
    r.Set("cancel_reason_code", _cancelReasonCode)
    return nil
}

// CancelReasonCode Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetCancelReasonCode() string {
    return r._cancelReasonCode
}
// CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetCancelReasonDes(_cancelReasonDes string) error {
    r._cancelReasonDes = _cancelReasonDes
    r.Set("cancel_reason_des", _cancelReasonDes)
    return nil
}

// CancelReasonDes Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetCancelReasonDes() string {
    return r._cancelReasonDes
}
// ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetExecuterName(_executerName string) error {
    r._executerName = _executerName
    r.Set("executer_name", _executerName)
    return nil
}

// ExecuterName Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetExecuterName() string {
    return r._executerName
}
// ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetExecuterCode(_executerCode string) error {
    r._executerCode = _executerCode
    r.Set("executer_code", _executerCode)
    return nil
}

// ExecuterCode Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetExecuterCode() string {
    return r._executerCode
}
// SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetSuperviserName(_superviserName string) error {
    r._superviserName = _superviserName
    r.Set("superviser_name", _superviserName)
    return nil
}

// SuperviserName Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetSuperviserName() string {
    return r._superviserName
}
// SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetSuperviserCode(_superviserCode string) error {
    r._superviserCode = _superviserCode
    r.Set("superviser_code", _superviserCode)
    return nil
}

// SuperviserCode Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetSuperviserCode() string {
    return r._superviserCode
}
// WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetWarehouseId(_warehouseId string) error {
    r._warehouseId = _warehouseId
    r.Set("warehouse_id", _warehouseId)
    return nil
}

// WarehouseId Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetWarehouseId() string {
    return r._warehouseId
}
// DrugId Setter
// 药品ID[企业自已系统的药品ID]
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetDrugId(_drugId string) error {
    r._drugId = _drugId
    r.Set("drug_id", _drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetDrugId() string {
    return r._drugId
}
// TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetTraceCodes(_traceCodes []string) error {
    r._traceCodes = _traceCodes
    r.Set("trace_codes", _traceCodes)
    return nil
}

// TraceCodes Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetTraceCodes() []string {
    return r._traceCodes
}
// QuReceivable Setter
// （协同平台数据合规）应收货总数量
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetQuReceivable(_quReceivable int64) error {
    r._quReceivable = _quReceivable
    r.Set("qu_receivable", _quReceivable)
    return nil
}

// QuReceivable Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetQuReceivable() int64 {
    return r._quReceivable
}
// XtIsCheck Setter
// （协同平台数据合规）是否验证，0：未通过验证，1：已验证
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetXtIsCheck(_xtIsCheck string) error {
    r._xtIsCheck = _xtIsCheck
    r.Set("xt_is_check", _xtIsCheck)
    return nil
}

// XtIsCheck Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetXtIsCheck() string {
    return r._xtIsCheck
}
// XtCheckCode Setter
// （协同平台数据合规）未验证通过原因【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetXtCheckCode(_xtCheckCode string) error {
    r._xtCheckCode = _xtCheckCode
    r.Set("xt_check_code", _xtCheckCode)
    return nil
}

// XtCheckCode Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetXtCheckCode() string {
    return r._xtCheckCode
}
// XtCheckCodeDesc Setter
// （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetXtCheckCodeDesc(_xtCheckCodeDesc string) error {
    r._xtCheckCodeDesc = _xtCheckCodeDesc
    r.Set("xt_check_code_desc", _xtCheckCodeDesc)
    return nil
}

// XtCheckCodeDesc Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetXtCheckCodeDesc() string {
    return r._xtCheckCodeDesc
}
// DrugListJson Setter
// （协同平台数据合规）药品列表Json："codeCount": 药品数量 "commDrugId": 国家药品唯一标识 "exprieDate": 生产日期 "physicInfo": 药品信息 "pkgSpec": 包状规格 "prepnCount": 制剂数量 "produceBatchNo":生产批次 "produceDate": 生产日期
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetDrugListJson(_drugListJson string) error {
    r._drugListJson = _drugListJson
    r.Set("drug_list_json", _drugListJson)
    return nil
}

// DrugListJson Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetDrugListJson() string {
    return r._drugListJson
}
// AssRefEntId Setter
// （协同平台数据合规）单据委托企业refEntId
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetAssRefEntId(_assRefEntId string) error {
    r._assRefEntId = _assRefEntId
    r.Set("ass_ref_ent_id", _assRefEntId)
    return nil
}

// AssRefEntId Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetAssRefEntId() string {
    return r._assRefEntId
}
// DisEntId Setter
// （协同平台数据合规）药品配送企业entId出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetDisEntId(_disEntId string) error {
    r._disEntId = _disEntId
    r.Set("dis_ent_id", _disEntId)
    return nil
}

// DisEntId Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetDisEntId() string {
    return r._disEntId
}
// DisRefEntId Setter
// （协同平台数据合规）药品配送企业出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetDisRefEntId(_disRefEntId string) error {
    r._disRefEntId = _disRefEntId
    r.Set("dis_ref_ent_id", _disRefEntId)
    return nil
}

// DisRefEntId Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetDisRefEntId() string {
    return r._disRefEntId
}
// ToPerson Setter
// （协同平台数据合规）收货人【必选】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetToPerson(_toPerson string) error {
    r._toPerson = _toPerson
    r.Set("to_person", _toPerson)
    return nil
}

// ToPerson Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetToPerson() string {
    return r._toPerson
}
// FromPerson Setter
// （协同平台数据合规）发货人【必选】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetFromPerson(_fromPerson string) error {
    r._fromPerson = _fromPerson
    r.Set("from_person", _fromPerson)
    return nil
}

// FromPerson Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetFromPerson() string {
    return r._fromPerson
}
// OrderCode Setter
// （协同平台数据合规）订货单编号【可选】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetOrderCode() string {
    return r._orderCode
}
// FromBillCode Setter
// （协同平台数据合规）发货单编号【必选】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetFromBillCode(_fromBillCode string) error {
    r._fromBillCode = _fromBillCode
    r.Set("from_bill_code", _fromBillCode)
    return nil
}

// FromBillCode Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetFromBillCode() string {
    return r._fromBillCode
}
// ToAddress Setter
// （协同平台数据合规）收货地址【必选】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetToAddress(_toAddress string) error {
    r._toAddress = _toAddress
    r.Set("to_address", _toAddress)
    return nil
}

// ToAddress Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetToAddress() string {
    return r._toAddress
}
// FromAddress Setter
// （协同平台数据合规）发货地址【必选】
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetFromAddress(_fromAddress string) error {
    r._fromAddress = _fromAddress
    r.Set("from_address", _fromAddress)
    return nil
}

// FromAddress Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetFromAddress() string {
    return r._fromAddress
}
// AssEntId Setter
// （协同平台数据合规）单据委托企业entId
func (r *AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) SetAssEntId(_assEntId string) error {
    r._assEntId = _assEntId
    r.Set("ass_ent_id", _assEntId)
    return nil
}

// AssEntId Getter
func (r AlibabaAlihealthDrugKytUploadinoutbillAPIRequest) GetAssEntId() string {
    return r._assEntId
}
