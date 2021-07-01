package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPurMediaStatisticsAPIRequest
新媒体统计信息 API请求
alibaba.pur.media.statistics

清博同步新媒体的统计信息给到采购平台 */
type AlibabaPurMediaStatisticsAPIRequest struct {
	model.Params
	// 新媒体统计对象
	_mediaStatisticsDTO []MediaStatisticsDto
}

// New
