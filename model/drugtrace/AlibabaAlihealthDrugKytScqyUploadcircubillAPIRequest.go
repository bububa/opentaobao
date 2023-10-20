package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest 生成批发单据上传 API请求
// alibaba.alihealth.drug.kyt.scqy.uploadcircubill
//
// 生产批发单据上传（非零售企业使用），包括101, &#34;生产入库&#34;；102, &#34;采购入库&#34;；103, &#34;退货入库&#34;；104, &#34;调拨入库&#34;；106, &#34;零头入库&#34;；107, &#34;供应入库&#34;；108, &#34;召回入库&#34;；110,&#34;赠品入库&#34;；111,&#34;盘盈入库&#34;；112,&#34;报废入库&#34;；113,&#34;其他入库&#34;
// 201, &#34;销售出库&#34;；202, &#34;退货出库&#34;；203, &#34;调拨出库&#34;；204, &#34;返工出库&#34;；205, &#34;销毁出库&#34;；206, &#34;抽检出库&#34;；207, &#34;直调出库&#34;；208, &#34;生产出库&#34;；209, &#34;供应出库&#34;；211, &#34;召回出库&#34;；212,&#34;赠品出库&#34;；214,&#34;盘亏出库&#34;；215,&#34;损坏出库&#34;；216,&#34;报废出库&#34;；217,&#34;其他出库&#34;；237, &#34;直调退货&#34;。
type AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest struct {
	model.Params
	// 单据编号【同一个企业不能上传相同单据号】
	_billCode string
	// 单据时间（扫码时间）
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
	// 操作人标识（写appkey编号）
	_operIcCode string
	// 单据提交者姓名
	_operIcName string
	// 单据文件体【bas64字符串】，看对接文档中的代码示例，示例中有相应说明。
	_fileContent string
	// 单据名称
	_uploadFileName string
	// 请求端类型【暂定都写2】
	_clientType string
	// 发货地址
	_fromAddress string
	// 收货地址
	_toAddress string
	// 发货单编号
	_fromBillCode string
	// 订货单编号
	_orderCode string
	// 发货人(特药出单据必填)
	_fromPerson string
	// 收货人(特药入单据必填)
	_toPerson string
	// 药品配送企业【出库单填写】
	_disRefEntId string
	// 药品配送企业entId【出库单填写】
	_disEntId string
	// 是否验证，0：未通过验证，1：已验证
	_xtIsCheck string
	// 未验证通过原因【验证未通过时填写】
	_xtCheckCode string
	// 未验证通过原因描述【验证未通过时填写】
	_xtCheckCodeDesc string
	// 【可不填】药品列表Json："codeCount":         药品数量 "commDrugId":     国家药品唯一标识 "exprieDate":         生产日期 "physicInfo":          药品信息 "pkgSpec":           包状规格 "prepnCount":       制剂数量 "produceBatchNo":生产批次 "produceDate":      生产日期
	_drugListJson string
	// 单据委托企业refEntId
	_assRefEntId string
	// 单据委托企业entId
	_assEntId string
	// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
	_billType int64
	// 药品类型【3普药2特药】一般情况下，89开头的码定义为特药，其它码定义成普药
	_physicType int64
	// 应收货总数量
	_quReceivable int64
}

// NewAlibabaalihealthdrugkytscqyuploadcircubillRequest 初始化AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest对象
func NewAlibabaalihealthdrugkytscqyuploadcircubillRequest() *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest {
	return &AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.uploadcircubill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillCode is BillCode Setter
// 单据编号【同一个企业不能上传相同单据号】
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据时间（扫码时间）
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetRefUserId is RefUserId Setter
// 货主企业（单据的所有者，上传人）【注意：该入参是ref_ent_id，不是ent_id】
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetAgentRefUserId is AgentRefUserId Setter
// 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// GetAgentRefUserId AgentRefUserId Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// SetFromUserId is FromUserId Setter
// 发货企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetToUserId is ToUserId Setter
// 收货企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// GetToUserId ToUserId Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetToUserId() string {
	return r._toUserId
}

// SetDestUserId is DestUserId Setter
// 直调企业【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetDestUserId(_destUserId string) error {
	r._destUserId = _destUserId
	r.Set("dest_user_id", _destUserId)
	return nil
}

// GetDestUserId DestUserId Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetDestUserId() string {
	return r._destUserId
}

// SetOperIcCode is OperIcCode Setter
// 操作人标识（写appkey编号）
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// GetOperIcCode OperIcCode Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// SetOperIcName is OperIcName Setter
// 单据提交者姓名
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// GetOperIcName OperIcName Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// SetFileContent is FileContent Setter
// 单据文件体【bas64字符串】，看对接文档中的代码示例，示例中有相应说明。
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetFileContent(_fileContent string) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// GetFileContent FileContent Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetFileContent() string {
	return r._fileContent
}

// SetUploadFileName is UploadFileName Setter
// 单据名称
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetUploadFileName(_uploadFileName string) error {
	r._uploadFileName = _uploadFileName
	r.Set("upload_file_name", _uploadFileName)
	return nil
}

// GetUploadFileName UploadFileName Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetUploadFileName() string {
	return r._uploadFileName
}

// SetClientType is ClientType Setter
// 请求端类型【暂定都写2】
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetClientType() string {
	return r._clientType
}

// SetFromAddress is FromAddress Setter
// 发货地址
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetFromAddress(_fromAddress string) error {
	r._fromAddress = _fromAddress
	r.Set("from_address", _fromAddress)
	return nil
}

// GetFromAddress FromAddress Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetFromAddress() string {
	return r._fromAddress
}

// SetToAddress is ToAddress Setter
// 收货地址
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetToAddress(_toAddress string) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// GetToAddress ToAddress Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetToAddress() string {
	return r._toAddress
}

// SetFromBillCode is FromBillCode Setter
// 发货单编号
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetFromBillCode(_fromBillCode string) error {
	r._fromBillCode = _fromBillCode
	r.Set("from_bill_code", _fromBillCode)
	return nil
}

// GetFromBillCode FromBillCode Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetFromBillCode() string {
	return r._fromBillCode
}

// SetOrderCode is OrderCode Setter
// 订货单编号
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetFromPerson is FromPerson Setter
// 发货人(特药出单据必填)
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetFromPerson(_fromPerson string) error {
	r._fromPerson = _fromPerson
	r.Set("from_person", _fromPerson)
	return nil
}

// GetFromPerson FromPerson Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetFromPerson() string {
	return r._fromPerson
}

// SetToPerson is ToPerson Setter
// 收货人(特药入单据必填)
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetToPerson(_toPerson string) error {
	r._toPerson = _toPerson
	r.Set("to_person", _toPerson)
	return nil
}

// GetToPerson ToPerson Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetToPerson() string {
	return r._toPerson
}

// SetDisRefEntId is DisRefEntId Setter
// 药品配送企业【出库单填写】
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetDisRefEntId(_disRefEntId string) error {
	r._disRefEntId = _disRefEntId
	r.Set("dis_ref_ent_id", _disRefEntId)
	return nil
}

// GetDisRefEntId DisRefEntId Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetDisRefEntId() string {
	return r._disRefEntId
}

// SetDisEntId is DisEntId Setter
// 药品配送企业entId【出库单填写】
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetDisEntId(_disEntId string) error {
	r._disEntId = _disEntId
	r.Set("dis_ent_id", _disEntId)
	return nil
}

// GetDisEntId DisEntId Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetDisEntId() string {
	return r._disEntId
}

// SetXtIsCheck is XtIsCheck Setter
// 是否验证，0：未通过验证，1：已验证
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetXtIsCheck(_xtIsCheck string) error {
	r._xtIsCheck = _xtIsCheck
	r.Set("xt_is_check", _xtIsCheck)
	return nil
}

// GetXtIsCheck XtIsCheck Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetXtIsCheck() string {
	return r._xtIsCheck
}

// SetXtCheckCode is XtCheckCode Setter
// 未验证通过原因【验证未通过时填写】
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetXtCheckCode(_xtCheckCode string) error {
	r._xtCheckCode = _xtCheckCode
	r.Set("xt_check_code", _xtCheckCode)
	return nil
}

// GetXtCheckCode XtCheckCode Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetXtCheckCode() string {
	return r._xtCheckCode
}

// SetXtCheckCodeDesc is XtCheckCodeDesc Setter
// 未验证通过原因描述【验证未通过时填写】
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetXtCheckCodeDesc(_xtCheckCodeDesc string) error {
	r._xtCheckCodeDesc = _xtCheckCodeDesc
	r.Set("xt_check_code_desc", _xtCheckCodeDesc)
	return nil
}

// GetXtCheckCodeDesc XtCheckCodeDesc Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetXtCheckCodeDesc() string {
	return r._xtCheckCodeDesc
}

// SetDrugListJson is DrugListJson Setter
// 【可不填】药品列表Json：&#34;codeCount&#34;:         药品数量 &#34;commDrugId&#34;:     国家药品唯一标识 &#34;exprieDate&#34;:         生产日期 &#34;physicInfo&#34;:          药品信息 &#34;pkgSpec&#34;:           包状规格 &#34;prepnCount&#34;:       制剂数量 &#34;produceBatchNo&#34;:生产批次 &#34;produceDate&#34;:      生产日期
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetDrugListJson(_drugListJson string) error {
	r._drugListJson = _drugListJson
	r.Set("drug_list_json", _drugListJson)
	return nil
}

// GetDrugListJson DrugListJson Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetDrugListJson() string {
	return r._drugListJson
}

// SetAssRefEntId is AssRefEntId Setter
// 单据委托企业refEntId
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetAssRefEntId(_assRefEntId string) error {
	r._assRefEntId = _assRefEntId
	r.Set("ass_ref_ent_id", _assRefEntId)
	return nil
}

// GetAssRefEntId AssRefEntId Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetAssRefEntId() string {
	return r._assRefEntId
}

// SetAssEntId is AssEntId Setter
// 单据委托企业entId
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetAssEntId(_assEntId string) error {
	r._assEntId = _assEntId
	r.Set("ass_ent_id", _assEntId)
	return nil
}

// GetAssEntId AssEntId Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetAssEntId() string {
	return r._assEntId
}

// SetBillType is BillType Setter
// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetBillType() int64 {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型【3普药2特药】一般情况下，89开头的码定义为特药，其它码定义成普药
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetPhysicType() int64 {
	return r._physicType
}

// SetQuReceivable is QuReceivable Setter
// 应收货总数量
func (r *AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) SetQuReceivable(_quReceivable int64) error {
	r._quReceivable = _quReceivable
	r.Set("qu_receivable", _quReceivable)
	return nil
}

// GetQuReceivable QuReceivable Getter
func (r AlibabaalihealthdrugkytscqyuploadcircubillAPIRequest) GetQuReceivable() int64 {
	return r._quReceivable
}
