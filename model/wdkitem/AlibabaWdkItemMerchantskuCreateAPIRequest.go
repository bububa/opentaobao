package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemMerchantskuCreateAPIRequest
商家商品信息新建 API请求
alibaba.wdk.item.merchantsku.create

商家商品信息新建 */
type AlibabaWdkItemMerchantskuCreateAPIRequest struct {
	model.Params
	// 新建商品参数，是个json字符串
	_params string
}

// New
