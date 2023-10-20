package xiamiopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// XiamiapisonglistenfilegetAPIResponse 获取歌曲试听文件 API返回值
// xiami.api.song.listenfile.get
//
// 获取歌曲试听文件
type XiamiapisonglistenfilegetAPIResponse struct {
	model.CommonResponse
	XiamiapisonglistenfilegetAPIResponseModel
}

// XiamiapisonglistenfilegetAPIResponseModel is 获取歌曲试听文件 成功返回结果
type XiamiapisonglistenfilegetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_api_song_listenfile_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 歌曲试听文件列表
	SongPlayInfoList []SongPlayInfoDo `json:"song_play_info_list,omitempty" xml:"song_play_info_list>song_play_info_do,omitempty"`
}
