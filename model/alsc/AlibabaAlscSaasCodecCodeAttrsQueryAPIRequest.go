package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscsaascodeccodeattrsqueryAPIRequest 码业务属性查询 API请求
// alibaba.alsc.saas.codec.code.attrs.query
//
// 码业务属性查询
type AlibabaalscsaascodeccodeattrsqueryAPIRequest struct {
	model.Params
	// 请求入参
	_queryCodeRequest *QueryCodeBizAttrRequest
}

// NewAlibabaalscsaascodeccodeattrsqueryRequest 初始化AlibabaalscsaascodeccodeattrsqueryAPIRequest对象
func NewAlibabaalscsaascodeccodeattrsqueryRequest() *AlibabaalscsaascodeccodeattrsqueryAPIRequest {
	return &AlibabaalscsaascodeccodeattrsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalscsaascodeccodeattrsqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.saas.codec.code.attrs.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalscsaascodeccodeattrsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalscsaascodeccodeattrsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryCodeRequest is QueryCodeRequest Setter
// 请求入参
func (r *AlibabaalscsaascodeccodeattrsqueryAPIRequest) SetQueryCodeRequest(_queryCodeRequest *QueryCodeBizAttrRequest) error {
	r._queryCodeRequest = _queryCodeRequest
	r.Set("query_code_request", _queryCodeRequest)
	return nil
}

// GetQueryCodeRequest QueryCodeRequest Getter
func (r AlibabaalscsaascodeccodeattrsqueryAPIRequest) GetQueryCodeRequest() *QueryCodeBizAttrRequest {
	return r._queryCodeRequest
}
