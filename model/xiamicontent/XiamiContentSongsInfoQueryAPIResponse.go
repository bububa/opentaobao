package xiamicontent

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentsongsinfoqueryAPIResponse 搜索歌曲列表 API返回值
// xiami.content.songs.info.query
//
// 多维度查询歌曲列表
type XiamicontentsongsinfoqueryAPIResponse struct {
	model.CommonResponse
	XiamicontentsongsinfoqueryAPIResponseModel
}

// XiamicontentsongsinfoqueryAPIResponseModel is 搜索歌曲列表 成功返回结果
type XiamicontentsongsinfoqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_songs_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的歌曲信息
	Songs *Page `json:"songs,omitempty" xml:"songs,omitempty"`
	// 系统自动生成
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
