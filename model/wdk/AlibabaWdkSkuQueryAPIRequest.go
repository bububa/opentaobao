package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuQueryAPIRequest
查询商品 API请求
alibaba.wdk.sku.query

查询商品 */
type AlibabaWdkSkuQueryAPIRequest struct {
	model.Params
	// 入参
	_param *SkuQueryDo
}

// New
