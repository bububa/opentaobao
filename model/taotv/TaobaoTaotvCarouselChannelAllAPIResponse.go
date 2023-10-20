package taotv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvCarouselChannelAllAPIResponse 获取所有频道列表 API返回值
// taobao.taotv.carousel.channel.all
//
// 获取所有频道列表，按照序号升序
type TaobaoTaotvCarouselChannelAllAPIResponse struct {
	model.CommonResponse
	TaobaoTaotvCarouselChannelAllAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTaotvCarouselChannelAllAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTaotvCarouselChannelAllAPIResponseModel).Reset()
}

// TaobaoTaotvCarouselChannelAllAPIResponseModel is 获取所有频道列表 成功返回结果
type TaobaoTaotvCarouselChannelAllAPIResponseModel struct {
	XMLName xml.Name `xml:"taotv_carousel_channel_all_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoTaotvCarouselChannelAllResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTaotvCarouselChannelAllAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTaotvCarouselChannelAllAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvCarouselChannelAllAPIResponse)
	},
}

// GetTaobaoTaotvCarouselChannelAllAPIResponse 从 sync.Pool 获取 TaobaoTaotvCarouselChannelAllAPIResponse
func GetTaobaoTaotvCarouselChannelAllAPIResponse() *TaobaoTaotvCarouselChannelAllAPIResponse {
	return poolTaobaoTaotvCarouselChannelAllAPIResponse.Get().(*TaobaoTaotvCarouselChannelAllAPIResponse)
}

// ReleaseTaobaoTaotvCarouselChannelAllAPIResponse 将 TaobaoTaotvCarouselChannelAllAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTaotvCarouselChannelAllAPIResponse(v *TaobaoTaotvCarouselChannelAllAPIResponse) {
	v.Reset()
	poolTaobaoTaotvCarouselChannelAllAPIResponse.Put(v)
}
