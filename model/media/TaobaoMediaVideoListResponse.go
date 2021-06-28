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
    // Response *TaobaoMediaVideoListResponse `json:"media_video_list_response,omitempty"` 
    TaobaoMediaVideoListResponse
}

/* model for simplify = false
type TaobaoMediaVideoListResponse struct {

    // 返回结果
    
    Result  *struct {
        SearchResultDO  *SearchResultDO `json:"search_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoMediaVideoListResponse struct {

    // 返回结果
    Result   *SearchResultDO `json:"result,omitempty"`

}
