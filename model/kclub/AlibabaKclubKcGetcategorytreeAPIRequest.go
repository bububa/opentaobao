package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaKclubKcGetcategorytreeAPIRequest
知识云-查询租户下类目树 API请求
alibaba.kclub.kc.getcategorytree

知识云-查询租户下类目树。通过租户id、类型(外部类目、帮助中心类目、内部类目)。 */
type AlibabaKclubKcGetcategorytreeAPIRequest struct {
	model.Params
	// 租户id
	_tenantId int64
	// 鉴权参数
	_auth *TenancyAuth
}

// New
