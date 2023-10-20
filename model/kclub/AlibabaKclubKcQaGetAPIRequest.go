package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabakclubkcqagetAPIRequest 知识云-查询单个知识详情 API请求
// alibaba.kclub.kc.qa.get
//
// 知识云-查询单个知识详情。通过租户id、问题id查询问题详情
type AlibabakclubkcqagetAPIRequest struct {
	model.Params
	// 问题id
	_questionId int64
	// 过滤条件
	_filter *KcQaFilter
	// 鉴权
	_auth *TenancyAuth
}

// NewAlibabakclubkcqagetRequest 初始化AlibabakclubkcqagetAPIRequest对象
func NewAlibabakclubkcqagetRequest() *AlibabakclubkcqagetAPIRequest {
	return &AlibabakclubkcqagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabakclubkcqagetAPIRequest) GetApiMethodName() string {
	return "alibaba.kclub.kc.qa.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabakclubkcqagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabakclubkcqagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuestionId is QuestionId Setter
// 问题id
func (r *AlibabakclubkcqagetAPIRequest) SetQuestionId(_questionId int64) error {
	r._questionId = _questionId
	r.Set("question_id", _questionId)
	return nil
}

// GetQuestionId QuestionId Getter
func (r AlibabakclubkcqagetAPIRequest) GetQuestionId() int64 {
	return r._questionId
}

// SetFilter is Filter Setter
// 过滤条件
func (r *AlibabakclubkcqagetAPIRequest) SetFilter(_filter *KcQaFilter) error {
	r._filter = _filter
	r.Set("filter", _filter)
	return nil
}

// GetFilter Filter Getter
func (r AlibabakclubkcqagetAPIRequest) GetFilter() *KcQaFilter {
	return r._filter
}

// SetAuth is Auth Setter
// 鉴权
func (r *AlibabakclubkcqagetAPIRequest) SetAuth(_auth *TenancyAuth) error {
	r._auth = _auth
	r.Set("auth", _auth)
	return nil
}

// GetAuth Auth Getter
func (r AlibabakclubkcqagetAPIRequest) GetAuth() *TenancyAuth {
	return r._auth
}
