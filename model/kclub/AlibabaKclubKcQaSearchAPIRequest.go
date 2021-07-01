package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaKclubKcQaSearchAPIRequest
知识云-知识检索 API请求
alibaba.kclub.kc.qa.search

知识云-知识搜索服务 */
type AlibabaKclubKcQaSearchAPIRequest struct {
	model.Params
	// 查询参数
	_query *KcSearchQuestionQuery
	// 鉴权
	_auth *TenancyAuth
}

// New
