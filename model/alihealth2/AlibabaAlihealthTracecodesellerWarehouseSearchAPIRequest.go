package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest
查询仓库api API请求
alibaba.alihealth.tracecodeseller.warehouse.search

查询仓库api */
type AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest struct {
	model.Params
	// 身份认证
	_appkey string
	// 商家id
	_entInfoId int64
	// 第几页
	_page int64
	// 每页多少条
	_pageSize int64
}

// New
