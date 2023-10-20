package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointRuleGetAPIRequest 查询积分规则 API请求
// alibaba.alsc.crm.point.rule.get
//
// 新增积分规则查询接口,传入includeLogicalDelete和maxUpdateTime时走同步下行逻辑不然则走普通积分规则查询接口
type AlibabaAlscCrmPointRuleGetAPIRequest struct {
	model.Params
	// 入参
	_paramQueryPointRuleOpenReq *QueryPointRuleOpenReq
}

// NewAlibabaAlscCrmPointRuleGetRequest 初始化AlibabaAlscCrmPointRuleGetAPIRequest对象
func NewAlibabaAlscCrmPointRuleGetRequest() *AlibabaAlscCrmPointRuleGetAPIRequest {
	return &AlibabaAlscCrmPointRuleGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmPointRuleGetAPIRequest) Reset() {
	r._paramQueryPointRuleOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointRuleGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.rule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointRuleGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmPointRuleGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryPointRuleOpenReq is ParamQueryPointRuleOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointRuleGetAPIRequest) SetParamQueryPointRuleOpenReq(_paramQueryPointRuleOpenReq *QueryPointRuleOpenReq) error {
	r._paramQueryPointRuleOpenReq = _paramQueryPointRuleOpenReq
	r.Set("param_query_point_rule_open_req", _paramQueryPointRuleOpenReq)
	return nil
}

// GetParamQueryPointRuleOpenReq ParamQueryPointRuleOpenReq Getter
func (r AlibabaAlscCrmPointRuleGetAPIRequest) GetParamQueryPointRuleOpenReq() *QueryPointRuleOpenReq {
	return r._paramQueryPointRuleOpenReq
}

var poolAlibabaAlscCrmPointRuleGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmPointRuleGetRequest()
	},
}

// GetAlibabaAlscCrmPointRuleGetRequest 从 sync.Pool 获取 AlibabaAlscCrmPointRuleGetAPIRequest
func GetAlibabaAlscCrmPointRuleGetAPIRequest() *AlibabaAlscCrmPointRuleGetAPIRequest {
	return poolAlibabaAlscCrmPointRuleGetAPIRequest.Get().(*AlibabaAlscCrmPointRuleGetAPIRequest)
}

// ReleaseAlibabaAlscCrmPointRuleGetAPIRequest 将 AlibabaAlscCrmPointRuleGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmPointRuleGetAPIRequest(v *AlibabaAlscCrmPointRuleGetAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmPointRuleGetAPIRequest.Put(v)
}
