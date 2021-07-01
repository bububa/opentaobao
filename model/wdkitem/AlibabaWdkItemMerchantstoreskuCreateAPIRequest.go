package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemMerchantstoreskuCreateAPIRequest
门店商品信息新建 API请求
alibaba.wdk.item.merchantstoresku.create

门店商品信息新建 */
type AlibabaWdkItemMerchantstoreskuCreateAPIRequest struct {
	model.Params
	// 门店编码
	_storeId string
	// 商品编码
	_skuCode string
	// 新建参数，是个json字符串
	_params string
	// 机构
	_orgCode string
}

// New
