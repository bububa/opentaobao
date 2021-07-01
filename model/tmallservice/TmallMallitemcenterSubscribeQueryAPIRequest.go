package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMallitemcenterSubscribeQueryAPIRequest
天猫服务订购信息查询接口 API请求
tmall.mallitemcenter.subscribe.query

查询商家服务订购信息 */
type TmallMallitemcenterSubscribeQueryAPIRequest struct {
	model.Params
	// 入参query
	_query *Spb2bOderQuery
}

// New
