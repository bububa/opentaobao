package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSeriesSkuRemoveAPIRequest
系列品商品变更-移除商品 API请求
alibaba.wdk.series.sku.remove

系列品商品变更-移除商品 */
type AlibabaWdkSeriesSkuRemoveAPIRequest struct {
	model.Params
	// 系列品移除商品请求
	_seriesSkus *SeriesSkuRequest
}

// New
