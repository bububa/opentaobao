package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUploadcircubillAPIRequest
生产批发单据上传 API请求
alibaba.alihealth.drug.kyt.uploadcircubill

生产批发单据上传（非零售企业使用），包括101, "生产入库"；102, "采购入库"；103, "退货入库"；104, "调拨入库"；106, "零头入库"；107, "供应入库"；108, "召回入库"；110,"赠品入库"；111,"盘盈入库"；112,"报废入库"；113,"其他入库"
201, "销售出库"；202, "退货出库"；203, "调拨出库"；204, "返工出库"；205, "销毁出库"；206, "抽检出库"；207, "直调出库"；208, "生产出库"；209, "供应出库"；211, "召回出库"；212,"赠品出库"；214,"盘亏出库"；215,"损坏出库"；216,"报废出库"；217,"其他出库"；237, "直调退货"。 */
type AlibabaAlihealthDrugKytUploadcircubillAPIRequest struct {
	model.Params
	// 单据编号【同一个企业不能上传相同单据号】
	_billCode string
	// 单据时间（扫码时间）
	_billTime string
	// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
	_billType int64
	// 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药
	_physicType int64
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
	// 单据名称
	_uploadFileName string
	// 客户端类型【暂定都写2】
	_clientType string
	// （协同平台数据合规）应收货总数量
	_quReceivable int64
	// （协同平台数据合规）是否验证，0：未通过验证，1：已验证
	_xtIsCheck string
	// （协同平台数据合规）未验证通过原因【验证未通过时填写】
	_xtCheckCode string
	// （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
	_xtCheckCodeDesc string
	// （协同平台数据合规）药品列表Json："codeCount":         药品数量 "commDrugId":     国家药品唯一标识 "exprieDate":         生产日期 "physicInfo":          药品信息 "pkgSpec":           包状规格 "prepnCount":       制剂数量 "produceBatchNo":生产批次 "produceDate":      生产日期
	_drugListJson string
	// （协同平台数据合规）单据委托企业refEntId
	_assRefEntId string
	// （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
	_disEntId string
	// （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
	_disRefEntId string
	// （协同平台数据合规）收货人【必选】
	_toPerson string
	// （协同平台数据合规）发货人【必选】
	_fromPerson string
	// （协同平台数据合规）订货单编号【可选】
	_orderCode string
	// （协同平台数据合规）发货单编号【必选】
	_fromBillCode string
	// （协同平台数据合规）收货地址【必选】
	_toAddress string
	// （协同平台数据合规）发货地址【必选】
	_fromAddress string
	// （协同平台数据合规）单据委托企业entId
	_assEntId string
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

// Set is BillCode Setter
// 单据编号【同一个企业不能上传相同单据号】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is BillTime Setter
// 单据时间（扫码时间）
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// Get BillTime Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetBillTime() string {
	return r._billTime
}

// Set is BillType Setter
// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// Get BillType Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetBillType() int64 {
	return r._billType
}

// Set is PhysicType Setter
// 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// Get PhysicType Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetPhysicType() int64 {
	return r._physicType
}

// Set is RefUserId Setter
// 货主企业（单据的所有者，上传人）【注意：该入参是ref_ent_id，不是ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// Get RefUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// Set is AgentRefUserId Setter
// 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// Get AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// Set is FromUserId Setter
// 发货企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// Get FromUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// Set is ToUserId Setter
// 收货企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// Get ToUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetToUserId() string {
	return r._toUserId
}

// Set is DestUserId Setter
// 直调企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetDestUserId(_destUserId string) error {
	r._destUserId = _destUserId
	r.Set("dest_user_id", _destUserId)
	return nil
}

// Get DestUserId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetDestUserId() string {
	return r._destUserId
}

// Set is OperIcCode Setter
// 操作人标识（appkey编号）
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// Get OperIcCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// Set is OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// Get OperIcName Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// Set is FileContent Setter
// 单据文件体【bas64字符串】，看对接文档中的代码示例，示例中有相应说明。
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetFileContent(_fileContent string) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// Get FileContent Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetFileContent() string {
	return r._fileContent
}

// Set is UploadFileName Setter
// 单据名称
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetUploadFileName(_uploadFileName string) error {
	r._uploadFileName = _uploadFileName
	r.Set("upload_file_name", _uploadFileName)
	return nil
}

// Get UploadFileName Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetUploadFileName() string {
	return r._uploadFileName
}

// Set is ClientType Setter
// 客户端类型【暂定都写2】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// Get ClientType Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetClientType() string {
	return r._clientType
}

// Set is QuReceivable Setter
// （协同平台数据合规）应收货总数量
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetQuReceivable(_quReceivable int64) error {
	r._quReceivable = _quReceivable
	r.Set("qu_receivable", _quReceivable)
	return nil
}

// Get QuReceivable Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetQuReceivable() int64 {
	return r._quReceivable
}

// Set is XtIsCheck Setter
// （协同平台数据合规）是否验证，0：未通过验证，1：已验证
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetXtIsCheck(_xtIsCheck string) error {
	r._xtIsCheck = _xtIsCheck
	r.Set("xt_is_check", _xtIsCheck)
	return nil
}

// Get XtIsCheck Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetXtIsCheck() string {
	return r._xtIsCheck
}

// Set is XtCheckCode Setter
// （协同平台数据合规）未验证通过原因【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetXtCheckCode(_xtCheckCode string) error {
	r._xtCheckCode = _xtCheckCode
	r.Set("xt_check_code", _xtCheckCode)
	return nil
}

// Get XtCheckCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetXtCheckCode() string {
	return r._xtCheckCode
}

// Set is XtCheckCodeDesc Setter
// （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetXtCheckCodeDesc(_xtCheckCodeDesc string) error {
	r._xtCheckCodeDesc = _xtCheckCodeDesc
	r.Set("xt_check_code_desc", _xtCheckCodeDesc)
	return nil
}

// Get XtCheckCodeDesc Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetXtCheckCodeDesc() string {
	return r._xtCheckCodeDesc
}

// Set is DrugListJson Setter
// （协同平台数据合规）药品列表Json："codeCount":         药品数量 "commDrugId":     国家药品唯一标识 "exprieDate":         生产日期 "physicInfo":          药品信息 "pkgSpec":           包状规格 "prepnCount":       制剂数量 "produceBatchNo":生产批次 "produceDate":      生产日期
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetDrugListJson(_drugListJson string) error {
	r._drugListJson = _drugListJson
	r.Set("drug_list_json", _drugListJson)
	return nil
}

// Get DrugListJson Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetDrugListJson() string {
	return r._drugListJson
}

// Set is AssRefEntId Setter
// （协同平台数据合规）单据委托企业refEntId
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetAssRefEntId(_assRefEntId string) error {
	r._assRefEntId = _assRefEntId
	r.Set("ass_ref_ent_id", _assRefEntId)
	return nil
}

// Get AssRefEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetAssRefEntId() string {
	return r._assRefEntId
}

// Set is DisEntId Setter
// （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetDisEntId(_disEntId string) error {
	r._disEntId = _disEntId
	r.Set("dis_ent_id", _disEntId)
	return nil
}

// Get DisEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetDisEntId() string {
	return r._disEntId
}

// Set is DisRefEntId Setter
// （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetDisRefEntId(_disRefEntId string) error {
	r._disRefEntId = _disRefEntId
	r.Set("dis_ref_ent_id", _disRefEntId)
	return nil
}

// Get DisRefEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetDisRefEntId() string {
	return r._disRefEntId
}

// Set is ToPerson Setter
// （协同平台数据合规）收货人【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetToPerson(_toPerson string) error {
	r._toPerson = _toPerson
	r.Set("to_person", _toPerson)
	return nil
}

// Get ToPerson Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetToPerson() string {
	return r._toPerson
}

// Set is FromPerson Setter
// （协同平台数据合规）发货人【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetFromPerson(_fromPerson string) error {
	r._fromPerson = _fromPerson
	r.Set("from_person", _fromPerson)
	return nil
}

// Get FromPerson Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetFromPerson() string {
	return r._fromPerson
}

// Set is OrderCode Setter
// （协同平台数据合规）订货单编号【可选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// Get OrderCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// Set is FromBillCode Setter
// （协同平台数据合规）发货单编号【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetFromBillCode(_fromBillCode string) error {
	r._fromBillCode = _fromBillCode
	r.Set("from_bill_code", _fromBillCode)
	return nil
}

// Get FromBillCode Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetFromBillCode() string {
	return r._fromBillCode
}

// Set is ToAddress Setter
// （协同平台数据合规）收货地址【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetToAddress(_toAddress string) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// Get ToAddress Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetToAddress() string {
	return r._toAddress
}

// Set is FromAddress Setter
// （协同平台数据合规）发货地址【必选】
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetFromAddress(_fromAddress string) error {
	r._fromAddress = _fromAddress
	r.Set("from_address", _fromAddress)
	return nil
}

// Get FromAddress Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetFromAddress() string {
	return r._fromAddress
}

// Set is AssEntId Setter
// （协同平台数据合规）单据委托企业entId
func (r *AlibabaAlihealthDrugKytUploadcircubillAPIRequest) SetAssEntId(_assEntId string) error {
	r._assEntId = _assEntId
	r.Set("ass_ent_id", _assEntId)
	return nil
}

// Get AssEntId Getter
func (r AlibabaAlihealthDrugKytUploadcircubillAPIRequest) GetAssEntId() string {
	return r._assEntId
}
