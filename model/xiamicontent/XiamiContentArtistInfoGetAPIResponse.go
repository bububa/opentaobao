package xiamicontent

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentArtistInfoGetAPIResponse 获取艺人信息 API返回值
// xiami.content.artist.info.get
//
// (批量)获取艺人信息
type XiamiContentArtistInfoGetAPIResponse struct {
	model.CommonResponse
	XiamiContentArtistInfoGetAPIResponseModel
}

// XiamiContentArtistInfoGetAPIResponseModel is 获取艺人信息 成功返回结果
type XiamiContentArtistInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_artist_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 艺人列表
	Artists []ArtistDto `json:"artists,omitempty" xml:"artists>artist_dto,omitempty"`
	// 结果
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
