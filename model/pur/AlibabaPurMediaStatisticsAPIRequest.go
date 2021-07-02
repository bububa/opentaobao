package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurMediaStatisticsAPIRequest 新媒体统计信息 API请求
// alibaba.pur.media.statistics
//
// 清博同步新媒体的统计信息给到采购平台
type AlibabaPurMediaStatisticsAPIRequest struct {
	model.Params
	// 新媒体统计对象
	_mediaStatisticsDTO []MediaStatisticsDto
}

// NewAlibabaPurMediaStatisticsRequest 初始化AlibabaPurMediaStatisticsAPIRequest对象
func NewAlibabaPurMediaStatisticsRequest() *AlibabaPurMediaStatisticsAPIRequest {
	return &AlibabaPurMediaStatisticsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurMediaStatisticsAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.media.statistics"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurMediaStatisticsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MediaStatisticsDTO Setter
// 新媒体统计对象
func (r *AlibabaPurMediaStatisticsAPIRequest) SetMediaStatisticsDTO(_mediaStatisticsDTO []MediaStatisticsDto) error {
	r._mediaStatisticsDTO = _mediaStatisticsDTO
	r.Set("media_statistics_d_t_o", _mediaStatisticsDTO)
	return nil
}

// Get MediaStatisticsDTO Getter
func (r AlibabaPurMediaStatisticsAPIRequest) GetMediaStatisticsDTO() []MediaStatisticsDto {
	return r._mediaStatisticsDTO
}
