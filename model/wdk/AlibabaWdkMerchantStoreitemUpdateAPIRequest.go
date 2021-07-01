package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMerchantStoreitemUpdateAPIRequest
修改门店商品 API请求
alibaba.wdk.merchant.storeitem.update

修改门店商品 */
type AlibabaWdkMerchantStoreitemUpdateAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 商家编码
	_merchantCode string
	// 门店编码
	_storeId string
	// 修改参数的json
	_params string
}

// New
