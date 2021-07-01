package xiamiatrist

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* XiamiContentArtistInfoQueryAPIResponse
搜索艺人列表 API返回值
xiami.content.artist.info.query

根据查询条件，搜索相关艺人列表 */
type XiamiContentArtistInfoQueryAPIResponse struct {
	model.CommonResponse
	XiamiContentArtistInfoQueryAPIResponseModel
}

// XiamiContentArtistInfoQueryAPIResponseModel is 搜索艺人列表 成功返回结果
type XiamiContentArtistInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_artist_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的艺人分页信息
	ArtistPage *Page `json:"artist_page,omitempty" xml:"artist_page,omitempty"`
	// 执行结果
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
