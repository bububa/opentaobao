package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMerchantStoreitemQueryAPIRequest
门店商品信心查询 API请求
alibaba.wdk.merchant.storeitem.query

门店商品信心查询 */
type AlibabaWdkMerchantStoreitemQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 商家编码
	_merchantCode string
	// 门店编码
	_storeId string
}

// New
