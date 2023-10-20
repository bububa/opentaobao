package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleagreementpayqueryAPIRequest 代扣详情查询 API请求
// alibaba.idle.agreement.pay.query
//
// 查询代扣结果
type AlibabaidleagreementpayqueryAPIRequest struct {
	model.Params
	// 入参
	_param *AgreementPayBillQueryParam
}

// NewAlibabaidleagreementpayqueryRequest 初始化AlibabaidleagreementpayqueryAPIRequest对象
func NewAlibabaidleagreementpayqueryRequest() *AlibabaidleagreementpayqueryAPIRequest {
	return &AlibabaidleagreementpayqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleagreementpayqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.agreement.pay.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleagreementpayqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleagreementpayqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaidleagreementpayqueryAPIRequest) SetParam(_param *AgreementPayBillQueryParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaidleagreementpayqueryAPIRequest) GetParam() *AgreementPayBillQueryParam {
	return r._param
}
