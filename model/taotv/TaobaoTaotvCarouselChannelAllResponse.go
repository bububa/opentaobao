package taotv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取所有频道列表 APIResponse
taobao.taotv.carousel.channel.all

获取所有频道列表，按照序号升序
*/
type TaobaoTaotvCarouselChannelAllAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"taotv_carousel_channel_all_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoTaotvCarouselChannelAllResult `json:"result,omitempty" xml:"