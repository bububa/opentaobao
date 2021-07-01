package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemStoreskuQueryAPIRequest
门店商品信息查询 API请求
alibaba.wdk.item.storesku.query

门店商品信息查询 */
type AlibabaWdkItemStoreskuQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 门店ID
	_storeId string
}

// New
