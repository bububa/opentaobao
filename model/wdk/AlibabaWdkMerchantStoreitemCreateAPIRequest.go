package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMerchantStoreitemCreateAPIRequest
新建门店商品 API请求
alibaba.wdk.merchant.storeitem.create

新建门店商品 */
type AlibabaWdkMerchantStoreitemCreateAPIRequest struct {
	model.Params
	// 门店id
	_storeId string
	// 商品编码
	_skuCode string
	// 商家编码
	_merchantCode string
	// 新增门店商品参数，具体字段详见文档
	_params string
}

// New
