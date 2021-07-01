package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest
删除换购活动商品 API请求
alibaba.wdk.marketing.itempool.stair.removeitem

删除换购商品 */
type AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest struct {
	model.Params
	// 商品sku信息
	_param0 *ItemPoolSku
	// 活动信息
	_param1 *CommonActivityParam
}

// New
