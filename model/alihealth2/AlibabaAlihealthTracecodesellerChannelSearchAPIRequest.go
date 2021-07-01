package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerChannelSearchAPIRequest
查询渠道商api API请求
alibaba.alihealth.tracecodeseller.channel.search

查询渠道商api */
type AlibabaAlihealthTracecodesellerChannelSearchAPIRequest struct {
	model.Params
	// 身份认证
	_skeyCode string
	// 商家id
	_entInfoId int64
	// 第几页
	_page int64
	// 每页几条
	_pageSize int64
	// 0 出库 2 入库
	_outInType int64
}

// New
