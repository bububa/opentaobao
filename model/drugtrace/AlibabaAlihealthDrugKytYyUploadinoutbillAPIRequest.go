package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest 医院出入库信息上传 API请求
// alibaba.alihealth.drug.kyt.yy.uploadinoutbill
//
// 医院上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）,
// 不包括对个人的零售出库，疫苗接种，领药出库。
type AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest struct {
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
	// 单据类型【102代表采购入库】
	_billType int64
	// 药品类型【3普药2特药】
	_physicType int64
}

// NewAlibabaalihealthdrugkytyyuploadinoutbillRequest 初始化AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest对象
func NewAlibabaalihealthdrugkytyyuploadinoutbillRequest() *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest {
	return &AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yy.uploadinoutbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceCodes is TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetTraceCodes(_traceCodes []string) error {
	r._traceCodes = _traceCodes
	r.Set("trace_codes", _traceCodes)
	return nil
}

// GetTraceCodes TraceCodes Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetTraceCodes() []string {
	return r._traceCodes
}

// SetBillCode is BillCode Setter
// 单据编码
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据时间
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetRefUserId is RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetAgentRefUserId is AgentRefUserId Setter
// 代理企业REF标识
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// GetAgentRefUserId AgentRefUserId Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// SetFromUserId is FromUserId Setter
// 发货企业entId
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetToUserId is ToUserId Setter
// 收货企业entId
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// GetToUserId ToUserId Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetToUserId() string {
	return r._toUserId
}

// SetDestUserId is DestUserId Setter
// 直调企业标识
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetDestUserId(_destUserId string) error {
	r._destUserId = _destUserId
	r.Set("dest_user_id", _destUserId)
	return nil
}

// GetDestUserId DestUserId Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetDestUserId() string {
	return r._destUserId
}

// SetOperIcCode is OperIcCode Setter
// 单据提交者（appkey编号）
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// GetOperIcCode OperIcCode Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// SetOperIcName is OperIcName Setter
// 单据提交者姓名
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// GetOperIcName OperIcName Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// SetWarehouseId is WarehouseId Setter
// 仓号
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetWarehouseId(_warehouseId string) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetWarehouseId() string {
	return r._warehouseId
}

// SetDrugId is DrugId Setter
// 药品ID[企业自已系统的药品ID]
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// GetDrugId DrugId Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetDrugId() string {
	return r._drugId
}

// SetClientType is ClientType Setter
// 客户端类型[必须填2]
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetClientType() string {
	return r._clientType
}

// SetReturnReasonCode is ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetReturnReasonCode(_returnReasonCode string) error {
	r._returnReasonCode = _returnReasonCode
	r.Set("return_reason_code", _returnReasonCode)
	return nil
}

// GetReturnReasonCode ReturnReasonCode Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetReturnReasonCode() string {
	return r._returnReasonCode
}

// SetReturnReasonDes is ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetReturnReasonDes(_returnReasonDes string) error {
	r._returnReasonDes = _returnReasonDes
	r.Set("return_reason_des", _returnReasonDes)
	return nil
}

// GetReturnReasonDes ReturnReasonDes Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetReturnReasonDes() string {
	return r._returnReasonDes
}

// SetCancelReasonCode is CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetCancelReasonCode(_cancelReasonCode string) error {
	r._cancelReasonCode = _cancelReasonCode
	r.Set("cancel_reason_code", _cancelReasonCode)
	return nil
}

// GetCancelReasonCode CancelReasonCode Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetCancelReasonCode() string {
	return r._cancelReasonCode
}

// SetCancelReasonDes is CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetCancelReasonDes(_cancelReasonDes string) error {
	r._cancelReasonDes = _cancelReasonDes
	r.Set("cancel_reason_des", _cancelReasonDes)
	return nil
}

// GetCancelReasonDes CancelReasonDes Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetCancelReasonDes() string {
	return r._cancelReasonDes
}

// SetExecuterName is ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetExecuterName(_executerName string) error {
	r._executerName = _executerName
	r.Set("executer_name", _executerName)
	return nil
}

// GetExecuterName ExecuterName Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetExecuterName() string {
	return r._executerName
}

// SetExecuterCode is ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetExecuterCode(_executerCode string) error {
	r._executerCode = _executerCode
	r.Set("executer_code", _executerCode)
	return nil
}

// GetExecuterCode ExecuterCode Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetExecuterCode() string {
	return r._executerCode
}

// SetSuperviserName is SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetSuperviserName(_superviserName string) error {
	r._superviserName = _superviserName
	r.Set("superviser_name", _superviserName)
	return nil
}

// GetSuperviserName SuperviserName Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetSuperviserName() string {
	return r._superviserName
}

// SetSuperviserCode is SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetSuperviserCode(_superviserCode string) error {
	r._superviserCode = _superviserCode
	r.Set("superviser_code", _superviserCode)
	return nil
}

// GetSuperviserCode SuperviserCode Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetSuperviserCode() string {
	return r._superviserCode
}

// SetBillType is BillType Setter
// 单据类型【102代表采购入库】
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetBillType() int64 {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型【3普药2特药】
func (r *AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaalihealthdrugkytyyuploadinoutbillAPIRequest) GetPhysicType() int64 {
	return r._physicType
}
