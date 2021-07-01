package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSeriesEditAPIRequest
系列品变更-更新系列 API请求
alibaba.wdk.series.edit

系列品变更-更新系列 */
type AlibabaWdkSeriesEditAPIRequest struct {
	model.Params
	// 商品系列修改请求
	_series *SkuSeriesEditRequest
}

// New
