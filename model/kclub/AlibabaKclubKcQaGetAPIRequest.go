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

// New
