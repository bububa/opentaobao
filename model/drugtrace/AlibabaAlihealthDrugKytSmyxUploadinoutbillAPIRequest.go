package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest 药店出入库信息上传 API请求
// alibaba.alihealth.drug.kyt.smyx.uploadinoutbill
//
// 药店上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）,
// 不包括对个人的零售出库，疫苗接种，领药出库。
type AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest struct {
	model.Params
	// 单据编码
	_billCode string
	// 单据时间
	_billTime string
	// 单据类型【102代表采购入库】
	_billType int64
	// 药品类型【3普药2特药】
	_physicType int64
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
	// 仓号
	_warehouseId string
	// 药品ID[企业自已系统的药品ID]
	_drugId string
	// 追溯码[多个时用逗号分开]
	_traceCodes []string
}

// NewAlibabaAlihealthDrugKytSmyxUploadinoutbillRequest 初始化AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest对象
func NewAlibabaAlihealthDrugKytSmyxUploadinoutbillRequest() *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest {
	return &AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.smyx.uploadinoutbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is BillTime Setter
// 单据时间
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// Get BillTime Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetBillTime() string {
	return r._billTime
}

// Set is BillType Setter
// 单据类型【102代表采购入库】
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// Get BillType Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetBillType() int64 {
	return r._billType
}

// Set is PhysicType Setter
// 药品类型【3普药2特药】
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// Get PhysicType Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetPhysicType() int64 {
	return r._physicType
}

// Set is RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// Get RefUserId Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// Set is AgentRefUserId Setter
// 代理企业REF标识
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// Get AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// Set is FromUserId Setter
// 发货企业entId
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// Get FromUserId Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// Set is ToUserId Setter
// 收货企业entId
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// Get ToUserId Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetToUserId() string {
	return r._toUserId
}

// Set is DestUserId Setter
// 直调企业标识
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetDestUserId(_destUserId string) error {
	r._destUserId = _destUserId
	r.Set("dest_user_id", _destUserId)
	return nil
}

// Get DestUserId Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetDestUserId() string {
	return r._destUserId
}

// Set is OperIcCode Setter
// 单据提交者（appkey编号）
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// Get OperIcCode Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// Set is OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// Get OperIcName Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// Set is ClientType Setter
// 客户端类型[必须填2]
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// Get ClientType Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetClientType() string {
	return r._clientType
}

// Set is ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetReturnReasonCode(_returnReasonCode string) error {
	r._returnReasonCode = _returnReasonCode
	r.Set("return_reason_code", _returnReasonCode)
	return nil
}

// Get ReturnReasonCode Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetReturnReasonCode() string {
	return r._returnReasonCode
}

// Set is ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetReturnReasonDes(_returnReasonDes string) error {
	r._returnReasonDes = _returnReasonDes
	r.Set("return_reason_des", _returnReasonDes)
	return nil
}

// Get ReturnReasonDes Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetReturnReasonDes() string {
	return r._returnReasonDes
}

// Set is CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetCancelReasonCode(_cancelReasonCode string) error {
	r._cancelReasonCode = _cancelReasonCode
	r.Set("cancel_reason_code", _cancelReasonCode)
	return nil
}

// Get CancelReasonCode Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetCancelReasonCode() string {
	return r._cancelReasonCode
}

// Set is CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetCancelReasonDes(_cancelReasonDes string) error {
	r._cancelReasonDes = _cancelReasonDes
	r.Set("cancel_reason_des", _cancelReasonDes)
	return nil
}

// Get CancelReasonDes Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetCancelReasonDes() string {
	return r._cancelReasonDes
}

// Set is ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetExecuterName(_executerName string) error {
	r._executerName = _executerName
	r.Set("executer_name", _executerName)
	return nil
}

// Get ExecuterName Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetExecuterName() string {
	return r._executerName
}

// Set is ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetExecuterCode(_executerCode string) error {
	r._executerCode = _executerCode
	r.Set("executer_code", _executerCode)
	return nil
}

// Get ExecuterCode Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetExecuterCode() string {
	return r._executerCode
}

// Set is SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetSuperviserName(_superviserName string) error {
	r._superviserName = _superviserName
	r.Set("superviser_name", _superviserName)
	return nil
}

// Get SuperviserName Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetSuperviserName() string {
	return r._superviserName
}

// Set is SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetSuperviserCode(_superviserCode string) error {
	r._superviserCode = _superviserCode
	r.Set("superviser_code", _superviserCode)
	return nil
}

// Get SuperviserCode Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetSuperviserCode() string {
	return r._superviserCode
}

// Set is WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetWarehouseId(_warehouseId string) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// Get WarehouseId Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetWarehouseId() string {
	return r._warehouseId
}

// Set is DrugId Setter
// 药品ID[企业自已系统的药品ID]
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// Get DrugId Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetDrugId() string {
	return r._drugId
}

// Set is TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) SetTraceCodes(_traceCodes []string) error {
	r._traceCodes = _traceCodes
	r.Set("trace_codes", _traceCodes)
	return nil
}

// Get TraceCodes Getter
func (r AlibabaAlihealthDrugKytSmyxUploadinoutbillAPIRequest) GetTraceCodes() []string {
	return r._traceCodes
}
