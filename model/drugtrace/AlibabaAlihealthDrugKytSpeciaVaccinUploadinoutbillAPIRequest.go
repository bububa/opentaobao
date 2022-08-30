package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest 疫苗企业出入库上传 API请求
// alibaba.alihealth.drug.kyt.specia.vaccin.uploadinoutbill
//
// 零售企业上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）,
// 不包括对个人的零售出库，疫苗接种，领药出库。
type AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest struct {
	model.Params
	// 追溯码[多个时用逗号分开]
	_traceCodes []string
	// 单据编码
	_billCode string
	// 单据时间
	_billTime string
	// 上传企业的单位编码
	_refUserId string
	// 代理企业REF标识
	_agentRefUserId string
	// 发货企业entId
	_fromUserId string
	// 收货企业entId
	_toUserId string
	// 直调企业标识
	_destUserId string
	// 单据提交者（appkey编号）
	_operIcCode string
	// 单据提交者姓名
	_operIcName string
	// 仓号
	_warehouseId string
	// 药品ID[企业自已系统的药品ID]
	_drugId string
	// 客户端类型[必须填2]
	_clientType string
	// 退货原因代码[退货入出库时填写]
	_returnReasonCode string
	// 退货原因描述[退货入出库时填写]
	_returnReasonDes string
	// 注销原因代码【销毁出库时填写】
	_cancelReasonCode string
	// 注销原因描述【销毁出库时填写】
	_cancelReasonDes string
	// 执行人姓名【销毁出库时填写】
	_executerName string
	// 执行人证件号【销毁出库时填写】
	_executerCode string
	// 监督人姓名【销毁出库时填写】
	_superviserName string
	// 监督人证件号【销毁出库时填写】
	_superviserCode string
	// （协同平台数据合规）发货地址【必选】
	_fromAddress string
	// （协同平台数据合规）收货地址【必选】
	_toAddress string
	// （协同平台数据合规）发货单编号【必选】
	_fromBillCode string
	// （协同平台数据合规）订货单编号【可选】
	_orderCode string
	// （协同平台数据合规）发货人【必选】
	_fromPerson string
	// （协同平台数据合规）收货人【必选】
	_toPerson string
	// （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
	_disRefEntId string
	// （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
	_disEntId string
	// （协同平台数据合规）是否验证，0：未通过验证，1：已验证
	_xtIsCheck string
	// （协同平台数据合规）未验证通过原因【验证未通过时填写】
	_xtCheckCode string
	// （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
	_xtCheckCodeDesc string
	// （协同平台数据合规）药品列表Json
	_drugListJson string
	// （协同平台数据合规）单据委托企业refEntId【疫苗药品出库单填写】
	_assRefEntId string
	// （协同平台数据合规）单据委托企业entId【疫苗药品出库单填写】
	_assEntId string
	// 单据类型【102代表采购入库】
	_billType int64
	// 药品类型【3普药2特药】
	_physicType int64
	// （协同平台数据合规）应收货总数量【必选】
	_quReceivable int64
}

// NewAlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillRequest 初始化AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest对象
func NewAlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillRequest() *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest {
	return &AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.specia.vaccin.uploadinoutbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTraceCodes is TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetTraceCodes(_traceCodes []string) error {
	r._traceCodes = _traceCodes
	r.Set("trace_codes", _traceCodes)
	return nil
}

// GetTraceCodes TraceCodes Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetTraceCodes() []string {
	return r._traceCodes
}

// SetBillCode is BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据时间
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetRefUserId is RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetAgentRefUserId is AgentRefUserId Setter
// 代理企业REF标识
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// GetAgentRefUserId AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// SetFromUserId is FromUserId Setter
// 发货企业entId
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetToUserId is ToUserId Setter
// 收货企业entId
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// GetToUserId ToUserId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetToUserId() string {
	return r._toUserId
}

// SetDestUserId is DestUserId Setter
// 直调企业标识
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetDestUserId(_destUserId string) error {
	r._destUserId = _destUserId
	r.Set("dest_user_id", _destUserId)
	return nil
}

// GetDestUserId DestUserId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetDestUserId() string {
	return r._destUserId
}

// SetOperIcCode is OperIcCode Setter
// 单据提交者（appkey编号）
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// GetOperIcCode OperIcCode Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// SetOperIcName is OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// GetOperIcName OperIcName Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// SetWarehouseId is WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetWarehouseId(_warehouseId string) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetWarehouseId() string {
	return r._warehouseId
}

// SetDrugId is DrugId Setter
// 药品ID[企业自已系统的药品ID]
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// GetDrugId DrugId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetDrugId() string {
	return r._drugId
}

// SetClientType is ClientType Setter
// 客户端类型[必须填2]
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetClientType() string {
	return r._clientType
}

// SetReturnReasonCode is ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetReturnReasonCode(_returnReasonCode string) error {
	r._returnReasonCode = _returnReasonCode
	r.Set("return_reason_code", _returnReasonCode)
	return nil
}

// GetReturnReasonCode ReturnReasonCode Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetReturnReasonCode() string {
	return r._returnReasonCode
}

// SetReturnReasonDes is ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetReturnReasonDes(_returnReasonDes string) error {
	r._returnReasonDes = _returnReasonDes
	r.Set("return_reason_des", _returnReasonDes)
	return nil
}

// GetReturnReasonDes ReturnReasonDes Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetReturnReasonDes() string {
	return r._returnReasonDes
}

// SetCancelReasonCode is CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetCancelReasonCode(_cancelReasonCode string) error {
	r._cancelReasonCode = _cancelReasonCode
	r.Set("cancel_reason_code", _cancelReasonCode)
	return nil
}

// GetCancelReasonCode CancelReasonCode Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetCancelReasonCode() string {
	return r._cancelReasonCode
}

// SetCancelReasonDes is CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetCancelReasonDes(_cancelReasonDes string) error {
	r._cancelReasonDes = _cancelReasonDes
	r.Set("cancel_reason_des", _cancelReasonDes)
	return nil
}

// GetCancelReasonDes CancelReasonDes Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetCancelReasonDes() string {
	return r._cancelReasonDes
}

// SetExecuterName is ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetExecuterName(_executerName string) error {
	r._executerName = _executerName
	r.Set("executer_name", _executerName)
	return nil
}

// GetExecuterName ExecuterName Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetExecuterName() string {
	return r._executerName
}

// SetExecuterCode is ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetExecuterCode(_executerCode string) error {
	r._executerCode = _executerCode
	r.Set("executer_code", _executerCode)
	return nil
}

// GetExecuterCode ExecuterCode Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetExecuterCode() string {
	return r._executerCode
}

// SetSuperviserName is SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetSuperviserName(_superviserName string) error {
	r._superviserName = _superviserName
	r.Set("superviser_name", _superviserName)
	return nil
}

// GetSuperviserName SuperviserName Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetSuperviserName() string {
	return r._superviserName
}

// SetSuperviserCode is SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetSuperviserCode(_superviserCode string) error {
	r._superviserCode = _superviserCode
	r.Set("superviser_code", _superviserCode)
	return nil
}

// GetSuperviserCode SuperviserCode Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetSuperviserCode() string {
	return r._superviserCode
}

// SetFromAddress is FromAddress Setter
// （协同平台数据合规）发货地址【必选】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetFromAddress(_fromAddress string) error {
	r._fromAddress = _fromAddress
	r.Set("from_address", _fromAddress)
	return nil
}

// GetFromAddress FromAddress Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetFromAddress() string {
	return r._fromAddress
}

// SetToAddress is ToAddress Setter
// （协同平台数据合规）收货地址【必选】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetToAddress(_toAddress string) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// GetToAddress ToAddress Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetToAddress() string {
	return r._toAddress
}

// SetFromBillCode is FromBillCode Setter
// （协同平台数据合规）发货单编号【必选】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetFromBillCode(_fromBillCode string) error {
	r._fromBillCode = _fromBillCode
	r.Set("from_bill_code", _fromBillCode)
	return nil
}

// GetFromBillCode FromBillCode Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetFromBillCode() string {
	return r._fromBillCode
}

// SetOrderCode is OrderCode Setter
// （协同平台数据合规）订货单编号【可选】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetFromPerson is FromPerson Setter
// （协同平台数据合规）发货人【必选】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetFromPerson(_fromPerson string) error {
	r._fromPerson = _fromPerson
	r.Set("from_person", _fromPerson)
	return nil
}

// GetFromPerson FromPerson Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetFromPerson() string {
	return r._fromPerson
}

// SetToPerson is ToPerson Setter
// （协同平台数据合规）收货人【必选】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetToPerson(_toPerson string) error {
	r._toPerson = _toPerson
	r.Set("to_person", _toPerson)
	return nil
}

// GetToPerson ToPerson Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetToPerson() string {
	return r._toPerson
}

// SetDisRefEntId is DisRefEntId Setter
// （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetDisRefEntId(_disRefEntId string) error {
	r._disRefEntId = _disRefEntId
	r.Set("dis_ref_ent_id", _disRefEntId)
	return nil
}

// GetDisRefEntId DisRefEntId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetDisRefEntId() string {
	return r._disRefEntId
}

// SetDisEntId is DisEntId Setter
// （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetDisEntId(_disEntId string) error {
	r._disEntId = _disEntId
	r.Set("dis_ent_id", _disEntId)
	return nil
}

// GetDisEntId DisEntId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetDisEntId() string {
	return r._disEntId
}

// SetXtIsCheck is XtIsCheck Setter
// （协同平台数据合规）是否验证，0：未通过验证，1：已验证
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetXtIsCheck(_xtIsCheck string) error {
	r._xtIsCheck = _xtIsCheck
	r.Set("xt_is_check", _xtIsCheck)
	return nil
}

// GetXtIsCheck XtIsCheck Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetXtIsCheck() string {
	return r._xtIsCheck
}

// SetXtCheckCode is XtCheckCode Setter
// （协同平台数据合规）未验证通过原因【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetXtCheckCode(_xtCheckCode string) error {
	r._xtCheckCode = _xtCheckCode
	r.Set("xt_check_code", _xtCheckCode)
	return nil
}

// GetXtCheckCode XtCheckCode Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetXtCheckCode() string {
	return r._xtCheckCode
}

// SetXtCheckCodeDesc is XtCheckCodeDesc Setter
// （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetXtCheckCodeDesc(_xtCheckCodeDesc string) error {
	r._xtCheckCodeDesc = _xtCheckCodeDesc
	r.Set("xt_check_code_desc", _xtCheckCodeDesc)
	return nil
}

// GetXtCheckCodeDesc XtCheckCodeDesc Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetXtCheckCodeDesc() string {
	return r._xtCheckCodeDesc
}

// SetDrugListJson is DrugListJson Setter
// （协同平台数据合规）药品列表Json
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetDrugListJson(_drugListJson string) error {
	r._drugListJson = _drugListJson
	r.Set("drug_list_json", _drugListJson)
	return nil
}

// GetDrugListJson DrugListJson Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetDrugListJson() string {
	return r._drugListJson
}

// SetAssRefEntId is AssRefEntId Setter
// （协同平台数据合规）单据委托企业refEntId【疫苗药品出库单填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetAssRefEntId(_assRefEntId string) error {
	r._assRefEntId = _assRefEntId
	r.Set("ass_ref_ent_id", _assRefEntId)
	return nil
}

// GetAssRefEntId AssRefEntId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetAssRefEntId() string {
	return r._assRefEntId
}

// SetAssEntId is AssEntId Setter
// （协同平台数据合规）单据委托企业entId【疫苗药品出库单填写】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetAssEntId(_assEntId string) error {
	r._assEntId = _assEntId
	r.Set("ass_ent_id", _assEntId)
	return nil
}

// GetAssEntId AssEntId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetAssEntId() string {
	return r._assEntId
}

// SetBillType is BillType Setter
// 单据类型【102代表采购入库】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetBillType() int64 {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型【3普药2特药】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetPhysicType() int64 {
	return r._physicType
}

// SetQuReceivable is QuReceivable Setter
// （协同平台数据合规）应收货总数量【必选】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) SetQuReceivable(_quReceivable int64) error {
	r._quReceivable = _quReceivable
	r.Set("qu_receivable", _quReceivable)
	return nil
}

// GetQuReceivable QuReceivable Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadinoutbillAPIRequest) GetQuReceivable() int64 {
	return r._quReceivable
}
