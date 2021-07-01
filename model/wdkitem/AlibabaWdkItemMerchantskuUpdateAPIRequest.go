package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemMerchantskuUpdateAPIRequest
商家商品修改 API请求
alibaba.wdk.item.merchantsku.update

商家商品修改 */
type AlibabaWdkItemMerchantskuUpdateAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 参数json
	_params string
	// 机构编码
	_orgCode string
}

// New
