package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新媒体统计信息 APIRequest
alibaba.pur.media.statistics

清博同步新媒体的统计信息给到采购平台
*/
type AlibabaPurMediaStatisticsRequest struct {
    model.Params

    // 新媒体统计对象
    mediaStatisticsDTO   []MediaStatisticsDto 

}

func NewAlibabaPurMediaStatisticsRequest() *AlibabaPurMediaStatisticsRequest{
    return &AlibabaPurMediaStatisticsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPurMediaStatisticsRequest) GetApiMethodName() string {
    return "alibaba.pur.media.statistics"
}

func (r AlibabaPurMediaStatisticsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPurMediaStatisticsRequest) SetMediaStatisticsDTO(mediaStatisticsDTO []MediaStatisticsDto) error {
    r.mediaStatisticsDTO = mediaStatisticsDTO
    r.Set("media_statistics_d_t_o", mediaStatisticsDTO)
    return nil
}

func (r AlibabaPurMediaStatisticsRequest) GetMediaStatisticsDTO() []MediaStatisticsDto {
    return r.mediaStatisticsDTO
}

