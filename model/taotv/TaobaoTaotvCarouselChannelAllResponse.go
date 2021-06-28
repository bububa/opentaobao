package taotv

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取所有频道列表 APIResponse
taobao.taotv.carousel.channel.all

获取所有频道列表，按照序号升序
*/
type TaobaoTaotvCarouselChannelAllAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTaotvCarouselChannelAllResponse `json:"taotv_carousel_channel_all_response,omitempty"` 
    TaobaoTaotvCarouselChannelAllResponse
}

/* model for simplify = false
type TaobaoTaotvCarouselChannelAllResponse struct {

    // result
    
    Result  *struct {
        TaobaoTaotvCarouselChannelAllResult  *TaobaoTaotvCarouselChannelAllResult `json:"taobao_taotv_carousel_channel_all_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTaotvCarouselChannelAllResponse struct {

    // result
    Result   *TaobaoTaotvCarouselChannelAllResult `json:"result,omitempty"`

}
