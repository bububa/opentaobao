package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolAdditemAPIRequest
增加商品池里面的商品 API请求
alibaba.wdk.marketing.itempool.additem

增加商品池里面的商品 */
type AlibabaWdkMarketingItempoolAdditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// New
