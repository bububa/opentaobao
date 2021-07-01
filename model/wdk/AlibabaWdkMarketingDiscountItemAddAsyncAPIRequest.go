package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest
特价批量新增商品 API请求
alibaba.wdk.marketing.discount.item.add.async

新分组模型下新增商品 */
type AlibabaWdkMarketingDiscountItemAddAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemDiscountSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// New
