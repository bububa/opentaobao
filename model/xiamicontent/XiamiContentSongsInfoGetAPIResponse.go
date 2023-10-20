package xiamicontent

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentsongsinfogetAPIResponse 获取歌曲信息 API返回值
// xiami.content.songs.info.get
//
// (批量)获取歌曲信息
type XiamicontentsongsinfogetAPIResponse struct {
	model.CommonResponse
	XiamicontentsongsinfogetAPIResponseModel
}

// XiamicontentsongsinfogetAPIResponseModel is 获取歌曲信息 成功返回结果
type XiamicontentsongsinfogetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_songs_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 歌曲信息
	Songs []SongInfoDto `json:"songs,omitempty" xml:"songs>song_info_dto,omitempty"`
	// 系统自动生成
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
