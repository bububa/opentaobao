package pur

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPurMediaStatisticsAPIRequest) Reset() {
	r._mediaStatisticsDTO = r._mediaStatisticsDTO[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurMediaStatisticsAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.media.statistics"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurMediaStatisticsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPurMediaStatisticsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMediaStatisticsDTO is MediaStatisticsDTO Setter
// 新媒体统计对象
func (r *AlibabaPurMediaStatisticsAPIRequest) SetMediaStatisticsDTO(_mediaStatisticsDTO []MediaStatisticsDto) error {
	r._mediaStatisticsDTO = _mediaStatisticsDTO
	r.Set("media_statistics_d_t_o", _mediaStatisticsDTO)
	return nil
}

// GetMediaStatisticsDTO MediaStatisticsDTO Getter
func (r AlibabaPurMediaStatisticsAPIRequest) GetMediaStatisticsDTO() []MediaStatisticsDto {
	return r._mediaStatisticsDTO
}

var poolAlibabaPurMediaStatisticsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPurMediaStatisticsRequest()
	},
}

// GetAlibabaPurMediaStatisticsRequest 从 sync.Pool 获取 AlibabaPurMediaStatisticsAPIRequest
func GetAlibabaPurMediaStatisticsAPIRequest() *AlibabaPurMediaStatisticsAPIRequest {
	return poolAlibabaPurMediaStatisticsAPIRequest.Get().(*AlibabaPurMediaStatisticsAPIRequest)
}

// ReleaseAlibabaPurMediaStatisticsAPIRequest 将 AlibabaPurMediaStatisticsAPIRequest 放入 sync.Pool
func ReleaseAlibabaPurMediaStatisticsAPIRequest(v *AlibabaPurMediaStatisticsAPIRequest) {
	v.Reset()
	poolAlibabaPurMediaStatisticsAPIRequest.Put(v)
}
