package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolQueryitemsAPIRequest
查询商品池活动下的商品 API请求
alibaba.wdk.marketing.itempool.queryitems

查询商品池活动下面的商品 */
type AlibabaWdkMarketingItempoolQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// New
