package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSeriesCreateAPIRequest
系列品变更-新增系列 API请求
alibaba.wdk.series.create

系列品变更-新增系列 */
type AlibabaWdkSeriesCreateAPIRequest struct {
	model.Params
	// 系列品创建系列请求
	_series *SkuSeriesCreateRequest
}

// New
