package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytuploadinoutbillAPIRequest 企业上传出入库信息 API请求
// alibaba.alihealth.drug.kyt.uploadinoutbill
//
// 企业上传出入库信息，包括101, &#34;生产入库&#34;；102, &#34;采购入库&#34;；103, &#34;退货入库&#34;；104, &#34;调拨入库&#34;；106, &#34;零头入库&#34;；107, &#34;供应入库&#34;；108, &#34;召回入库&#34;；110,&#34;赠品入库&#34;；111,&#34;盘盈入库&#34;；112,&#34;报废入库&#34;；113,&#34;其他入库&#34;
// 201, &#34;销售出库&#34;；202, &#34;退货出库&#34;；203, &#34;调拨出库&#34;；204, &#34;返工出库&#34;；205, &#34;销毁出库&#34;；206, &#34;抽检出库&#34;；207, &#34;直调出库&#34;；208, &#34;生产出库&#34;；209, &#34;供应出库&#34;；211, &#34;召回出库&#34;；212,&#34;赠品出库&#34;；214,&#34;盘亏出库&#34;；215,&#34;损坏出库&#34;；216,&#34;报废出库&#34;；217,&#34;其他出库&#34;；237, &#34;直调退货&#34;。
// 不包括对个人的零售出库，疫苗接种，领药出库。
// 本接口与uploadcircubill接口的主要区别的，本接口入参中直接上传追溯码（多个码时用逗号分隔）。uploadcircubill接口入参中，需要上传码的单据文件（用扫码枪生成的xml文件），一般情况下使用uploadcircubill接口上传单据文件。
type AlibabaalihealthdrugkytuploadinoutbillAPIRequest struct {
	model.Params
	// 追溯码【多个码时用逗号拼接的字符串。要求数量在3500个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_traceCodes []string
	// 单据编号【同一个企业不能上传相同单据号】
	_billCode string
	// 单据时间（扫码时间）
	_billTime string
	// 货主（单据的所有者，上传人），上传企业的单位维一编码【注意：该入参是ref_ent_id，不是ent_id】
	_refUserId string
	// 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
	_agentRefUserId string
	// 发货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
	_fromUserId string
	// 收货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
	_toUserId string
	// 直调企业标识【注意：该入参是ent_id,并不是ref_ent_id】
	_destUserId string
	// 单据提交者（调用接口时的appkey编号）
	_operIcCode string
	// 单据提交者姓名（出入库单上传人的名子）
	_operIcName string
	// 仓号
	_warehouseId string
	// 药品ID[企业自已系统的药品ID]
	_drugId string
	// 调用方类型[必须填2]
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
	// 发货地址
	_fromAddress string
	// 收货地址
	_toAddress string
	// 发货单编号
	_fromBillCode string
	// 订货单编号【可选】
	_orderCode string
	// 发货人
	_fromPerson string
	// 收货人
	_toPerson string
	// 药品配送ref_enti_d企业
	_disRefEntId string
	// 药品配送企业ent_Id出库单
	_disEntId string
	// 是否验证，0：未通过验证，1：已验证
	_xtIsCheck string
	// 未验证通过原因【验证未通过时填写】
	_xtCheckCode string
	// 未验证通过原因描述【验证未通过时填写】
	_xtCheckCodeDesc string
	// 药品列表Json："codeCount": 药品数量 "commDrugId": 国家药品唯一标识 "exprieDate": 生产日期 "physicInfo": 药品信息 "pkgSpec": 包状规格 "prepnCount": 制剂数量 "produceBatchNo":生产批次 "produceDate": 生产日期
	_drugListJson string
	// 单据委托企业refEntId
	_assRefEntId string
	// 单据委托企业entId
	_assEntId string
	// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
	_billType int64
	// 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药【可以随便填写，单据上传后会以实际为准】
	_physicType int64
	// 应收货总数量
	_quReceivable int64
}

// NewAlibabaalihealthdrugkytuploadinoutbillRequest 初始化AlibabaalihealthdrugkytuploadinoutbillAPIRequest对象
func NewAlibabaalihealthdrugkytuploadinoutbillRequest() *AlibabaalihealthdrugkytuploadinoutbillAPIRequest {
	return &AlibabaalihealthdrugkytuploadinoutbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.uploadinoutbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceCodes is TraceCodes Setter
// 追溯码【多个码时用逗号拼接的字符串。要求数量在3500个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetTraceCodes(_traceCodes []string) error {
	r._traceCodes = _traceCodes
	r.Set("trace_codes", _traceCodes)
	return nil
}

// GetTraceCodes TraceCodes Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetTraceCodes() []string {
	return r._traceCodes
}

// SetBillCode is BillCode Setter
// 单据编号【同一个企业不能上传相同单据号】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据时间（扫码时间）
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetRefUserId is RefUserId Setter
// 货主（单据的所有者，上传人），上传企业的单位维一编码【注意：该入参是ref_ent_id，不是ent_id】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetAgentRefUserId is AgentRefUserId Setter
// 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// GetAgentRefUserId AgentRefUserId Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// SetFromUserId is FromUserId Setter
// 发货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetToUserId is ToUserId Setter
// 收货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// GetToUserId ToUserId Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetToUserId() string {
	return r._toUserId
}

// SetDestUserId is DestUserId Setter
// 直调企业标识【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetDestUserId(_destUserId string) error {
	r._destUserId = _destUserId
	r.Set("dest_user_id", _destUserId)
	return nil
}

// GetDestUserId DestUserId Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetDestUserId() string {
	return r._destUserId
}

// SetOperIcCode is OperIcCode Setter
// 单据提交者（调用接口时的appkey编号）
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// GetOperIcCode OperIcCode Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// SetOperIcName is OperIcName Setter
// 单据提交者姓名（出入库单上传人的名子）
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// GetOperIcName OperIcName Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// SetWarehouseId is WarehouseId Setter
// 仓号
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetWarehouseId(_warehouseId string) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetWarehouseId() string {
	return r._warehouseId
}

// SetDrugId is DrugId Setter
// 药品ID[企业自已系统的药品ID]
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// GetDrugId DrugId Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetDrugId() string {
	return r._drugId
}

// SetClientType is ClientType Setter
// 调用方类型[必须填2]
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetClientType() string {
	return r._clientType
}

// SetReturnReasonCode is ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetReturnReasonCode(_returnReasonCode string) error {
	r._returnReasonCode = _returnReasonCode
	r.Set("return_reason_code", _returnReasonCode)
	return nil
}

// GetReturnReasonCode ReturnReasonCode Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetReturnReasonCode() string {
	return r._returnReasonCode
}

// SetReturnReasonDes is ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetReturnReasonDes(_returnReasonDes string) error {
	r._returnReasonDes = _returnReasonDes
	r.Set("return_reason_des", _returnReasonDes)
	return nil
}

// GetReturnReasonDes ReturnReasonDes Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetReturnReasonDes() string {
	return r._returnReasonDes
}

// SetCancelReasonCode is CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetCancelReasonCode(_cancelReasonCode string) error {
	r._cancelReasonCode = _cancelReasonCode
	r.Set("cancel_reason_code", _cancelReasonCode)
	return nil
}

// GetCancelReasonCode CancelReasonCode Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetCancelReasonCode() string {
	return r._cancelReasonCode
}

// SetCancelReasonDes is CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetCancelReasonDes(_cancelReasonDes string) error {
	r._cancelReasonDes = _cancelReasonDes
	r.Set("cancel_reason_des", _cancelReasonDes)
	return nil
}

// GetCancelReasonDes CancelReasonDes Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetCancelReasonDes() string {
	return r._cancelReasonDes
}

// SetExecuterName is ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetExecuterName(_executerName string) error {
	r._executerName = _executerName
	r.Set("executer_name", _executerName)
	return nil
}

// GetExecuterName ExecuterName Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetExecuterName() string {
	return r._executerName
}

// SetExecuterCode is ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetExecuterCode(_executerCode string) error {
	r._executerCode = _executerCode
	r.Set("executer_code", _executerCode)
	return nil
}

// GetExecuterCode ExecuterCode Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetExecuterCode() string {
	return r._executerCode
}

// SetSuperviserName is SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetSuperviserName(_superviserName string) error {
	r._superviserName = _superviserName
	r.Set("superviser_name", _superviserName)
	return nil
}

// GetSuperviserName SuperviserName Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetSuperviserName() string {
	return r._superviserName
}

// SetSuperviserCode is SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetSuperviserCode(_superviserCode string) error {
	r._superviserCode = _superviserCode
	r.Set("superviser_code", _superviserCode)
	return nil
}

// GetSuperviserCode SuperviserCode Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetSuperviserCode() string {
	return r._superviserCode
}

// SetFromAddress is FromAddress Setter
// 发货地址
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetFromAddress(_fromAddress string) error {
	r._fromAddress = _fromAddress
	r.Set("from_address", _fromAddress)
	return nil
}

// GetFromAddress FromAddress Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetFromAddress() string {
	return r._fromAddress
}

// SetToAddress is ToAddress Setter
// 收货地址
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetToAddress(_toAddress string) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// GetToAddress ToAddress Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetToAddress() string {
	return r._toAddress
}

// SetFromBillCode is FromBillCode Setter
// 发货单编号
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetFromBillCode(_fromBillCode string) error {
	r._fromBillCode = _fromBillCode
	r.Set("from_bill_code", _fromBillCode)
	return nil
}

// GetFromBillCode FromBillCode Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetFromBillCode() string {
	return r._fromBillCode
}

// SetOrderCode is OrderCode Setter
// 订货单编号【可选】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetFromPerson is FromPerson Setter
// 发货人
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetFromPerson(_fromPerson string) error {
	r._fromPerson = _fromPerson
	r.Set("from_person", _fromPerson)
	return nil
}

// GetFromPerson FromPerson Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetFromPerson() string {
	return r._fromPerson
}

// SetToPerson is ToPerson Setter
// 收货人
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetToPerson(_toPerson string) error {
	r._toPerson = _toPerson
	r.Set("to_person", _toPerson)
	return nil
}

// GetToPerson ToPerson Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetToPerson() string {
	return r._toPerson
}

// SetDisRefEntId is DisRefEntId Setter
// 药品配送ref_enti_d企业
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetDisRefEntId(_disRefEntId string) error {
	r._disRefEntId = _disRefEntId
	r.Set("dis_ref_ent_id", _disRefEntId)
	return nil
}

// GetDisRefEntId DisRefEntId Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetDisRefEntId() string {
	return r._disRefEntId
}

// SetDisEntId is DisEntId Setter
// 药品配送企业ent_Id出库单
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetDisEntId(_disEntId string) error {
	r._disEntId = _disEntId
	r.Set("dis_ent_id", _disEntId)
	return nil
}

// GetDisEntId DisEntId Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetDisEntId() string {
	return r._disEntId
}

// SetXtIsCheck is XtIsCheck Setter
// 是否验证，0：未通过验证，1：已验证
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetXtIsCheck(_xtIsCheck string) error {
	r._xtIsCheck = _xtIsCheck
	r.Set("xt_is_check", _xtIsCheck)
	return nil
}

// GetXtIsCheck XtIsCheck Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetXtIsCheck() string {
	return r._xtIsCheck
}

// SetXtCheckCode is XtCheckCode Setter
// 未验证通过原因【验证未通过时填写】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetXtCheckCode(_xtCheckCode string) error {
	r._xtCheckCode = _xtCheckCode
	r.Set("xt_check_code", _xtCheckCode)
	return nil
}

// GetXtCheckCode XtCheckCode Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetXtCheckCode() string {
	return r._xtCheckCode
}

// SetXtCheckCodeDesc is XtCheckCodeDesc Setter
// 未验证通过原因描述【验证未通过时填写】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetXtCheckCodeDesc(_xtCheckCodeDesc string) error {
	r._xtCheckCodeDesc = _xtCheckCodeDesc
	r.Set("xt_check_code_desc", _xtCheckCodeDesc)
	return nil
}

// GetXtCheckCodeDesc XtCheckCodeDesc Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetXtCheckCodeDesc() string {
	return r._xtCheckCodeDesc
}

// SetDrugListJson is DrugListJson Setter
// 药品列表Json：&#34;codeCount&#34;: 药品数量 &#34;commDrugId&#34;: 国家药品唯一标识 &#34;exprieDate&#34;: 生产日期 &#34;physicInfo&#34;: 药品信息 &#34;pkgSpec&#34;: 包状规格 &#34;prepnCount&#34;: 制剂数量 &#34;produceBatchNo&#34;:生产批次 &#34;produceDate&#34;: 生产日期
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetDrugListJson(_drugListJson string) error {
	r._drugListJson = _drugListJson
	r.Set("drug_list_json", _drugListJson)
	return nil
}

// GetDrugListJson DrugListJson Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetDrugListJson() string {
	return r._drugListJson
}

// SetAssRefEntId is AssRefEntId Setter
// 单据委托企业refEntId
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetAssRefEntId(_assRefEntId string) error {
	r._assRefEntId = _assRefEntId
	r.Set("ass_ref_ent_id", _assRefEntId)
	return nil
}

// GetAssRefEntId AssRefEntId Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetAssRefEntId() string {
	return r._assRefEntId
}

// SetAssEntId is AssEntId Setter
// 单据委托企业entId
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetAssEntId(_assEntId string) error {
	r._assEntId = _assEntId
	r.Set("ass_ent_id", _assEntId)
	return nil
}

// GetAssEntId AssEntId Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetAssEntId() string {
	return r._assEntId
}

// SetBillType is BillType Setter
// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetBillType() int64 {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药【可以随便填写，单据上传后会以实际为准】
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetPhysicType() int64 {
	return r._physicType
}

// SetQuReceivable is QuReceivable Setter
// 应收货总数量
func (r *AlibabaalihealthdrugkytuploadinoutbillAPIRequest) SetQuReceivable(_quReceivable int64) error {
	r._quReceivable = _quReceivable
	r.Set("qu_receivable", _quReceivable)
	return nil
}

// GetQuReceivable QuReceivable Getter
func (r AlibabaalihealthdrugkytuploadinoutbillAPIRequest) GetQuReceivable() int64 {
	return r._quReceivable
}
