package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest
特价批量移除商品 API请求
alibaba.wdk.marketing.discount.item.remove.async

特价批量移除商品 */
type AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemDiscountSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// New
