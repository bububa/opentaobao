package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSeriesDefaultskuResetAPIRequest
系列品商品变更-重置默认商品 API请求
alibaba.wdk.series.defaultsku.reset

系列品商品变更-重置默认商品 */
type AlibabaWdkSeriesDefaultskuResetAPIRequest struct {
	model.Params
	// 系列品重置默认商品请求
	_seriesSku *SeriesSkuRequest
}

// New
