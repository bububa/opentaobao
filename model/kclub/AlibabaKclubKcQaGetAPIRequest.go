package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaKclubKcQaGetAPIRequest
知识云-查询单个知识详情 API请求
alibaba.kclub.kc.qa.get

知识云-查询单个知识详情。通过租户id、问题id查询问题详情 */
type AlibabaKclubKcQaGetAPIRequest struct {
	model.Params
	// 问题id
	_questionId int64
	// 过滤条件
	_filter *KcQaFilter
	// 鉴权
	_auth *TenancyAuth
}

// NewAlibabaKclubKcQaGetRequest 初始化AlibabaKclubKcQaGetAPIRequest对象
func NewAlibabaKclubKcQaGetRequest() *AlibabaKclubKcQaGetAPIRequest {
	return &AlibabaKclubKcQaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaKclubKcQaGetAPIRequest) GetApiMethodName() string {
	return "alibaba.kclub.kc.qa.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaKclubKcQaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is QuestionId Setter
// 问题id
func (r *AlibabaKclubKcQaGetAPIRequest) SetQuestionId(_questionId int64) error {
	r._questionId = _questionId
	r.Set("question_id", _questionId)
	return nil
}

// Get QuestionId Getter
func (r AlibabaKclubKcQaGetAPIRequest) GetQuestionId() int64 {
	return r._questionId
}

// Set is Filter Setter
// 过滤条件
func (r *AlibabaKclubKcQaGetAPIRequest) SetFilter(_filter *KcQaFilter) error {
	r._filter = _filter
	r.Set("filter", _filter)
	return nil
}

// Get Filter Getter
func (r AlibabaKclubKcQaGetAPIRequest) GetFilter() *KcQaFilter {
	return r._filter
}

// Set is Auth Setter
// 鉴权
func (r *AlibabaKclubKcQaGetAPIRequest) SetAuth(_auth *TenancyAuth) error {
	r._auth = _auth
	r.Set("auth", _auth)
	return nil
}

// Get Auth Getter
func (r AlibabaKclubKcQaGetAPIRequest) GetAuth() *TenancyAuth {
	return r._auth
}
