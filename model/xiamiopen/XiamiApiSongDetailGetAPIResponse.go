package xiamiopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// XiamiapisongdetailgetAPIResponse 获取歌曲详情 API返回值
// xiami.api.song.detail.get
//
// 获取歌曲详情
type XiamiapisongdetailgetAPIResponse struct {
	model.CommonResponse
	XiamiapisongdetailgetAPIResponseModel
}

// XiamiapisongdetailgetAPIResponseModel is 获取歌曲详情 成功返回结果
type XiamiapisongdetailgetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_api_song_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 歌曲信息
	SongDetailList []SongDetailDo `json:"song_detail_list,omitempty" xml:"song_detail_list>song_detail_do,omitempty"`
}
