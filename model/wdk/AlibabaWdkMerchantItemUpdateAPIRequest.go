package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMerchantItemUpdateAPIRequest
修改商家商品 API请求
alibaba.wdk.merchant.item.update

修改商家商品 */
type AlibabaWdkMerchantItemUpdateAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 门店编码
	_merchantCode string
	// 修改字段的json
	_params string
}

// New
