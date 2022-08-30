package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUploadcircubillAPIRequest 生产批发单据上传 API请求
// alibaba.alihealth.drug.kyt.uploadcircubill
//
// 生产批发单据上传（非零售企业使用），包括101, "生产入库"；102, "采购入库"；103, "退货入库"；104, "调拨入库"；106, "零头入库"；107, "供应入库"；108, "召回入库"；110,"赠品入库"；111,"盘盈入库"；112,"报废入库"；113,"其他入库"
// 201, "销售出库"；202, "退货出库"；203, "调拨出库"；204, "返工出库"；205, "销毁出库"；206, "抽检出库"；207, "直调出库"；208, "生产出库"；209, "供应出库"；211, "召回出库"；212,"赠品出库"；214,"盘亏出库"；215,"损坏出库"；216,"报废出库"；217,"其他出库"；237, "直调退货"。
type AlibabaAlihealthDrugKytUploadcircubillAPIRequest struct {
	model.Params
	// 单据编号【同一个企业不能上传相同单据号】
	_billCode string
	// 单据时间（扫码时间、出入库时间）
	_billTime string
	// 货主企业（单据的所有者，上传人）【注意：该入参是ref_ent_id，不是ent_id】
	_refUserId string
	// 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
	_agentRefUserId string
	// 发货企业【注意：该入参是ent_id,并不是ref_ent_id】
	_fromUserId string
	// 收货企业【注意：该入参是ent_id,并不是ref_ent_id】
	_toUserId string
	// 直调企业【注意：该入参是ent_id,并不是ref_ent_id】
	_destUserId string
	// 操作人标识（appkey编号）
	_operIcCode string
	// 单据提交者姓名
	_operIcName string
	// 单据文件体【bas64字符串】，看对接文档中的代码示例，示例中有相应说明。
	_fileContent string
	// 单据文件名
	_uploadFileName string
	// 客户端类型【暂定都写2】
	_clientType string
	// 发货地址【必选】
	_fromAddress string
	// 收货地址【必选】
	_toAddress string
	// 发货单编号【必选】
	_fromBillCode string
	// 订货单编号【可选】
	_orderCode string
	// 发货人【特药时必选】
	_fromPerson string
	// 收货人【特药时必选】
	_toPerson string
	// 药品配送企业【填写 ref_ent_Id，与dis_ent_id属性选填一个就行】
	_disRefEntId string
	// 药品配送企业【填写 ent_Id】
	_disEntId string
	// 是否验证，0：未通过验证，1：已验证
	_xtIsCheck string
	// 未验证通过原因【验证未通过时填写】
	_xtCheckCode string
	// 未验证通过原因描述【验证未通过时填写】
	_xtCheckCodeDesc string
	// 药品列表【可不填写，不填写时会通过码进行计算】Json："codeCount":         药品数量 "commDrugId":     国家药品唯一标识 "exprieDate":         生产日期 "physicInfo":          药品信息 "pkgSpec":           包状规格 "prepnCount":       制剂数量 "produceBatchNo":生产批次 "produceDate":      生产日期
	_drugListJson string
	// 单据委托企业【填写 ref_ent_Id，与ass_ent_id属性选填一个就行】
	_assRefEntId string
	// 单据委托企业【填写 ent_Id】
	_assEntId string
	// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
	_billType int64
	// 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药【可以随便填写，单据上传后会以实际为准】
	_physicType int64
	// 应收货总数量
	_quReceivable int64
}

// NewAlibabaAlihealthDrugKytUploadcircubillRequest 初始化AlibabaAlihealthDrugKytUploadcircubillAPIRequest对象
func NewAlibabaAlihealthDrugKytUploadcircubillRequest() *AlibabaAlihealthDrugKytUploadcircubillAPIRequest {
	return &AlibabaAlihealthDrugKytUploadcircubillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.uploadcircubill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBillCode is BillCode Setter
// 单据编号【同一个企业不能上传相同单据号】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据时间（扫码时间、出入库时间）
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetRefUserId is RefUserId Setter
// 货主企业（单据的所有者，上传人）【注意：该入参是ref_ent_id，不是ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetAgentRefUserId is AgentRefUserId Setter
// 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// GetAgentRefUserId AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// SetFromUserId is FromUserId Setter
// 发货企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetToUserId is ToUserId Setter
// 收货企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// GetToUserId ToUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetToUserId() string {
	return r._toUserId
}

// SetDestUserId is DestUserId Setter
// 直调企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetDestUserId(_destUserId string) error {
	r._destUserId = _destUserId
	r.Set("dest_user_id", _destUserId)
	return nil
}

// GetDestUserId DestUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetDestUserId() string {
	return r._destUserId
}

// SetOperIcCode is OperIcCode Setter
// 操作人标识（appkey编号）
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// GetOperIcCode OperIcCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// SetOperIcName is OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// GetOperIcName OperIcName Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// SetFileContent is FileContent Setter
// 单据文件体【bas64字符串】，看对接文档中的代码示例，示例中有相应说明。
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetFileContent(_fileContent string) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// GetFileContent FileContent Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetFileContent() string {
	return r._fileContent
}

// SetUploadFileName is UploadFileName Setter
// 单据文件名
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetUploadFileName(_uploadFileName string) error {
	r._uploadFileName = _uploadFileName
	r.Set("upload_file_name", _uploadFileName)
	return nil
}

// GetUploadFileName UploadFileName Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetUploadFileName() string {
	return r._uploadFileName
}

// SetClientType is ClientType Setter
// 客户端类型【暂定都写2】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetClientType() string {
	return r._clientType
}

// SetFromAddress is FromAddress Setter
// 发货地址【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetFromAddress(_fromAddress string) error {
	r._fromAddress = _fromAddress
	r.Set("from_address", _fromAddress)
	return nil
}

// GetFromAddress FromAddress Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetFromAddress() string {
	return r._fromAddress
}

// SetToAddress is ToAddress Setter
// 收货地址【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetToAddress(_toAddress string) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// GetToAddress ToAddress Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetToAddress() string {
	return r._toAddress
}

// SetFromBillCode is FromBillCode Setter
// 发货单编号【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetFromBillCode(_fromBillCode string) error {
	r._fromBillCode = _fromBillCode
	r.Set("from_bill_code", _fromBillCode)
	return nil
}

// GetFromBillCode FromBillCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetFromBillCode() string {
	return r._fromBillCode
}

// SetOrderCode is OrderCode Setter
// 订货单编号【可选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetFromPerson is FromPerson Setter
// 发货人【特药时必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetFromPerson(_fromPerson string) error {
	r._fromPerson = _fromPerson
	r.Set("from_person", _fromPerson)
	return nil
}

// GetFromPerson FromPerson Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetFromPerson() string {
	return r._fromPerson
}

// SetToPerson is ToPerson Setter
// 收货人【特药时必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetToPerson(_toPerson string) error {
	r._toPerson = _toPerson
	r.Set("to_person", _toPerson)
	return nil
}

// GetToPerson ToPerson Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetToPerson() string {
	return r._toPerson
}

// SetDisRefEntId is DisRefEntId Setter
// 药品配送企业【填写 ref_ent_Id，与dis_ent_id属性选填一个就行】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetDisRefEntId(_disRefEntId string) error {
	r._disRefEntId = _disRefEntId
	r.Set("dis_ref_ent_id", _disRefEntId)
	return nil
}

// GetDisRefEntId DisRefEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetDisRefEntId() string {
	return r._disRefEntId
}

// SetDisEntId is DisEntId Setter
// 药品配送企业【填写 ent_Id】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetDisEntId(_disEntId string) error {
	r._disEntId = _disEntId
	r.Set("dis_ent_id", _disEntId)
	return nil
}

// GetDisEntId DisEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetDisEntId() string {
	return r._disEntId
}

// SetXtIsCheck is XtIsCheck Setter
// 是否验证，0：未通过验证，1：已验证
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetXtIsCheck(_xtIsCheck string) error {
	r._xtIsCheck = _xtIsCheck
	r.Set("xt_is_check", _xtIsCheck)
	return nil
}

// GetXtIsCheck XtIsCheck Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetXtIsCheck() string {
	return r._xtIsCheck
}

// SetXtCheckCode is XtCheckCode Setter
// 未验证通过原因【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetXtCheckCode(_xtCheckCode string) error {
	r._xtCheckCode = _xtCheckCode
	r.Set("xt_check_code", _xtCheckCode)
	return nil
}

// GetXtCheckCode XtCheckCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetXtCheckCode() string {
	return r._xtCheckCode
}

// SetXtCheckCodeDesc is XtCheckCodeDesc Setter
// 未验证通过原因描述【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetXtCheckCodeDesc(_xtCheckCodeDesc string) error {
	r._xtCheckCodeDesc = _xtCheckCodeDesc
	r.Set("xt_check_code_desc", _xtCheckCodeDesc)
	return nil
}

// GetXtCheckCodeDesc XtCheckCodeDesc Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetXtCheckCodeDesc() string {
	return r._xtCheckCodeDesc
}

// SetDrugListJson is DrugListJson Setter
// 药品列表【可不填写，不填写时会通过码进行计算】Json："codeCount":         药品数量 "commDrugId":     国家药品唯一标识 "exprieDate":         生产日期 "physicInfo":          药品信息 "pkgSpec":           包状规格 "prepnCount":       制剂数量 "produceBatchNo":生产批次 "produceDate":      生产日期
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetDrugListJson(_drugListJson string) error {
	r._drugListJson = _drugListJson
	r.Set("drug_list_json", _drugListJson)
	return nil
}

// GetDrugListJson DrugListJson Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetDrugListJson() string {
	return r._drugListJson
}

// SetAssRefEntId is AssRefEntId Setter
// 单据委托企业【填写 ref_ent_Id，与ass_ent_id属性选填一个就行】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetAssRefEntId(_assRefEntId string) error {
	r._assRefEntId = _assRefEntId
	r.Set("ass_ref_ent_id", _assRefEntId)
	return nil
}

// GetAssRefEntId AssRefEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetAssRefEntId() string {
	return r._assRefEntId
}

// SetAssEntId is AssEntId Setter
// 单据委托企业【填写 ent_Id】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetAssEntId(_assEntId string) error {
	r._assEntId = _assEntId
	r.Set("ass_ent_id", _assEntId)
	return nil
}

// GetAssEntId AssEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetAssEntId() string {
	return r._assEntId
}

// SetBillType is BillType Setter
// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetBillType() int64 {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药【可以随便填写，单据上传后会以实际为准】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetPhysicType() int64 {
	return r._physicType
}

// SetQuReceivable is QuReceivable Setter
// 应收货总数量
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetQuReceivable(_quReceivable int64) error {
	r._quReceivable = _quReceivable
	r.Set("qu_receivable", _quReceivable)
	return nil
}

// GetQuReceivable QuReceivable Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetQuReceivable() int64 {
	return r._quReceivable
}
