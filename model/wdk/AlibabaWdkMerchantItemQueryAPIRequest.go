package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMerchantItemQueryAPIRequest
商家商品查询 API请求
alibaba.wdk.merchant.item.query

商家商品查询 */
type AlibabaWdkMerchantItemQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 商家编码
	_merchantCode string
}

// New
