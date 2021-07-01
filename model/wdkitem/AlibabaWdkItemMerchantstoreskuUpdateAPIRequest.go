package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemMerchantstoreskuUpdateAPIRequest
门店商品信息修改 API请求
alibaba.wdk.item.merchantstoresku.update

门店商品信息修改 */
type AlibabaWdkItemMerchantstoreskuUpdateAPIRequest struct {
	model.Params
	// 门店ID
	_storeId string
	// 商品编码
	_skuCode string
	// 修改参数，是个json字符串
	_params string
}

// New
