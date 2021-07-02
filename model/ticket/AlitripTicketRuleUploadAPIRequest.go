package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTicketRuleUploadAPIRequest 【门票API2.0】景点门票规则维护接口 API请求
// alitrip.ticket.rule.upload
//
// 景点门票规则维护接口。该接口同时支持新发规则和编辑现有规则，如果out_rule_id下没有发布过规则，则系统将判断为新发一个规则，否则认为是编辑现有规则。
// 对于新发布规则的情况，有些参数是必填的，请仔细查看各字段说明。对于编辑的情况，除out_rule_id外都是可选，编辑情况支持增量更新（某个参数不传则使用该规则上原有值）
type AlitripTicketRuleUploadAPIRequest struct {
	model.Params
	// 商户票种规则id
	_outRuleId string
	// 新发布规则时必填。商户票种规则名称
	_outRuleName string
	// 商户景点编码。新发布规则时必填，out_scenic_id与ali_scenic_id二选一，至少填写其中一个
	_outScenicId string
	// 阿里旅行景点编码。新发布规则时必填，out_scenic_id与ali_scenic_id二选一，至少填写其中一个
	_aliScenicId int64
	// 可选，票种规则类型：0-实体票规则，1-电子票（包含手工票和直连票）规则。新发规则时不传 默认为1
	_ruleType int64
	// 可选，规则状态。0-生效，-1-失效。新发规则时不传 默认生效状态。
	_ruleStatus int64
	// 新发布规则时必填。退票类型。1-无条件退改， 2-有条件退改， 3-不可退改。
	_refundType int64
	// （该字段已废弃，请使用结构化退改规则字段替代：refund_custom_rules）新发布规则时特殊选填。退票描述。当refund_type为2时，该字段必填
	_refundDesc string
	// 可选，是否支持自动退款，0或为空时表示不支持。0-不支持，1-只支持"售中自动退款"，2-只支持"过期未使用自动退款"，3-同时支持"售中自动退款"和"过期未使用自动退款"
	_autoRefundSupport int64
	// 结构化自定义退款规则（json数组格式），当refund_type为2时，该字段必填。【type字段说明（特别注意：2和3不能同时存在）：1（游玩日期前退改规则），2（游玩日期当日退改规则），3（区间票 游玩日期有效期内退改规则），4（游玩日期后退改规则），5（其他情况退改规则）。amount字段说明：收取的手续费。unit字段说明：手续费单位，1（票价百分比），2（固定金额，单位分）。】  【示例含义说明：1、游玩日期前2天16点30分前申请退款，每张票收取票价20%的手续费；2、(单日票）游玩日当天12点00分前申请退款，每张票收取票价50%手续费；3、（区间票）游玩日期有效期，最后一天11点00分前申请退款，每张票收取票价60%手续费；4、游玩日期后7天23点59分前申请退款，每张票收取票价80%手续费；5、其他情况，每张票收取固定100元手续费。】
	_refundCustomRules string
	// 新发布规则时必填。出游人信息设置。1、不需要，2、仅需一位游客信息。3、需要所有游客信息。不填默认为1（不需要）。注：实体票（rule_type=2）不需要出游人信息，电子票（rule_type=1）一般是需要出游人信息
	_visitorRequire int64
	// 新发布规则时特殊选填。需要的出游人信息，需要出游人信息时必填。 2:手机号,3:身份证,4:姓名,17：港澳居民居住证 18：台湾居民居住证 19：护照 20：港澳台居民往返内地通行证（回乡证）。示例格式：2,3,4
	_visitorInfos string
	// 新发布规则时必填。出游人 是否限购。1:限购,2:不限购
	_visitorLimitAble int64
	// 新发布规则时特殊选填。限购模式：mode为1按天, mode为2按周, mode为3按月
	_visitorLimitMode int64
	// 新发布规则时特殊选填。限购类型。0-身份证限购， 1-手机号限购
	_visitorLimitType int64
	// 新发布规则时特殊选填。限购数量
	_visitorLimitNum int64
	// 新发布规则时必填。入园类型。1-用兑换凭证直接入园，2-用兑换凭证换票入园
	_enterType int64
	// 新发布规则时必填。入园使用的凭证类型。1、二维码，2、身份证，3、二维码或身份证，4:数字码，5、手机号，6、其它。当enter_type为1时，该字段只能选择1～3 。而当enterType为2时，该字段可设置1～6
	_enterVoucherType int64
	// 其他入园凭证类型。当enter_voucher_type=6时，填写其他入园凭证类型。
	_enterVoucherValue string
	// 新发布规则时特殊选填。换票地址。当enter_type为2时，该字段必填。
	_ticketChangeAdderss string
	// 新发布规则时必填。景区入园地址。游客在景区的详细入园地址，请仔细填写。
	_enterAddress string
	// 新发布规则时必填。门票费用包含，请详细说明该门票商品包含的费用信息。
	_feeInclude string
	// 新发布规则时必填。门票商品一些游客须知的补充说明。1600字符限制
	_extraDesc string
}

// NewAlitripTicketRuleUploadRequest 初始化AlitripTicketRuleUploadAPIRequest对象
func NewAlitripTicketRuleUploadRequest() *AlitripTicketRuleUploadAPIRequest {
	return &AlitripTicketRuleUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTicketRuleUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.ticket.rule.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTicketRuleUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOutRuleId is OutRuleId Setter
// 商户票种规则id
func (r *AlitripTicketRuleUploadAPIRequest) SetOutRuleId(_outRuleId string) error {
	r._outRuleId = _outRuleId
	r.Set("out_rule_id", _outRuleId)
	return nil
}

// GetOutRuleId OutRuleId Getter
func (r AlitripTicketRuleUploadAPIRequest) GetOutRuleId() string {
	return r._outRuleId
}

// SetOutRuleName is OutRuleName Setter
// 新发布规则时必填。商户票种规则名称
func (r *AlitripTicketRuleUploadAPIRequest) SetOutRuleName(_outRuleName string) error {
	r._outRuleName = _outRuleName
	r.Set("out_rule_name", _outRuleName)
	return nil
}

// GetOutRuleName OutRuleName Getter
func (r AlitripTicketRuleUploadAPIRequest) GetOutRuleName() string {
	return r._outRuleName
}

// SetOutScenicId is OutScenicId Setter
// 商户景点编码。新发布规则时必填，out_scenic_id与ali_scenic_id二选一，至少填写其中一个
func (r *AlitripTicketRuleUploadAPIRequest) SetOutScenicId(_outScenicId string) error {
	r._outScenicId = _outScenicId
	r.Set("out_scenic_id", _outScenicId)
	return nil
}

// GetOutScenicId OutScenicId Getter
func (r AlitripTicketRuleUploadAPIRequest) GetOutScenicId() string {
	return r._outScenicId
}

// SetAliScenicId is AliScenicId Setter
// 阿里旅行景点编码。新发布规则时必填，out_scenic_id与ali_scenic_id二选一，至少填写其中一个
func (r *AlitripTicketRuleUploadAPIRequest) SetAliScenicId(_aliScenicId int64) error {
	r._aliScenicId = _aliScenicId
	r.Set("ali_scenic_id", _aliScenicId)
	return nil
}

// GetAliScenicId AliScenicId Getter
func (r AlitripTicketRuleUploadAPIRequest) GetAliScenicId() int64 {
	return r._aliScenicId
}

// SetRuleType is RuleType Setter
// 可选，票种规则类型：0-实体票规则，1-电子票（包含手工票和直连票）规则。新发规则时不传 默认为1
func (r *AlitripTicketRuleUploadAPIRequest) SetRuleType(_ruleType int64) error {
	r._ruleType = _ruleType
	r.Set("rule_type", _ruleType)
	return nil
}

// GetRuleType RuleType Getter
func (r AlitripTicketRuleUploadAPIRequest) GetRuleType() int64 {
	return r._ruleType
}

// SetRuleStatus is RuleStatus Setter
// 可选，规则状态。0-生效，-1-失效。新发规则时不传 默认生效状态。
func (r *AlitripTicketRuleUploadAPIRequest) SetRuleStatus(_ruleStatus int64) error {
	r._ruleStatus = _ruleStatus
	r.Set("rule_status", _ruleStatus)
	return nil
}

// GetRuleStatus RuleStatus Getter
func (r AlitripTicketRuleUploadAPIRequest) GetRuleStatus() int64 {
	return r._ruleStatus
}

// SetRefundType is RefundType Setter
// 新发布规则时必填。退票类型。1-无条件退改， 2-有条件退改， 3-不可退改。
func (r *AlitripTicketRuleUploadAPIRequest) SetRefundType(_refundType int64) error {
	r._refundType = _refundType
	r.Set("refund_type", _refundType)
	return nil
}

// GetRefundType RefundType Getter
func (r AlitripTicketRuleUploadAPIRequest) GetRefundType() int64 {
	return r._refundType
}

// SetRefundDesc is RefundDesc Setter
// （该字段已废弃，请使用结构化退改规则字段替代：refund_custom_rules）新发布规则时特殊选填。退票描述。当refund_type为2时，该字段必填
func (r *AlitripTicketRuleUploadAPIRequest) SetRefundDesc(_refundDesc string) error {
	r._refundDesc = _refundDesc
	r.Set("refund_desc", _refundDesc)
	return nil
}

// GetRefundDesc RefundDesc Getter
func (r AlitripTicketRuleUploadAPIRequest) GetRefundDesc() string {
	return r._refundDesc
}

// SetAutoRefundSupport is AutoRefundSupport Setter
// 可选，是否支持自动退款，0或为空时表示不支持。0-不支持，1-只支持"售中自动退款"，2-只支持"过期未使用自动退款"，3-同时支持"售中自动退款"和"过期未使用自动退款"
func (r *AlitripTicketRuleUploadAPIRequest) SetAutoRefundSupport(_autoRefundSupport int64) error {
	r._autoRefundSupport = _autoRefundSupport
	r.Set("auto_refund_support", _autoRefundSupport)
	return nil
}

// GetAutoRefundSupport AutoRefundSupport Getter
func (r AlitripTicketRuleUploadAPIRequest) GetAutoRefundSupport() int64 {
	return r._autoRefundSupport
}

// SetRefundCustomRules is RefundCustomRules Setter
// 结构化自定义退款规则（json数组格式），当refund_type为2时，该字段必填。【type字段说明（特别注意：2和3不能同时存在）：1（游玩日期前退改规则），2（游玩日期当日退改规则），3（区间票 游玩日期有效期内退改规则），4（游玩日期后退改规则），5（其他情况退改规则）。amount字段说明：收取的手续费。unit字段说明：手续费单位，1（票价百分比），2（固定金额，单位分）。】  【示例含义说明：1、游玩日期前2天16点30分前申请退款，每张票收取票价20%的手续费；2、(单日票）游玩日当天12点00分前申请退款，每张票收取票价50%手续费；3、（区间票）游玩日期有效期，最后一天11点00分前申请退款，每张票收取票价60%手续费；4、游玩日期后7天23点59分前申请退款，每张票收取票价80%手续费；5、其他情况，每张票收取固定100元手续费。】
func (r *AlitripTicketRuleUploadAPIRequest) SetRefundCustomRules(_refundCustomRules string) error {
	r._refundCustomRules = _refundCustomRules
	r.Set("refund_custom_rules", _refundCustomRules)
	return nil
}

// GetRefundCustomRules RefundCustomRules Getter
func (r AlitripTicketRuleUploadAPIRequest) GetRefundCustomRules() string {
	return r._refundCustomRules
}

// SetVisitorRequire is VisitorRequire Setter
// 新发布规则时必填。出游人信息设置。1、不需要，2、仅需一位游客信息。3、需要所有游客信息。不填默认为1（不需要）。注：实体票（rule_type=2）不需要出游人信息，电子票（rule_type=1）一般是需要出游人信息
func (r *AlitripTicketRuleUploadAPIRequest) SetVisitorRequire(_visitorRequire int64) error {
	r._visitorRequire = _visitorRequire
	r.Set("visitor_require", _visitorRequire)
	return nil
}

// GetVisitorRequire VisitorRequire Getter
func (r AlitripTicketRuleUploadAPIRequest) GetVisitorRequire() int64 {
	return r._visitorRequire
}

// SetVisitorInfos is VisitorInfos Setter
// 新发布规则时特殊选填。需要的出游人信息，需要出游人信息时必填。 2:手机号,3:身份证,4:姓名,17：港澳居民居住证 18：台湾居民居住证 19：护照 20：港澳台居民往返内地通行证（回乡证）。示例格式：2,3,4
func (r *AlitripTicketRuleUploadAPIRequest) SetVisitorInfos(_visitorInfos string) error {
	r._visitorInfos = _visitorInfos
	r.Set("visitor_infos", _visitorInfos)
	return nil
}

// GetVisitorInfos VisitorInfos Getter
func (r AlitripTicketRuleUploadAPIRequest) GetVisitorInfos() string {
	return r._visitorInfos
}

// SetVisitorLimitAble is VisitorLimitAble Setter
// 新发布规则时必填。出游人 是否限购。1:限购,2:不限购
func (r *AlitripTicketRuleUploadAPIRequest) SetVisitorLimitAble(_visitorLimitAble int64) error {
	r._visitorLimitAble = _visitorLimitAble
	r.Set("visitor_limit_able", _visitorLimitAble)
	return nil
}

// GetVisitorLimitAble VisitorLimitAble Getter
func (r AlitripTicketRuleUploadAPIRequest) GetVisitorLimitAble() int64 {
	return r._visitorLimitAble
}

// SetVisitorLimitMode is VisitorLimitMode Setter
// 新发布规则时特殊选填。限购模式：mode为1按天, mode为2按周, mode为3按月
func (r *AlitripTicketRuleUploadAPIRequest) SetVisitorLimitMode(_visitorLimitMode int64) error {
	r._visitorLimitMode = _visitorLimitMode
	r.Set("visitor_limit_mode", _visitorLimitMode)
	return nil
}

// GetVisitorLimitMode VisitorLimitMode Getter
func (r AlitripTicketRuleUploadAPIRequest) GetVisitorLimitMode() int64 {
	return r._visitorLimitMode
}

// SetVisitorLimitType is VisitorLimitType Setter
// 新发布规则时特殊选填。限购类型。0-身份证限购， 1-手机号限购
func (r *AlitripTicketRuleUploadAPIRequest) SetVisitorLimitType(_visitorLimitType int64) error {
	r._visitorLimitType = _visitorLimitType
	r.Set("visitor_limit_type", _visitorLimitType)
	return nil
}

// GetVisitorLimitType VisitorLimitType Getter
func (r AlitripTicketRuleUploadAPIRequest) GetVisitorLimitType() int64 {
	return r._visitorLimitType
}

// SetVisitorLimitNum is VisitorLimitNum Setter
// 新发布规则时特殊选填。限购数量
func (r *AlitripTicketRuleUploadAPIRequest) SetVisitorLimitNum(_visitorLimitNum int64) error {
	r._visitorLimitNum = _visitorLimitNum
	r.Set("visitor_limit_num", _visitorLimitNum)
	return nil
}

// GetVisitorLimitNum VisitorLimitNum Getter
func (r AlitripTicketRuleUploadAPIRequest) GetVisitorLimitNum() int64 {
	return r._visitorLimitNum
}

// SetEnterType is EnterType Setter
// 新发布规则时必填。入园类型。1-用兑换凭证直接入园，2-用兑换凭证换票入园
func (r *AlitripTicketRuleUploadAPIRequest) SetEnterType(_enterType int64) error {
	r._enterType = _enterType
	r.Set("enter_type", _enterType)
	return nil
}

// GetEnterType EnterType Getter
func (r AlitripTicketRuleUploadAPIRequest) GetEnterType() int64 {
	return r._enterType
}

// SetEnterVoucherType is EnterVoucherType Setter
// 新发布规则时必填。入园使用的凭证类型。1、二维码，2、身份证，3、二维码或身份证，4:数字码，5、手机号，6、其它。当enter_type为1时，该字段只能选择1～3 。而当enterType为2时，该字段可设置1～6
func (r *AlitripTicketRuleUploadAPIRequest) SetEnterVoucherType(_enterVoucherType int64) error {
	r._enterVoucherType = _enterVoucherType
	r.Set("enter_voucher_type", _enterVoucherType)
	return nil
}

// GetEnterVoucherType EnterVoucherType Getter
func (r AlitripTicketRuleUploadAPIRequest) GetEnterVoucherType() int64 {
	return r._enterVoucherType
}

// SetEnterVoucherValue is EnterVoucherValue Setter
// 其他入园凭证类型。当enter_voucher_type=6时，填写其他入园凭证类型。
func (r *AlitripTicketRuleUploadAPIRequest) SetEnterVoucherValue(_enterVoucherValue string) error {
	r._enterVoucherValue = _enterVoucherValue
	r.Set("enter_voucher_value", _enterVoucherValue)
	return nil
}

// GetEnterVoucherValue EnterVoucherValue Getter
func (r AlitripTicketRuleUploadAPIRequest) GetEnterVoucherValue() string {
	return r._enterVoucherValue
}

// SetTicketChangeAdderss is TicketChangeAdderss Setter
// 新发布规则时特殊选填。换票地址。当enter_type为2时，该字段必填。
func (r *AlitripTicketRuleUploadAPIRequest) SetTicketChangeAdderss(_ticketChangeAdderss string) error {
	r._ticketChangeAdderss = _ticketChangeAdderss
	r.Set("ticket_change_adderss", _ticketChangeAdderss)
	return nil
}

// GetTicketChangeAdderss TicketChangeAdderss Getter
func (r AlitripTicketRuleUploadAPIRequest) GetTicketChangeAdderss() string {
	return r._ticketChangeAdderss
}

// SetEnterAddress is EnterAddress Setter
// 新发布规则时必填。景区入园地址。游客在景区的详细入园地址，请仔细填写。
func (r *AlitripTicketRuleUploadAPIRequest) SetEnterAddress(_enterAddress string) error {
	r._enterAddress = _enterAddress
	r.Set("enter_address", _enterAddress)
	return nil
}

// GetEnterAddress EnterAddress Getter
func (r AlitripTicketRuleUploadAPIRequest) GetEnterAddress() string {
	return r._enterAddress
}

// SetFeeInclude is FeeInclude Setter
// 新发布规则时必填。门票费用包含，请详细说明该门票商品包含的费用信息。
func (r *AlitripTicketRuleUploadAPIRequest) SetFeeInclude(_feeInclude string) error {
	r._feeInclude = _feeInclude
	r.Set("fee_include", _feeInclude)
	return nil
}

// GetFeeInclude FeeInclude Getter
func (r AlitripTicketRuleUploadAPIRequest) GetFeeInclude() string {
	return r._feeInclude
}

// SetExtraDesc is ExtraDesc Setter
// 新发布规则时必填。门票商品一些游客须知的补充说明。1600字符限制
func (r *AlitripTicketRuleUploadAPIRequest) SetExtraDesc(_extraDesc string) error {
	r._extraDesc = _extraDesc
	r.Set("extra_desc", _extraDesc)
	return nil
}

// GetExtraDesc ExtraDesc Getter
func (r AlitripTicketRuleUploadAPIRequest) GetExtraDesc() string {
	return r._extraDesc
}
