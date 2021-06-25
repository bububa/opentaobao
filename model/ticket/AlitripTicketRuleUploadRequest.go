package ticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】景点门票规则维护接口 APIRequest
alitrip.ticket.rule.upload

景点门票规则维护接口。该接口同时支持新发规则和编辑现有规则，如果out_rule_id下没有发布过规则，则系统将判断为新发一个规则，否则认为是编辑现有规则。
对于新发布规则的情况，有些参数是必填的，请仔细查看各字段说明。对于编辑的情况，除out_rule_id外都是可选，编辑情况支持增量更新（某个参数不传则使用该规则上原有值）
*/
type AlitripTicketRuleUploadRequest struct {
    model.Params

    // 商户票种规则id
    outRuleId   string 

    // 新发布规则时必填。商户票种规则名称
    outRuleName   string 

    // 商户景点编码。新发布规则时必填，out_scenic_id与ali_scenic_id二选一，至少填写其中一个
    outScenicId   string 

    // 阿里旅行景点编码。新发布规则时必填，out_scenic_id与ali_scenic_id二选一，至少填写其中一个
    aliScenicId   int64 

    // 可选，票种规则类型：0-实体票规则，1-电子票（包含手工票和直连票）规则。新发规则时不传 默认为1
    ruleType   int64 

    // 可选，规则状态。0-生效，-1-失效。新发规则时不传 默认生效状态。
    ruleStatus   int64 

    // 新发布规则时必填。退票类型。1-无条件退改， 2-有条件退改， 3-不可退改。
    refundType   int64 

    // （该字段已废弃，请使用结构化退改规则字段替代：refund_custom_rules）新发布规则时特殊选填。退票描述。当refund_type为2时，该字段必填
    refundDesc   string 

    // 可选，是否支持自动退款，0或为空时表示不支持。0-不支持，1-只支持"售中自动退款"，2-只支持"过期未使用自动退款"，3-同时支持"售中自动退款"和"过期未使用自动退款"
    autoRefundSupport   int64 

    // 结构化自定义退款规则（json数组格式），当refund_type为2时，该字段必填。【type字段说明（特别注意：2和3不能同时存在）：1（游玩日期前退改规则），2（游玩日期当日退改规则），3（区间票 游玩日期有效期内退改规则），4（游玩日期后退改规则），5（其他情况退改规则）。amount字段说明：收取的手续费。unit字段说明：手续费单位，1（票价百分比），2（固定金额，单位分）。】  【示例含义说明：1、游玩日期前2天16点30分前申请退款，每张票收取票价20%的手续费；2、(单日票）游玩日当天12点00分前申请退款，每张票收取票价50%手续费；3、（区间票）游玩日期有效期，最后一天11点00分前申请退款，每张票收取票价60%手续费；4、游玩日期后7天23点59分前申请退款，每张票收取票价80%手续费；5、其他情况，每张票收取固定100元手续费。】
    refundCustomRules   string 

    // 新发布规则时必填。出游人信息设置。1、不需要，2、仅需一位游客信息。3、需要所有游客信息。不填默认为1（不需要）。注：实体票（rule_type=2）不需要出游人信息，电子票（rule_type=1）一般是需要出游人信息
    visitorRequire   int64 

    // 新发布规则时特殊选填。需要的出游人信息，需要出游人信息时必填。 2:手机号,3:身份证,4:姓名,17：港澳居民居住证 18：台湾居民居住证 19：护照 20：港澳台居民往返内地通行证（回乡证）。示例格式：2,3,4
    visitorInfos   string 

    // 新发布规则时必填。出游人 是否限购。1:限购,2:不限购
    visitorLimitAble   int64 

    // 新发布规则时特殊选填。限购模式：mode为1按天, mode为2按周, mode为3按月
    visitorLimitMode   int64 

    // 新发布规则时特殊选填。限购类型。0-身份证限购， 1-手机号限购
    visitorLimitType   int64 

    // 新发布规则时特殊选填。限购数量
    visitorLimitNum   int64 

    // 新发布规则时必填。入园类型。1-用兑换凭证直接入园，2-用兑换凭证换票入园
    enterType   int64 

    // 新发布规则时必填。入园使用的凭证类型。1、二维码，2、身份证，3、二维码或身份证，4:数字码，5、手机号，6、其它。当enter_type为1时，该字段只能选择1～3 。而当enterType为2时，该字段可设置1～6
    enterVoucherType   int64 

    // 其他入园凭证类型。当enter_voucher_type=6时，填写其他入园凭证类型。
    enterVoucherValue   string 

    // 新发布规则时特殊选填。换票地址。当enter_type为2时，该字段必填。
    ticketChangeAdderss   string 

    // 新发布规则时必填。景区入园地址。游客在景区的详细入园地址，请仔细填写。
    enterAddress   string 

    // 新发布规则时必填。门票费用包含，请详细说明该门票商品包含的费用信息。
    feeInclude   string 

    // 新发布规则时必填。门票商品一些游客须知的补充说明。1600字符限制
    extraDesc   string 

}

func NewAlitripTicketRuleUploadRequest() *AlitripTicketRuleUploadRequest{
    return &AlitripTicketRuleUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTicketRuleUploadRequest) GetApiMethodName() string {
    return "alitrip.ticket.rule.upload"
}

func (r AlitripTicketRuleUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTicketRuleUploadRequest) SetOutRuleId(outRuleId string) error {
    r.outRuleId = outRuleId
    r.Set("out_rule_id", outRuleId)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetOutRuleId() string {
    return r.outRuleId
}

func (r *AlitripTicketRuleUploadRequest) SetOutRuleName(outRuleName string) error {
    r.outRuleName = outRuleName
    r.Set("out_rule_name", outRuleName)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetOutRuleName() string {
    return r.outRuleName
}

func (r *AlitripTicketRuleUploadRequest) SetOutScenicId(outScenicId string) error {
    r.outScenicId = outScenicId
    r.Set("out_scenic_id", outScenicId)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetOutScenicId() string {
    return r.outScenicId
}

func (r *AlitripTicketRuleUploadRequest) SetAliScenicId(aliScenicId int64) error {
    r.aliScenicId = aliScenicId
    r.Set("ali_scenic_id", aliScenicId)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetAliScenicId() int64 {
    return r.aliScenicId
}

func (r *AlitripTicketRuleUploadRequest) SetRuleType(ruleType int64) error {
    r.ruleType = ruleType
    r.Set("rule_type", ruleType)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetRuleType() int64 {
    return r.ruleType
}

func (r *AlitripTicketRuleUploadRequest) SetRuleStatus(ruleStatus int64) error {
    r.ruleStatus = ruleStatus
    r.Set("rule_status", ruleStatus)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetRuleStatus() int64 {
    return r.ruleStatus
}

func (r *AlitripTicketRuleUploadRequest) SetRefundType(refundType int64) error {
    r.refundType = refundType
    r.Set("refund_type", refundType)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetRefundType() int64 {
    return r.refundType
}

func (r *AlitripTicketRuleUploadRequest) SetRefundDesc(refundDesc string) error {
    r.refundDesc = refundDesc
    r.Set("refund_desc", refundDesc)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetRefundDesc() string {
    return r.refundDesc
}

func (r *AlitripTicketRuleUploadRequest) SetAutoRefundSupport(autoRefundSupport int64) error {
    r.autoRefundSupport = autoRefundSupport
    r.Set("auto_refund_support", autoRefundSupport)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetAutoRefundSupport() int64 {
    return r.autoRefundSupport
}

func (r *AlitripTicketRuleUploadRequest) SetRefundCustomRules(refundCustomRules string) error {
    r.refundCustomRules = refundCustomRules
    r.Set("refund_custom_rules", refundCustomRules)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetRefundCustomRules() string {
    return r.refundCustomRules
}

func (r *AlitripTicketRuleUploadRequest) SetVisitorRequire(visitorRequire int64) error {
    r.visitorRequire = visitorRequire
    r.Set("visitor_require", visitorRequire)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetVisitorRequire() int64 {
    return r.visitorRequire
}

func (r *AlitripTicketRuleUploadRequest) SetVisitorInfos(visitorInfos string) error {
    r.visitorInfos = visitorInfos
    r.Set("visitor_infos", visitorInfos)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetVisitorInfos() string {
    return r.visitorInfos
}

func (r *AlitripTicketRuleUploadRequest) SetVisitorLimitAble(visitorLimitAble int64) error {
    r.visitorLimitAble = visitorLimitAble
    r.Set("visitor_limit_able", visitorLimitAble)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetVisitorLimitAble() int64 {
    return r.visitorLimitAble
}

func (r *AlitripTicketRuleUploadRequest) SetVisitorLimitMode(visitorLimitMode int64) error {
    r.visitorLimitMode = visitorLimitMode
    r.Set("visitor_limit_mode", visitorLimitMode)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetVisitorLimitMode() int64 {
    return r.visitorLimitMode
}

func (r *AlitripTicketRuleUploadRequest) SetVisitorLimitType(visitorLimitType int64) error {
    r.visitorLimitType = visitorLimitType
    r.Set("visitor_limit_type", visitorLimitType)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetVisitorLimitType() int64 {
    return r.visitorLimitType
}

func (r *AlitripTicketRuleUploadRequest) SetVisitorLimitNum(visitorLimitNum int64) error {
    r.visitorLimitNum = visitorLimitNum
    r.Set("visitor_limit_num", visitorLimitNum)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetVisitorLimitNum() int64 {
    return r.visitorLimitNum
}

func (r *AlitripTicketRuleUploadRequest) SetEnterType(enterType int64) error {
    r.enterType = enterType
    r.Set("enter_type", enterType)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetEnterType() int64 {
    return r.enterType
}

func (r *AlitripTicketRuleUploadRequest) SetEnterVoucherType(enterVoucherType int64) error {
    r.enterVoucherType = enterVoucherType
    r.Set("enter_voucher_type", enterVoucherType)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetEnterVoucherType() int64 {
    return r.enterVoucherType
}

func (r *AlitripTicketRuleUploadRequest) SetEnterVoucherValue(enterVoucherValue string) error {
    r.enterVoucherValue = enterVoucherValue
    r.Set("enter_voucher_value", enterVoucherValue)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetEnterVoucherValue() string {
    return r.enterVoucherValue
}

func (r *AlitripTicketRuleUploadRequest) SetTicketChangeAdderss(ticketChangeAdderss string) error {
    r.ticketChangeAdderss = ticketChangeAdderss
    r.Set("ticket_change_adderss", ticketChangeAdderss)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetTicketChangeAdderss() string {
    return r.ticketChangeAdderss
}

func (r *AlitripTicketRuleUploadRequest) SetEnterAddress(enterAddress string) error {
    r.enterAddress = enterAddress
    r.Set("enter_address", enterAddress)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetEnterAddress() string {
    return r.enterAddress
}

func (r *AlitripTicketRuleUploadRequest) SetFeeInclude(feeInclude string) error {
    r.feeInclude = feeInclude
    r.Set("fee_include", feeInclude)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetFeeInclude() string {
    return r.feeInclude
}

func (r *AlitripTicketRuleUploadRequest) SetExtraDesc(extraDesc string) error {
    r.extraDesc = extraDesc
    r.Set("extra_desc", extraDesc)
    return nil
}

func (r AlitripTicketRuleUploadRequest) GetExtraDesc() string {
    return r.extraDesc
}

