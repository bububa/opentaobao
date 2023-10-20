package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest 企业上传出入库信息 API请求
// alibaba.alihealth.drug.kyt.wes.uploadinoutbill
//
// 企业上传出入库信息，包括101, &#34;生产入库&#34;；102, &#34;采购入库&#34;；103, &#34;退货入库&#34;；104, &#34;调拨入库&#34;；106, &#34;零头入库&#34;；107, &#34;供应入库&#34;；108, &#34;召回入库&#34;；110,&#34;赠品入库&#34;；111,&#34;盘盈入库&#34;；112,&#34;报废入库&#34;；113,&#34;其他入库&#34;
// 201, &#34;销售出库&#34;；202, &#34;退货出库&#34;；203, &#34;调拨出库&#34;；204, &#34;返工出库&#34;；205, &#34;销毁出库&#34;；206, &#34;抽检出库&#34;；207, &#34;直调出库&#34;；208, &#34;生产出库&#34;；209, &#34;供应出库&#34;；211, &#34;召回出库&#34;；212,&#34;赠品出库&#34;；214,&#34;盘亏出库&#34;；215,&#34;损坏出库&#34;；216,&#34;报废出库&#34;；217,&#34;其他出库&#34;；237, &#34;直调退货&#34;。
// 不包括对个人的零售出库，疫苗接种，领药出库。
// 本接口与uploadcircubill接口的主要区别的，本接口入参中直接上传追溯码（多个码时用逗号分隔）。uploadcircubill接口入参中，需要上传码的单据文件（用扫码枪生成的xml文件），一般情况下使用uploadcircubill接口上传单据文件。
type AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest struct {
	model.Params
	// 货主（单据的所有者，上传人），上传企业的单位维一编码【注意：该入参是ref_ent_id，不是ent_id】
	_refEntId string
	// 服务时间
	_licenseToken string
	// 单据编号【同一个企业不能上传相同单据号】
	_billCode string
	// 单据时间（扫码时间）
	_billTime string
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
	// 追溯码【多个码时用逗号拼接的字符串。要求数量在3500个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_traceCodes string
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
	// 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
	_agentRefUserId int64
	// 应收货总数量
	_quReceivable int64
}

// NewAlibabaAlihealthDrugKytWesUploadinoutbillRequest 初始化AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest对象
func NewAlibabaAlihealthDrugKytWesUploadinoutbillRequest() *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest {
	return &AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest{
		Params: model.NewParams(39),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) Reset() {
	r._refEntId = ""
	r._licenseToken = ""
	r._billCode = ""
	r._billTime = ""
	r._fromUserId = ""
	r._toUserId = ""
	r._destUserId = ""
	r._operIcCode = ""
	r._operIcName = ""
	r._warehouseId = ""
	r._drugId = ""
	r._traceCodes = ""
	r._clientType = ""
	r._returnReasonCode = ""
	r._returnReasonDes = ""
	r._cancelReasonCode = ""
	r._cancelReasonDes = ""
	r._executerName = ""
	r._executerCode = ""
	r._superviserName = ""
	r._superviserCode = ""
	r._fromAddress = ""
	r._toAddress = ""
	r._fromBillCode = ""
	r._orderCode = ""
	r._fromPerson = ""
	r._toPerson = ""
	r._disRefEntId = ""
	r._disEntId = ""
	r._xtIsCheck = ""
	r._xtCheckCode = ""
	r._xtCheckCodeDesc = ""
	r._drugListJson = ""
	r._assRefEntId = ""
	r._assEntId = ""
	r._billType = 0
	r._physicType = 0
	r._agentRefUserId = 0
	r._quReceivable = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.uploadinoutbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 货主（单据的所有者，上传人），上传企业的单位维一编码【注意：该入参是ref_ent_id，不是ent_id】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetBillCode is BillCode Setter
// 单据编号【同一个企业不能上传相同单据号】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据时间（扫码时间）
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetFromUserId is FromUserId Setter
// 发货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetToUserId is ToUserId Setter
// 收货企业entId【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// GetToUserId ToUserId Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetToUserId() string {
	return r._toUserId
}

// SetDestUserId is DestUserId Setter
// 直调企业标识【注意：该入参是ent_id,并不是ref_ent_id】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetDestUserId(_destUserId string) error {
	r._destUserId = _destUserId
	r.Set("dest_user_id", _destUserId)
	return nil
}

// GetDestUserId DestUserId Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetDestUserId() string {
	return r._destUserId
}

// SetOperIcCode is OperIcCode Setter
// 单据提交者（调用接口时的appkey编号）
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// GetOperIcCode OperIcCode Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// SetOperIcName is OperIcName Setter
// 单据提交者姓名（出入库单上传人的名子）
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// GetOperIcName OperIcName Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// SetWarehouseId is WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetWarehouseId(_warehouseId string) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetWarehouseId() string {
	return r._warehouseId
}

// SetDrugId is DrugId Setter
// 药品ID[企业自已系统的药品ID]
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// GetDrugId DrugId Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetDrugId() string {
	return r._drugId
}

// SetTraceCodes is TraceCodes Setter
// 追溯码【多个码时用逗号拼接的字符串。要求数量在3500个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetTraceCodes(_traceCodes string) error {
	r._traceCodes = _traceCodes
	r.Set("trace_codes", _traceCodes)
	return nil
}

// GetTraceCodes TraceCodes Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetTraceCodes() string {
	return r._traceCodes
}

// SetClientType is ClientType Setter
// 调用方类型[必须填2]
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetClientType() string {
	return r._clientType
}

// SetReturnReasonCode is ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetReturnReasonCode(_returnReasonCode string) error {
	r._returnReasonCode = _returnReasonCode
	r.Set("return_reason_code", _returnReasonCode)
	return nil
}

// GetReturnReasonCode ReturnReasonCode Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetReturnReasonCode() string {
	return r._returnReasonCode
}

// SetReturnReasonDes is ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetReturnReasonDes(_returnReasonDes string) error {
	r._returnReasonDes = _returnReasonDes
	r.Set("return_reason_des", _returnReasonDes)
	return nil
}

// GetReturnReasonDes ReturnReasonDes Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetReturnReasonDes() string {
	return r._returnReasonDes
}

// SetCancelReasonCode is CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetCancelReasonCode(_cancelReasonCode string) error {
	r._cancelReasonCode = _cancelReasonCode
	r.Set("cancel_reason_code", _cancelReasonCode)
	return nil
}

// GetCancelReasonCode CancelReasonCode Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetCancelReasonCode() string {
	return r._cancelReasonCode
}

// SetCancelReasonDes is CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetCancelReasonDes(_cancelReasonDes string) error {
	r._cancelReasonDes = _cancelReasonDes
	r.Set("cancel_reason_des", _cancelReasonDes)
	return nil
}

// GetCancelReasonDes CancelReasonDes Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetCancelReasonDes() string {
	return r._cancelReasonDes
}

// SetExecuterName is ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetExecuterName(_executerName string) error {
	r._executerName = _executerName
	r.Set("executer_name", _executerName)
	return nil
}

// GetExecuterName ExecuterName Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetExecuterName() string {
	return r._executerName
}

// SetExecuterCode is ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetExecuterCode(_executerCode string) error {
	r._executerCode = _executerCode
	r.Set("executer_code", _executerCode)
	return nil
}

// GetExecuterCode ExecuterCode Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetExecuterCode() string {
	return r._executerCode
}

// SetSuperviserName is SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetSuperviserName(_superviserName string) error {
	r._superviserName = _superviserName
	r.Set("superviser_name", _superviserName)
	return nil
}

// GetSuperviserName SuperviserName Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetSuperviserName() string {
	return r._superviserName
}

// SetSuperviserCode is SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetSuperviserCode(_superviserCode string) error {
	r._superviserCode = _superviserCode
	r.Set("superviser_code", _superviserCode)
	return nil
}

// GetSuperviserCode SuperviserCode Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetSuperviserCode() string {
	return r._superviserCode
}

// SetFromAddress is FromAddress Setter
// 发货地址
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetFromAddress(_fromAddress string) error {
	r._fromAddress = _fromAddress
	r.Set("from_address", _fromAddress)
	return nil
}

// GetFromAddress FromAddress Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetFromAddress() string {
	return r._fromAddress
}

// SetToAddress is ToAddress Setter
// 收货地址
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetToAddress(_toAddress string) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// GetToAddress ToAddress Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetToAddress() string {
	return r._toAddress
}

// SetFromBillCode is FromBillCode Setter
// 发货单编号
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetFromBillCode(_fromBillCode string) error {
	r._fromBillCode = _fromBillCode
	r.Set("from_bill_code", _fromBillCode)
	return nil
}

// GetFromBillCode FromBillCode Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetFromBillCode() string {
	return r._fromBillCode
}

// SetOrderCode is OrderCode Setter
// 订货单编号【可选】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetFromPerson is FromPerson Setter
// 发货人
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetFromPerson(_fromPerson string) error {
	r._fromPerson = _fromPerson
	r.Set("from_person", _fromPerson)
	return nil
}

// GetFromPerson FromPerson Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetFromPerson() string {
	return r._fromPerson
}

// SetToPerson is ToPerson Setter
// 收货人
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetToPerson(_toPerson string) error {
	r._toPerson = _toPerson
	r.Set("to_person", _toPerson)
	return nil
}

// GetToPerson ToPerson Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetToPerson() string {
	return r._toPerson
}

// SetDisRefEntId is DisRefEntId Setter
// 药品配送ref_enti_d企业
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetDisRefEntId(_disRefEntId string) error {
	r._disRefEntId = _disRefEntId
	r.Set("dis_ref_ent_id", _disRefEntId)
	return nil
}

// GetDisRefEntId DisRefEntId Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetDisRefEntId() string {
	return r._disRefEntId
}

// SetDisEntId is DisEntId Setter
// 药品配送企业ent_Id出库单
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetDisEntId(_disEntId string) error {
	r._disEntId = _disEntId
	r.Set("dis_ent_id", _disEntId)
	return nil
}

// GetDisEntId DisEntId Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetDisEntId() string {
	return r._disEntId
}

// SetXtIsCheck is XtIsCheck Setter
// 是否验证，0：未通过验证，1：已验证
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetXtIsCheck(_xtIsCheck string) error {
	r._xtIsCheck = _xtIsCheck
	r.Set("xt_is_check", _xtIsCheck)
	return nil
}

// GetXtIsCheck XtIsCheck Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetXtIsCheck() string {
	return r._xtIsCheck
}

// SetXtCheckCode is XtCheckCode Setter
// 未验证通过原因【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetXtCheckCode(_xtCheckCode string) error {
	r._xtCheckCode = _xtCheckCode
	r.Set("xt_check_code", _xtCheckCode)
	return nil
}

// GetXtCheckCode XtCheckCode Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetXtCheckCode() string {
	return r._xtCheckCode
}

// SetXtCheckCodeDesc is XtCheckCodeDesc Setter
// 未验证通过原因描述【验证未通过时填写】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetXtCheckCodeDesc(_xtCheckCodeDesc string) error {
	r._xtCheckCodeDesc = _xtCheckCodeDesc
	r.Set("xt_check_code_desc", _xtCheckCodeDesc)
	return nil
}

// GetXtCheckCodeDesc XtCheckCodeDesc Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetXtCheckCodeDesc() string {
	return r._xtCheckCodeDesc
}

// SetDrugListJson is DrugListJson Setter
// 药品列表Json：&#34;codeCount&#34;: 药品数量 &#34;commDrugId&#34;: 国家药品唯一标识 &#34;exprieDate&#34;: 生产日期 &#34;physicInfo&#34;: 药品信息 &#34;pkgSpec&#34;: 包状规格 &#34;prepnCount&#34;: 制剂数量 &#34;produceBatchNo&#34;:生产批次 &#34;produceDate&#34;: 生产日期
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetDrugListJson(_drugListJson string) error {
	r._drugListJson = _drugListJson
	r.Set("drug_list_json", _drugListJson)
	return nil
}

// GetDrugListJson DrugListJson Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetDrugListJson() string {
	return r._drugListJson
}

// SetAssRefEntId is AssRefEntId Setter
// 单据委托企业refEntId
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetAssRefEntId(_assRefEntId string) error {
	r._assRefEntId = _assRefEntId
	r.Set("ass_ref_ent_id", _assRefEntId)
	return nil
}

// GetAssRefEntId AssRefEntId Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetAssRefEntId() string {
	return r._assRefEntId
}

// SetAssEntId is AssEntId Setter
// 单据委托企业entId
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetAssEntId(_assEntId string) error {
	r._assEntId = _assEntId
	r.Set("ass_ent_id", _assEntId)
	return nil
}

// GetAssEntId AssEntId Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetAssEntId() string {
	return r._assEntId
}

// SetBillType is BillType Setter
// 单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetBillType() int64 {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型【3普药2特药】89开头的码定义为特药，其它码定义成普药【可以随便填写，单据上传后会以实际为准】
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetPhysicType() int64 {
	return r._physicType
}

// SetAgentRefUserId is AgentRefUserId Setter
// 第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetAgentRefUserId(_agentRefUserId int64) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// GetAgentRefUserId AgentRefUserId Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetAgentRefUserId() int64 {
	return r._agentRefUserId
}

// SetQuReceivable is QuReceivable Setter
// 应收货总数量
func (r *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) SetQuReceivable(_quReceivable int64) error {
	r._quReceivable = _quReceivable
	r.Set("qu_receivable", _quReceivable)
	return nil
}

// GetQuReceivable QuReceivable Getter
func (r AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) GetQuReceivable() int64 {
	return r._quReceivable
}

var poolAlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytWesUploadinoutbillRequest()
	},
}

// GetAlibabaAlihealthDrugKytWesUploadinoutbillRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest
func GetAlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest() *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest {
	return poolAlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest.Get().(*AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest 将 AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest(v *AlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesUploadinoutbillAPIRequest.Put(v)
}
