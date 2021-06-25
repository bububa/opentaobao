package taotv

import (
    "github.com/bububa/opentaobao/model"
)

/* 
TV桌面为你推荐接口 APIResponse
youku.tv.desktop.toyou.recommend

提供为你推荐数据
*/
type YoukuTvDesktopToyouRecommendAPIResponse struct {
    model.CommonResponse
    Response *YoukuTvDesktopToyouRecommendResponse `json:"youku_tv_desktop_toyou_recommend_response,omitempty"`
}

type YoukuTvDesktopToyouRecommendResponse struct {

    // 响应的结果列表
    Results   []V5BaseItemRbo `json:"results,omitempty"`

}
