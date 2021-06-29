package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新媒体统计信息 API请求
alibaba.pur.media.statistics

清博同步新媒体的统计信息给到采购平台
*/
type AlibabaPurMediaStatisticsRequest struct {
    model.Params
    // 新媒体统计对象
    _mediaStatisticsDTO   []MediaStatisticsDTO
}

// 初始化AlibabaPurMediaStatisticsRequest对象
func NewAlibabaPurMediaStatisticsRequest() *AlibabaPurMediaStatisticsRequest{
    return &AlibabaPurMediaStatisticsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPurMediaStatisticsRequest) GetApiMethodName() string {
    return "alibaba.pur.media.statistics"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPurMediaStatisticsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MediaStatisticsDTO Setter
// 新媒体统计对象
func (r *AlibabaPurMediaStatisticsRequest) SetMediaStatisticsDTO(_mediaStatisticsDTO []MediaStatisticsDTO) error {
    r._mediaStatisticsDTO = _mediaStatisticsDTO
    r.Set("media_statistics_d_t_o", _mediaStatisticsDTO)
    return nil
}

// MediaStatisticsDTO Getter
func (r AlibabaPurMediaStatisticsRequest) GetMediaStatisticsDTO() []MediaStatisticsDTO {
    return r._mediaStatisticsDTO
}
