package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest
批量发布买赠商品 API请求
alibaba.wdk.marketing.buygift.item.add.async

批量发布买赠商品 */
type AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemBuyGiftSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// New
