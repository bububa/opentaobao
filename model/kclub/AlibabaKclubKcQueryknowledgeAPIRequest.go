package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaKclubKcQueryknowledgeAPIRequest
知识云-通用知识查询服务 API请求
alibaba.kclub.kc.queryknowledge

知识云-通用知识查询服务。通过租户id、类目id、知识类型、知识状态等条件查询类目。 */
type AlibabaKclubKcQueryknowledgeAPIRequest struct {
	model.Params
	// 查询条件
	_kcQaQuery *KcQaQuery
	// 鉴权
	_auth *TenancyAuth
}

// New
