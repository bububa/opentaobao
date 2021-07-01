package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗企业出入库上传 API请求
alibaba.alihealth.drug.kyt.dr.uploadinoutbill

零售企业上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）, 
不包括对个人的零售出库，疫苗接种，领药出库。
*/
type AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest struct {
    model.Params
    // 单据编码
    _billCode   string
    // 单据时间
    _billTime   string
    // 单据类型【102代表采购入库】
    _billType   int64
    // 药品类型【3普药2特药】
    _physicType   int64
    // 上传企业的单位编码
    _refUserId   string
    // 代理企业REF标识
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
    // （协同平台数据合规）药品列表Json
    _drugListJson   string
    // （协同平台数据合规）单据委托企业refEntId【疫苗药品出库单填写】
    _assRefEntId   string
    // （协同平台数据合规）单据委托企业entId【疫苗药品出库单填写】
    _assEntId   string
}

// 初始化AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest对象
func NewAlibabaAlihealthDrugKytDrUploadinoutbillRequest() *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest{
    return &AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.uploadinoutbill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetBillCode() string {
    return r._billCode
}
// BillTime Setter
// 单据时间
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetBillTime(_billTime string) error {
    r._billTime = _billTime
    r.Set("bill_time", _billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetBillTime() string {
    return r._billTime
}
// BillType Setter
// 单据类型【102代表采购入库】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetBillType(_billType int64) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetBillType() int64 {
    return r._billType
}
// PhysicType Setter
// 药品类型【3普药2特药】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetPhysicType(_physicType int64) error {
    r._physicType = _physicType
    r.Set("physic_type", _physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetPhysicType() int64 {
    return r._physicType
}
// RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetRefUserId(_refUserId string) error {
    r._refUserId = _refUserId
    r.Set("ref_user_id", _refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetRefUserId() string {
    return r._refUserId
}
// AgentRefUserId Setter
// 代理企业REF标识
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
    r._agentRefUserId = _agentRefUserId
    r.Set("agent_ref_user_id", _agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetAgentRefUserId() string {
    return r._agentRefUserId
}
// FromUserId Setter
// 发货企业entId
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetFromUserId(_fromUserId string) error {
    r._fromUserId = _fromUserId
    r.Set("from_user_id", _fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetFromUserId() string {
    return r._fromUserId
}
// ToUserId Setter
// 收货企业entId
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetToUserId(_toUserId string) error {
    r._toUserId = _toUserId
    r.Set("to_user_id", _toUserId)
    return nil
}

// ToUserId Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetToUserId() string {
    return r._toUserId
}
// DestUserId Setter
// 直调企业标识
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetDestUserId(_destUserId string) error {
    r._destUserId = _destUserId
    r.Set("dest_user_id", _destUserId)
    return nil
}

// DestUserId Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetDestUserId() string {
    return r._destUserId
}
// OperIcCode Setter
// 单据提交者（appkey编号）
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetOperIcCode(_operIcCode string) error {
    r._operIcCode = _operIcCode
    r.Set("oper_ic_code", _operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetOperIcCode() string {
    return r._operIcCode
}
// OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetOperIcName(_operIcName string) error {
    r._operIcName = _operIcName
    r.Set("oper_ic_name", _operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetOperIcName() string {
    return r._operIcName
}
// ClientType Setter
// 客户端类型[必须填2]
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetClientType() string {
    return r._clientType
}
// ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetReturnReasonCode(_returnReasonCode string) error {
    r._returnReasonCode = _returnReasonCode
    r.Set("return_reason_code", _returnReasonCode)
    return nil
}

// ReturnReasonCode Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetReturnReasonCode() string {
    return r._returnReasonCode
}
// ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetReturnReasonDes(_returnReasonDes string) error {
    r._returnReasonDes = _returnReasonDes
    r.Set("return_reason_des", _returnReasonDes)
    return nil
}

// ReturnReasonDes Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetReturnReasonDes() string {
    return r._returnReasonDes
}
// CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetCancelReasonCode(_cancelReasonCode string) error {
    r._cancelReasonCode = _cancelReasonCode
    r.Set("cancel_reason_code", _cancelReasonCode)
    return nil
}

// CancelReasonCode Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetCancelReasonCode() string {
    return r._cancelReasonCode
}
// CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetCancelReasonDes(_cancelReasonDes string) error {
    r._cancelReasonDes = _cancelReasonDes
    r.Set("cancel_reason_des", _cancelReasonDes)
    return nil
}

// CancelReasonDes Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetCancelReasonDes() string {
    return r._cancelReasonDes
}
// ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetExecuterName(_executerName string) error {
    r._executerName = _executerName
    r.Set("executer_name", _executerName)
    return nil
}

// ExecuterName Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetExecuterName() string {
    return r._executerName
}
// ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetExecuterCode(_executerCode string) error {
    r._executerCode = _executerCode
    r.Set("executer_code", _executerCode)
    return nil
}

// ExecuterCode Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetExecuterCode() string {
    return r._executerCode
}
// SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetSuperviserName(_superviserName string) error {
    r._superviserName = _superviserName
    r.Set("superviser_name", _superviserName)
    return nil
}

// SuperviserName Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetSuperviserName() string {
    return r._superviserName
}
// SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetSuperviserCode(_superviserCode string) error {
    r._superviserCode = _superviserCode
    r.Set("superviser_code", _superviserCode)
    return nil
}

// SuperviserCode Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetSuperviserCode() string {
    return r._superviserCode
}
// WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetWarehouseId(_warehouseId string) error {
    r._warehouseId = _warehouseId
    r.Set("warehouse_id", _warehouseId)
    return nil
}

// WarehouseId Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetWarehouseId() string {
    return r._warehouseId
}
// DrugId Setter
// 药品ID[企业自已系统的药品ID]
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetDrugId(_drugId string) error {
    r._drugId = _drugId
    r.Set("drug_id", _drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetDrugId() string {
    return r._drugId
}
// TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetTraceCodes(_traceCodes []string) error {
    r._traceCodes = _traceCodes
    r.Set("trace_codes", _traceCodes)
    return nil
}

// TraceCodes Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetTraceCodes() []string {
    return r._traceCodes
}
// FromAddress Setter
// （协同平台数据合规）发货地址【必选】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetFromAddress(_fromAddress string) error {
    r._fromAddress = _fromAddress
    r.Set("from_address", _fromAddress)
    return nil
}

// FromAddress Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetFromAddress() string {
    return r._fromAddress
}
// ToAddress Setter
// （协同平台数据合规）收货地址【必选】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetToAddress(_toAddress string) error {
    r._toAddress = _toAddress
    r.Set("to_address", _toAddress)
    return nil
}

// ToAddress Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetToAddress() string {
    return r._toAddress
}
// FromBillCode Setter
// （协同平台数据合规）发货单编号【必选】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetFromBillCode(_fromBillCode string) error {
    r._fromBillCode = _fromBillCode
    r.Set("from_bill_code", _fromBillCode)
    return nil
}

// FromBillCode Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetFromBillCode() string {
    return r._fromBillCode
}
// OrderCode Setter
// （协同平台数据合规）订货单编号【可选】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetOrderCode() string {
    return r._orderCode
}
// FromPerson Setter
// （协同平台数据合规）发货人【必选】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetFromPerson(_fromPerson string) error {
    r._fromPerson = _fromPerson
    r.Set("from_person", _fromPerson)
    return nil
}

// FromPerson Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetFromPerson() string {
    return r._fromPerson
}
// ToPerson Setter
// （协同平台数据合规）收货人【必选】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetToPerson(_toPerson string) error {
    r._toPerson = _toPerson
    r.Set("to_person", _toPerson)
    return nil
}

// ToPerson Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetToPerson() string {
    return r._toPerson
}
// DisRefEntId Setter
// （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetDisRefEntId(_disRefEntId string) error {
    r._disRefEntId = _disRefEntId
    r.Set("dis_ref_ent_id", _disRefEntId)
    return nil
}

// DisRefEntId Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetDisRefEntId() string {
    return r._disRefEntId
}
// DisEntId Setter
// （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetDisEntId(_disEntId string) error {
    r._disEntId = _disEntId
    r.Set("dis_ent_id", _disEntId)
    return nil
}

// DisEntId Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetDisEntId() string {
    return r._disEntId
}
// QuReceivable Setter
// （协同平台数据合规）应收货总数量【必选】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetQuReceivable(_quReceivable int64) error {
    r._quReceivable = _quReceivable
    r.Set("qu_receivable", _quReceivable)
    return nil
}

// QuReceivable Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetQuReceivable() int64 {
    return r._quReceivable
}
// XtIsCheck Setter
// （协同平台数据合规）是否验证，0：未通过验证，1：已验证
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetXtIsCheck(_xtIsCheck string) error {
    r._xtIsCheck = _xtIsCheck
    r.Set("xt_is_check", _xtIsCheck)
    return nil
}

// XtIsCheck Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetXtIsCheck() string {
    return r._xtIsCheck
}
// XtCheckCode Setter
// （协同平台数据合规）未验证通过原因【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetXtCheckCode(_xtCheckCode string) error {
    r._xtCheckCode = _xtCheckCode
    r.Set("xt_check_code", _xtCheckCode)
    return nil
}

// XtCheckCode Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetXtCheckCode() string {
    return r._xtCheckCode
}
// XtCheckCodeDesc Setter
// （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetXtCheckCodeDesc(_xtCheckCodeDesc string) error {
    r._xtCheckCodeDesc = _xtCheckCodeDesc
    r.Set("xt_check_code_desc", _xtCheckCodeDesc)
    return nil
}

// XtCheckCodeDesc Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetXtCheckCodeDesc() string {
    return r._xtCheckCodeDesc
}
// DrugListJson Setter
// （协同平台数据合规）药品列表Json
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetDrugListJson(_drugListJson string) error {
    r._drugListJson = _drugListJson
    r.Set("drug_list_json", _drugListJson)
    return nil
}

// DrugListJson Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetDrugListJson() string {
    return r._drugListJson
}
// AssRefEntId Setter
// （协同平台数据合规）单据委托企业refEntId【疫苗药品出库单填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetAssRefEntId(_assRefEntId string) error {
    r._assRefEntId = _assRefEntId
    r.Set("ass_ref_ent_id", _assRefEntId)
    return nil
}

// AssRefEntId Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetAssRefEntId() string {
    return r._assRefEntId
}
// AssEntId Setter
// （协同平台数据合规）单据委托企业entId【疫苗药品出库单填写】
func (r *AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) SetAssEntId(_assEntId string) error {
    r._assEntId = _assEntId
    r.Set("ass_ent_id", _assEntId)
    return nil
}

// AssEntId Getter
func (r AlibabaAlihealthDrugKytDrUploadinoutbillAPIRequest) GetAssEntId() string {
    return r._assEntId
}
