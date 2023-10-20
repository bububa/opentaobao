package taotv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotaotvcarouselchannelallAPIResponse 获取所有频道列表 API返回值
// taobao.taotv.carousel.channel.all
//
// 获取所有频道列表，按照序号升序
type TaobaotaotvcarouselchannelallAPIResponse struct {
	model.CommonResponse
	TaobaotaotvcarouselchannelallAPIResponseModel
}

// TaobaotaotvcarouselchannelallAPIResponseModel is 获取所有频道列表 成功返回结果
type TaobaotaotvcarouselchannelallAPIResponseModel struct {
	XMLName xml.Name `xml:"taotv_carousel_channel_all_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaotaotvcarouselchannelallResult `json:"result,omitempty" xml:"result,omitempty"`
}
