package taotv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotaotvvideoplaylistgetAPIResponse 根据频道ID获取频道下节目单以及当前播放 API返回值
// taobao.taotv.video.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放
type TaobaotaotvvideoplaylistgetAPIResponse struct {
	model.CommonResponse
	TaobaotaotvvideoplaylistgetAPIResponseModel
}

// TaobaotaotvvideoplaylistgetAPIResponseModel is 根据频道ID获取频道下节目单以及当前播放 成功返回结果
type TaobaotaotvvideoplaylistgetAPIResponseModel struct {
	XMLName xml.Name `xml:"taotv_video_playlist_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaotaotvvideoplaylistgetResult `json:"result,omitempty" xml:"result,omitempty"`
}
