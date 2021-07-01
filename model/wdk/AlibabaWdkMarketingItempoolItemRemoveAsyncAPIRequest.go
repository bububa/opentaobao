package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest
商品池删除商品 API请求
alibaba.wdk.marketing.itempool.item.remove.async

新模型下删除商品 */
type AlibabaWdkMarketingItempoolItemRemoveAsyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemPoolSku
	// 活动信息
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// New
