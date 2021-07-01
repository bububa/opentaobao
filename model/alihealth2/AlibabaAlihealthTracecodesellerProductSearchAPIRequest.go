package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerProductSearchAPIRequest
查询商品api API请求
alibaba.alihealth.tracecodeseller.product.search

查询商品api */
type AlibabaAlihealthTracecodesellerProductSearchAPIRequest struct {
	model.Params
	// 身份认证
	_skeyCode string
	// 商家id
	_entInfoId int64
	// 页数
	_page int64
	// 每页条数
	_pageSize int64
}

// New
