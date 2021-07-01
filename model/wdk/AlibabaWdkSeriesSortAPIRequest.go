package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSeriesSortAPIRequest
系列品-商品排序 API请求
alibaba.wdk.series.sort

系列品商品变更-商品排序 */
type AlibabaWdkSeriesSortAPIRequest struct {
	model.Params
	// 自定义排序请求
	_sort *SeriesSortRequest
}

// New
