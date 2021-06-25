package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取商家视频列表 APIResponse
taobao.media.video.list

用于获取授权商家的视频列表
*/
type TaobaoMediaVideoListAPIResponse struct {
    model.CommonResponse
    Response *TaobaoMediaVideoListResponse `json:"taobao_media_video_list_response,omitempty"`
}

type TaobaoMediaVideoListResponse struct {

    // 返回结果
    Result   *SearchResultDO `json:"result,omitempty"`

}
