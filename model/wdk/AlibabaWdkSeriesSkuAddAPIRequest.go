package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSeriesSkuAddAPIRequest
系列品商品变更-添加商品 API请求
alibaba.wdk.series.sku.add

系列品商品变更-添加商品 */
type AlibabaWdkSeriesSkuAddAPIRequest struct {
	model.Params
	// 系列品添加商品请求
	_seriesSkus *SeriesSkuRequest
}

// New
