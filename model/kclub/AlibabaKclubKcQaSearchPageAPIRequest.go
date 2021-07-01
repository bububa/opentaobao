package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaKclubKcQaSearchPageAPIRequest
知识云-知识检索(分页) API请求
alibaba.kclub.kc.qa.search.page

知识云-知识搜索服务 */
type AlibabaKclubKcQaSearchPageAPIRequest struct {
	model.Params
	// 查询参数
	_query *KcSearchQuestionQuery
	// 鉴权
	_auth *TenancyAuth
}

// New
