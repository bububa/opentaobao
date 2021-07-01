package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuStoreskuScrollQueryAPIRequest
门店商品批量查询接口 API请求
alibaba.wdk.sku.storesku.scroll.query

提供门店商品批量查询接口 */
type AlibabaWdkSkuStoreskuScrollQueryAPIRequest struct {
	model.Params
	// 经营的id
	_storeId string
	// 历游标，首次调用传递null，后续传递ScrollPageResult.getScrollId()
	_scrollId string
}

// New
