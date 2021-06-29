package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出入库单据上传 API请求
alibaba.alihealth.drugtrace.top.lsyd.uploadinoutbill

零售企业上传出入库信息，包括101, "生产入库"；102, "采购入库"；103, "退货入库"；104, "调拨入库"；106, "零头入库"；107, "供应入库"；108, "召回入库"；110,"赠品入库"；111,"盘盈入库"；112,"报废入库"；113,"其他入库"
201, "销售出库"；202, "退货出库"；203, "调拨出库"；204, "返工出库"；205, "销毁出库"；206, "抽检出库"；207, "直调出库"；208, "生产出库"；209, "供应出库"；211, "召回出库"；212,"赠品出库"；214,"盘亏出库"；215,"损坏出库"；216,"报废出库"；217,"其他出库"；237, "直调退货"。
不包括对个人的零售出库，疫苗接种，领药出库。
*/
type AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest struct {
    model.Params
    // 单据编码（每个单据号只能上传一次）
    _billCode   string
    // 单据时间
    _billTime   string
    // 单据类型【102代表采购入库】
    _billType   int64
    // 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药
    _physicType   int64
    // 上传单据企业的单位编码【注意：该入参是ref_ent_id，不是ent_id】
    _refUserId   string
    // 代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
    _agentRefUserId   string
    // 发货企业entId
    _fromUserId   string
    // 收货企业entId
    _toUserId   string
    // 直调企业标识
    _destUserId   string
    // 单据提交者（appkey编号）
    _operIcCode   string
    // 单据提交者姓名
    _operIcName   string
    // 客户端类型[必须填2]
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
    // （协同平台数据合规）发货地址【必选】
    _fromAddress   string
    // （协同平台数据合规）收货地址【必选】
    _toAddress   string
    // （协同平台数据合规）发货单编号【必选】
    _fromBillCode   string
    // （协同平台数据合规）订货单编号【可选】
    _orderCode   string
    // （协同平台数据合规）发货人【必选】
    _fromPerson   string
    // （协同平台数据合规）收货人【必选】
    _toPerson   string
    // （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
    _disRefEntId   string
    // （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
    _disEntId   string
    // （协同平台数据合规）应收货总数量【必选】
    _quReceivable   int64
    // （协同平台数据合规）是否验证，0：未通过验证，1：已验证
    _xtIsCheck   string
    // （协同平台数据合规）未验证通过原因【验证未通过时填写】
    _xtCheckCode   string
    // （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
    _xtCheckCodeDesc   string
    // （协同平台数据合规）药品列表Json："codeCount": 药品数量 "commDrugId": 国家药品唯一标识 "exprieDate": 生产日期 "physicInfo": 药品信息 "pkgSpec": 包状规格 "prepnCount": 制剂数量 "produceBatchNo":生产批次 "produceDate": 生产日期
    _drugListJson   string
    // （协同平台数据合规）单据委托企业refEntId【疫苗药品出库单填写】
    _assRefEntId   string
    // （协同平台数据合规）单据委托企业entId【疫苗药品出库单填写】
    _assEntId   string
}

// 初始化AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest() *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest{
    return &AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.lsyd.uploadinoutbill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编码（每个单据号只能上传一次）
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetBillCode() string {
    return r._billCode
}
// BillTime Setter
// 单据时间
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetBillTime(_billTime string) error {
    r._billTime = _billTime
    r.Set("bill_time", _billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetBillTime() string {
    return r._billTime
}
// BillType Setter
// 单据类型【102代表采购入库】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetBillType(_billType int64) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetBillType() int64 {
    return r._billType
}
// PhysicType Setter
// 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetPhysicType(_physicType int64) error {
    r._physicType = _physicType
    r.Set("physic_type", _physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetPhysicType() int64 {
    return r._physicType
}
// RefUserId Setter
// 上传单据企业的单位编码【注意：该入参是ref_ent_id，不是ent_id】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetRefUserId(_refUserId string) error {
    r._refUserId = _refUserId
    r.Set("ref_user_id", _refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetRefUserId() string {
    return r._refUserId
}
// AgentRefUserId Setter
// 代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetAgentRefUserId(_agentRefUserId string) error {
    r._agentRefUserId = _agentRefUserId
    r.Set("agent_ref_user_id", _agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetAgentRefUserId() string {
    return r._agentRefUserId
}
// FromUserId Setter
// 发货企业entId
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetFromUserId(_fromUserId string) error {
    r._fromUserId = _fromUserId
    r.Set("from_user_id", _fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetFromUserId() string {
    return r._fromUserId
}
// ToUserId Setter
// 收货企业entId
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetToUserId(_toUserId string) error {
    r._toUserId = _toUserId
    r.Set("to_user_id", _toUserId)
    return nil
}

// ToUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetToUserId() string {
    return r._toUserId
}
// DestUserId Setter
// 直调企业标识
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetDestUserId(_destUserId string) error {
    r._destUserId = _destUserId
    r.Set("dest_user_id", _destUserId)
    return nil
}

// DestUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetDestUserId() string {
    return r._destUserId
}
// OperIcCode Setter
// 单据提交者（appkey编号）
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetOperIcCode(_operIcCode string) error {
    r._operIcCode = _operIcCode
    r.Set("oper_ic_code", _operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetOperIcCode() string {
    return r._operIcCode
}
// OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetOperIcName(_operIcName string) error {
    r._operIcName = _operIcName
    r.Set("oper_ic_name", _operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetOperIcName() string {
    return r._operIcName
}
// ClientType Setter
// 客户端类型[必须填2]
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetClientType() string {
    return r._clientType
}
// ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetReturnReasonCode(_returnReasonCode string) error {
    r._returnReasonCode = _returnReasonCode
    r.Set("return_reason_code", _returnReasonCode)
    return nil
}

// ReturnReasonCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetReturnReasonCode() string {
    return r._returnReasonCode
}
// ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetReturnReasonDes(_returnReasonDes string) error {
    r._returnReasonDes = _returnReasonDes
    r.Set("return_reason_des", _returnReasonDes)
    return nil
}

// ReturnReasonDes Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetReturnReasonDes() string {
    return r._returnReasonDes
}
// CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetCancelReasonCode(_cancelReasonCode string) error {
    r._cancelReasonCode = _cancelReasonCode
    r.Set("cancel_reason_code", _cancelReasonCode)
    return nil
}

// CancelReasonCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetCancelReasonCode() string {
    return r._cancelReasonCode
}
// CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetCancelReasonDes(_cancelReasonDes string) error {
    r._cancelReasonDes = _cancelReasonDes
    r.Set("cancel_reason_des", _cancelReasonDes)
    return nil
}

// CancelReasonDes Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetCancelReasonDes() string {
    return r._cancelReasonDes
}
// ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetExecuterName(_executerName string) error {
    r._executerName = _executerName
    r.Set("executer_name", _executerName)
    return nil
}

// ExecuterName Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetExecuterName() string {
    return r._executerName
}
// ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetExecuterCode(_executerCode string) error {
    r._executerCode = _executerCode
    r.Set("executer_code", _executerCode)
    return nil
}

// ExecuterCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetExecuterCode() string {
    return r._executerCode
}
// SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetSuperviserName(_superviserName string) error {
    r._superviserName = _superviserName
    r.Set("superviser_name", _superviserName)
    return nil
}

// SuperviserName Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetSuperviserName() string {
    return r._superviserName
}
// SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetSuperviserCode(_superviserCode string) error {
    r._superviserCode = _superviserCode
    r.Set("superviser_code", _superviserCode)
    return nil
}

// SuperviserCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetSuperviserCode() string {
    return r._superviserCode
}
// WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetWarehouseId(_warehouseId string) error {
    r._warehouseId = _warehouseId
    r.Set("warehouse_id", _warehouseId)
    return nil
}

// WarehouseId Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetWarehouseId() string {
    return r._warehouseId
}
// DrugId Setter
// 药品ID[企业自已系统的药品ID]
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetDrugId(_drugId string) error {
    r._drugId = _drugId
    r.Set("drug_id", _drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetDrugId() string {
    return r._drugId
}
// TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetTraceCodes(_traceCodes []string) error {
    r._traceCodes = _traceCodes
    r.Set("trace_codes", _traceCodes)
    return nil
}

// TraceCodes Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetTraceCodes() []string {
    return r._traceCodes
}
// FromAddress Setter
// （协同平台数据合规）发货地址【必选】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetFromAddress(_fromAddress string) error {
    r._fromAddress = _fromAddress
    r.Set("from_address", _fromAddress)
    return nil
}

// FromAddress Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetFromAddress() string {
    return r._fromAddress
}
// ToAddress Setter
// （协同平台数据合规）收货地址【必选】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetToAddress(_toAddress string) error {
    r._toAddress = _toAddress
    r.Set("to_address", _toAddress)
    return nil
}

// ToAddress Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetToAddress() string {
    return r._toAddress
}
// FromBillCode Setter
// （协同平台数据合规）发货单编号【必选】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetFromBillCode(_fromBillCode string) error {
    r._fromBillCode = _fromBillCode
    r.Set("from_bill_code", _fromBillCode)
    return nil
}

// FromBillCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetFromBillCode() string {
    return r._fromBillCode
}
// OrderCode Setter
// （协同平台数据合规）订货单编号【可选】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetOrderCode() string {
    return r._orderCode
}
// FromPerson Setter
// （协同平台数据合规）发货人【必选】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetFromPerson(_fromPerson string) error {
    r._fromPerson = _fromPerson
    r.Set("from_person", _fromPerson)
    return nil
}

// FromPerson Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetFromPerson() string {
    return r._fromPerson
}
// ToPerson Setter
// （协同平台数据合规）收货人【必选】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetToPerson(_toPerson string) error {
    r._toPerson = _toPerson
    r.Set("to_person", _toPerson)
    return nil
}

// ToPerson Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetToPerson() string {
    return r._toPerson
}
// DisRefEntId Setter
// （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetDisRefEntId(_disRefEntId string) error {
    r._disRefEntId = _disRefEntId
    r.Set("dis_ref_ent_id", _disRefEntId)
    return nil
}

// DisRefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetDisRefEntId() string {
    return r._disRefEntId
}
// DisEntId Setter
// （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetDisEntId(_disEntId string) error {
    r._disEntId = _disEntId
    r.Set("dis_ent_id", _disEntId)
    return nil
}

// DisEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetDisEntId() string {
    return r._disEntId
}
// QuReceivable Setter
// （协同平台数据合规）应收货总数量【必选】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetQuReceivable(_quReceivable int64) error {
    r._quReceivable = _quReceivable
    r.Set("qu_receivable", _quReceivable)
    return nil
}

// QuReceivable Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetQuReceivable() int64 {
    return r._quReceivable
}
// XtIsCheck Setter
// （协同平台数据合规）是否验证，0：未通过验证，1：已验证
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetXtIsCheck(_xtIsCheck string) error {
    r._xtIsCheck = _xtIsCheck
    r.Set("xt_is_check", _xtIsCheck)
    return nil
}

// XtIsCheck Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetXtIsCheck() string {
    return r._xtIsCheck
}
// XtCheckCode Setter
// （协同平台数据合规）未验证通过原因【验证未通过时填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetXtCheckCode(_xtCheckCode string) error {
    r._xtCheckCode = _xtCheckCode
    r.Set("xt_check_code", _xtCheckCode)
    return nil
}

// XtCheckCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetXtCheckCode() string {
    return r._xtCheckCode
}
// XtCheckCodeDesc Setter
// （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetXtCheckCodeDesc(_xtCheckCodeDesc string) error {
    r._xtCheckCodeDesc = _xtCheckCodeDesc
    r.Set("xt_check_code_desc", _xtCheckCodeDesc)
    return nil
}

// XtCheckCodeDesc Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetXtCheckCodeDesc() string {
    return r._xtCheckCodeDesc
}
// DrugListJson Setter
// （协同平台数据合规）药品列表Json："codeCount": 药品数量 "commDrugId": 国家药品唯一标识 "exprieDate": 生产日期 "physicInfo": 药品信息 "pkgSpec": 包状规格 "prepnCount": 制剂数量 "produceBatchNo":生产批次 "produceDate": 生产日期
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetDrugListJson(_drugListJson string) error {
    r._drugListJson = _drugListJson
    r.Set("drug_list_json", _drugListJson)
    return nil
}

// DrugListJson Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetDrugListJson() string {
    return r._drugListJson
}
// AssRefEntId Setter
// （协同平台数据合规）单据委托企业refEntId【疫苗药品出库单填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetAssRefEntId(_assRefEntId string) error {
    r._assRefEntId = _assRefEntId
    r.Set("ass_ref_ent_id", _assRefEntId)
    return nil
}

// AssRefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetAssRefEntId() string {
    return r._assRefEntId
}
// AssEntId Setter
// （协同平台数据合规）单据委托企业entId【疫苗药品出库单填写】
func (r *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) SetAssEntId(_assEntId string) error {
    r._assEntId = _assEntId
    r.Set("ass_ent_id", _assEntId)
    return nil
}

// AssEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydUploadinoutbillRequest) GetAssEntId() string {
    return r._assEntId
}