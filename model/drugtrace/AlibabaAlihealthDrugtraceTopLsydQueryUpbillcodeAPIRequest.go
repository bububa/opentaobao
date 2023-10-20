package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest 通过一个码，查询这个码对应的上游企业出库单的单据号 API请求
// alibaba.alihealth.drugtrace.top.lsyd.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
type AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业REF_ENT_ID （当前企业的唯一标识）
	_refEntId string
}

// NewAlibabaalihealthdrugtracetoplsydqueryupbillcodeRequest 初始化AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest对象
func NewAlibabaalihealthdrugtracetoplsydqueryupbillcodeRequest() *AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest {
	return &AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.query.upbillcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest) GetCode() string {
	return r._code
}

// SetRefEntId is RefEntId Setter
// 企业REF_ENT_ID （当前企业的唯一标识）
func (r *AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugtracetoplsydqueryupbillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
