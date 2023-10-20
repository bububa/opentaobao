package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapurmediastatisticsAPIRequest 新媒体统计信息 API请求
// alibaba.pur.media.statistics
//
// 清博同步新媒体的统计信息给到采购平台
type AlibabapurmediastatisticsAPIRequest struct {
	model.Params
	// 新媒体统计对象
	_mediaStatisticsDTO []MediaStatisticsDto
}

// NewAlibabapurmediastatisticsRequest 初始化AlibabapurmediastatisticsAPIRequest对象
func NewAlibabapurmediastatisticsRequest() *AlibabapurmediastatisticsAPIRequest {
	return &AlibabapurmediastatisticsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapurmediastatisticsAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.media.statistics"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapurmediastatisticsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapurmediastatisticsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMediaStatisticsDTO is MediaStatisticsDTO Setter
// 新媒体统计对象
func (r *AlibabapurmediastatisticsAPIRequest) SetMediaStatisticsDTO(_mediaStatisticsDTO []MediaStatisticsDto) error {
	r._mediaStatisticsDTO = _mediaStatisticsDTO
	r.Set("media_statistics_d_t_o", _mediaStatisticsDTO)
	return nil
}

// GetMediaStatisticsDTO MediaStatisticsDTO Getter
func (r AlibabapurmediastatisticsAPIRequest) GetMediaStatisticsDTO() []MediaStatisticsDto {
	return r._mediaStatisticsDTO
}
