package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest
仓商品遍历接口(游标) API请求
alibaba.wdk.sku.warehousesku.scroll.query

提供仓商品数据接口查询 */
type AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest struct {
	model.Params
	// 仓库编码
	_warehouseCode string
	// 游标
	_scrollId string
}

// New
