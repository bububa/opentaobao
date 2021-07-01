package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest
批量删除买赠商品 API请求
alibaba.wdk.marketing.buygift.item.remove.async

批量删除买赠商品 */
type AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemBuyGiftSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// New
