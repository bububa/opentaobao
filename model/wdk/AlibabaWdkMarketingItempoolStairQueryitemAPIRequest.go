package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolStairQueryitemAPIRequest
换购商品查询 API请求
alibaba.wdk.marketing.itempool.stair.queryitem

换购商品查询 */
type AlibabaWdkMarketingItempoolStairQueryitemAPIRequest struct {
	model.Params
	// 换购商品查询参数
	_param0 *ActivitySkuQuery
}

// New
