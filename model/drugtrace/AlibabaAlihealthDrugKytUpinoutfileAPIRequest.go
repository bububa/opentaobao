package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUpinoutfileAPIRequest 上传出入库单据(传文件) API请求
// alibaba.alihealth.drug.kyt.upinoutfile
//
// 上传出入库单据(传文件)
type AlibabaAlihealthDrugKytUpinoutfileAPIRequest struct {
	model.Params
	// 单据编码
	_billCode string
	// 单据时间
	_billTime string
	// 上传企业的单位编码
	_refUserId string
	// 代理企业REF标识
	_agentRefUserId string
	// 收货企业entId
	_fromUserId string
	// 发货企业entId
	_toUserId string
	// 直调企业标识
	_destUserId string
	// 单据提交者（key编号）
	_operIcCode string
	// 单据提交者姓名
	_operIcName string
	// 仓号
	_warehouseId string
	// 药品ID
	_drugId string
	// 文件内容
	_fileContent string
	// 文件名
	_uploadFileName string
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

// NewAlibabaAlihealthDrugKytUpinoutfileRequest 初始化AlibabaAlihealthDrugKytUpinoutfileAPIRequest对象
func NewAlibabaAlihealthDrugKytUpinoutfileRequest() *AlibabaAlihealthDrugKytUpinoutfileAPIRequest {
	return &AlibabaAlihealthDrugKytUpinoutfileAPIRequest{
		Params: model.NewParams(24),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) Reset() {
	r._billCode = ""
	r._billTime = ""
	r._refUserId = ""
	r._agentRefUserId = ""
	r._fromUserId = ""
	r._toUserId = ""
	r._destUserId = ""
	r._operIcCode = ""
	r._operIcName = ""
	r._warehouseId = ""
	r._drugId = ""
	r._fileContent = ""
	r._uploadFileName = ""
	r._clientType = ""
	r._returnReasonCode = ""
	r._returnReasonDes = ""
	r._cancelReasonCode = ""
	r._cancelReasonDes = ""
	r._executerName = ""
	r._executerCode = ""
	r._superviserName = ""
	r._superviserCode = ""
	r._billType = 0
	r._physicType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.upinoutfile"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillCode is BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据时间
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetRefUserId is RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetAgentRefUserId is AgentRefUserId Setter
// 代理企业REF标识
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// GetAgentRefUserId AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// SetFromUserId is FromUserId Setter
// 收货企业entId
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetToUserId is ToUserId Setter
// 发货企业entId
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// GetToUserId ToUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetToUserId() string {
	return r._toUserId
}

// SetDestUserId is DestUserId Setter
// 直调企业标识
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetDestUserId(_destUserId string) error {
	r._destUserId = _destUserId
	r.Set("dest_user_id", _destUserId)
	return nil
}

// GetDestUserId DestUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetDestUserId() string {
	return r._destUserId
}

// SetOperIcCode is OperIcCode Setter
// 单据提交者（key编号）
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// GetOperIcCode OperIcCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// SetOperIcName is OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// GetOperIcName OperIcName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// SetWarehouseId is WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetWarehouseId(_warehouseId string) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetWarehouseId() string {
	return r._warehouseId
}

// SetDrugId is DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// GetDrugId DrugId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetDrugId() string {
	return r._drugId
}

// SetFileContent is FileContent Setter
// 文件内容
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetFileContent(_fileContent string) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// GetFileContent FileContent Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetFileContent() string {
	return r._fileContent
}

// SetUploadFileName is UploadFileName Setter
// 文件名
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetUploadFileName(_uploadFileName string) error {
	r._uploadFileName = _uploadFileName
	r.Set("upload_file_name", _uploadFileName)
	return nil
}

// GetUploadFileName UploadFileName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetUploadFileName() string {
	return r._uploadFileName
}

// SetClientType is ClientType Setter
// 客户端类型[必须填2]
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetClientType() string {
	return r._clientType
}

// SetReturnReasonCode is ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetReturnReasonCode(_returnReasonCode string) error {
	r._returnReasonCode = _returnReasonCode
	r.Set("return_reason_code", _returnReasonCode)
	return nil
}

// GetReturnReasonCode ReturnReasonCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetReturnReasonCode() string {
	return r._returnReasonCode
}

// SetReturnReasonDes is ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetReturnReasonDes(_returnReasonDes string) error {
	r._returnReasonDes = _returnReasonDes
	r.Set("return_reason_des", _returnReasonDes)
	return nil
}

// GetReturnReasonDes ReturnReasonDes Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetReturnReasonDes() string {
	return r._returnReasonDes
}

// SetCancelReasonCode is CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetCancelReasonCode(_cancelReasonCode string) error {
	r._cancelReasonCode = _cancelReasonCode
	r.Set("cancel_reason_code", _cancelReasonCode)
	return nil
}

// GetCancelReasonCode CancelReasonCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetCancelReasonCode() string {
	return r._cancelReasonCode
}

// SetCancelReasonDes is CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetCancelReasonDes(_cancelReasonDes string) error {
	r._cancelReasonDes = _cancelReasonDes
	r.Set("cancel_reason_des", _cancelReasonDes)
	return nil
}

// GetCancelReasonDes CancelReasonDes Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetCancelReasonDes() string {
	return r._cancelReasonDes
}

// SetExecuterName is ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetExecuterName(_executerName string) error {
	r._executerName = _executerName
	r.Set("executer_name", _executerName)
	return nil
}

// GetExecuterName ExecuterName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetExecuterName() string {
	return r._executerName
}

// SetExecuterCode is ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetExecuterCode(_executerCode string) error {
	r._executerCode = _executerCode
	r.Set("executer_code", _executerCode)
	return nil
}

// GetExecuterCode ExecuterCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetExecuterCode() string {
	return r._executerCode
}

// SetSuperviserName is SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetSuperviserName(_superviserName string) error {
	r._superviserName = _superviserName
	r.Set("superviser_name", _superviserName)
	return nil
}

// GetSuperviserName SuperviserName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetSuperviserName() string {
	return r._superviserName
}

// SetSuperviserCode is SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetSuperviserCode(_superviserCode string) error {
	r._superviserCode = _superviserCode
	r.Set("superviser_code", _superviserCode)
	return nil
}

// GetSuperviserCode SuperviserCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetSuperviserCode() string {
	return r._superviserCode
}

// SetBillType is BillType Setter
// 单据类型【102代表采购入库】
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetBillType() int64 {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型【3普药2特药】
func (r *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaAlihealthDrugKytUpinoutfileAPIRequest) GetPhysicType() int64 {
	return r._physicType
}

var poolAlibabaAlihealthDrugKytUpinoutfileAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytUpinoutfileRequest()
	},
}

// GetAlibabaAlihealthDrugKytUpinoutfileRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytUpinoutfileAPIRequest
func GetAlibabaAlihealthDrugKytUpinoutfileAPIRequest() *AlibabaAlihealthDrugKytUpinoutfileAPIRequest {
	return poolAlibabaAlihealthDrugKytUpinoutfileAPIRequest.Get().(*AlibabaAlihealthDrugKytUpinoutfileAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytUpinoutfileAPIRequest 将 AlibabaAlihealthDrugKytUpinoutfileAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytUpinoutfileAPIRequest(v *AlibabaAlihealthDrugKytUpinoutfileAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytUpinoutfileAPIRequest.Put(v)
}
