package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuWarehouseskuQueryAPIRequest
仓商品查询接口(指定商品编码) API请求
alibaba.wdk.sku.warehousesku.query

提供指定仓商品编码查询 */
type AlibabaWdkSkuWarehouseskuQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCodeList []string
	// 仓编码
	_warehouseCode string
}

// New
