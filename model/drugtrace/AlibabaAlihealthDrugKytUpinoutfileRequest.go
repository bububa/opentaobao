package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传出入库单据(传文件) API请求
alibaba.alihealth.drug.kyt.upinoutfile

上传出入库单据(传文件)
*/
type AlibabaAlihealthDrugKytUpinoutfileRequest struct {
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
    // 收货企业entId
    _fromUserId   string
    // 发货企业entId
    _toUserId   string
    // 直调企业标识
    _destUserId   string
    // 单据提交者（key编号）
    _operIcCode   string
    // 单据提交者姓名
    _operIcName   string
    // 仓号
    _warehouseId   string
    // 药品ID
    _drugId   string
    // 文件内容
    _fileContent   string
    // 文件名
    _uploadFileName   string
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
}

// 初始化AlibabaAlihealthDrugKytUpinoutfileRequest对象
func NewAlibabaAlihealthDrugKytUpinoutfileRequest() *AlibabaAlihealthDrugKytUpinoutfileRequest{
    return &AlibabaAlihealthDrugKytUpinoutfileRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.upinoutfile"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetBillCode() string {
    return r._billCode
}
// BillTime Setter
// 单据时间
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetBillTime(_billTime string) error {
    r._billTime = _billTime
    r.Set("bill_time", _billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetBillTime() string {
    return r._billTime
}
// BillType Setter
// 单据类型【102代表采购入库】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetBillType(_billType int64) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetBillType() int64 {
    return r._billType
}
// PhysicType Setter
// 药品类型【3普药2特药】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetPhysicType(_physicType int64) error {
    r._physicType = _physicType
    r.Set("physic_type", _physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetPhysicType() int64 {
    return r._physicType
}
// RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetRefUserId(_refUserId string) error {
    r._refUserId = _refUserId
    r.Set("ref_user_id", _refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetRefUserId() string {
    return r._refUserId
}
// AgentRefUserId Setter
// 代理企业REF标识
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetAgentRefUserId(_agentRefUserId string) error {
    r._agentRefUserId = _agentRefUserId
    r.Set("agent_ref_user_id", _agentRefUserId)
    return nil
}

// AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetAgentRefUserId() string {
    return r._agentRefUserId
}
// FromUserId Setter
// 收货企业entId
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetFromUserId(_fromUserId string) error {
    r._fromUserId = _fromUserId
    r.Set("from_user_id", _fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetFromUserId() string {
    return r._fromUserId
}
// ToUserId Setter
// 发货企业entId
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetToUserId(_toUserId string) error {
    r._toUserId = _toUserId
    r.Set("to_user_id", _toUserId)
    return nil
}

// ToUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetToUserId() string {
    return r._toUserId
}
// DestUserId Setter
// 直调企业标识
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetDestUserId(_destUserId string) error {
    r._destUserId = _destUserId
    r.Set("dest_user_id", _destUserId)
    return nil
}

// DestUserId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetDestUserId() string {
    return r._destUserId
}
// OperIcCode Setter
// 单据提交者（key编号）
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetOperIcCode(_operIcCode string) error {
    r._operIcCode = _operIcCode
    r.Set("oper_ic_code", _operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetOperIcCode() string {
    return r._operIcCode
}
// OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetOperIcName(_operIcName string) error {
    r._operIcName = _operIcName
    r.Set("oper_ic_name", _operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetOperIcName() string {
    return r._operIcName
}
// WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetWarehouseId(_warehouseId string) error {
    r._warehouseId = _warehouseId
    r.Set("warehouse_id", _warehouseId)
    return nil
}

// WarehouseId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetWarehouseId() string {
    return r._warehouseId
}
// DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetDrugId(_drugId string) error {
    r._drugId = _drugId
    r.Set("drug_id", _drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetDrugId() string {
    return r._drugId
}
// FileContent Setter
// 文件内容
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetFileContent(_fileContent string) error {
    r._fileContent = _fileContent
    r.Set("file_content", _fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetFileContent() string {
    return r._fileContent
}
// UploadFileName Setter
// 文件名
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetUploadFileName(_uploadFileName string) error {
    r._uploadFileName = _uploadFileName
    r.Set("upload_file_name", _uploadFileName)
    return nil
}

// UploadFileName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetUploadFileName() string {
    return r._uploadFileName
}
// ClientType Setter
// 客户端类型[必须填2]
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetClientType() string {
    return r._clientType
}
// ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetReturnReasonCode(_returnReasonCode string) error {
    r._returnReasonCode = _returnReasonCode
    r.Set("return_reason_code", _returnReasonCode)
    return nil
}

// ReturnReasonCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetReturnReasonCode() string {
    return r._returnReasonCode
}
// ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetReturnReasonDes(_returnReasonDes string) error {
    r._returnReasonDes = _returnReasonDes
    r.Set("return_reason_des", _returnReasonDes)
    return nil
}

// ReturnReasonDes Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetReturnReasonDes() string {
    return r._returnReasonDes
}
// CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetCancelReasonCode(_cancelReasonCode string) error {
    r._cancelReasonCode = _cancelReasonCode
    r.Set("cancel_reason_code", _cancelReasonCode)
    return nil
}

// CancelReasonCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetCancelReasonCode() string {
    return r._cancelReasonCode
}
// CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetCancelReasonDes(_cancelReasonDes string) error {
    r._cancelReasonDes = _cancelReasonDes
    r.Set("cancel_reason_des", _cancelReasonDes)
    return nil
}

// CancelReasonDes Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetCancelReasonDes() string {
    return r._cancelReasonDes
}
// ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetExecuterName(_executerName string) error {
    r._executerName = _executerName
    r.Set("executer_name", _executerName)
    return nil
}

// ExecuterName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetExecuterName() string {
    return r._executerName
}
// ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetExecuterCode(_executerCode string) error {
    r._executerCode = _executerCode
    r.Set("executer_code", _executerCode)
    return nil
}

// ExecuterCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetExecuterCode() string {
    return r._executerCode
}
// SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetSuperviserName(_superviserName string) error {
    r._superviserName = _superviserName
    r.Set("superviser_name", _superviserName)
    return nil
}

// SuperviserName Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetSuperviserName() string {
    return r._superviserName
}
// SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytUpinoutfileRequest) SetSuperviserCode(_superviserCode string) error {
    r._superviserCode = _superviserCode
    r.Set("superviser_code", _superviserCode)
    return nil
}

// SuperviserCode Getter
func (r AlibabaAlihealthDrugKytUpinoutfileRequest) GetSuperviserCode() string {
    return r._superviserCode
}