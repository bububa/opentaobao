package xiamicontent

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentAlbumInfoGetAPIResponse 获取专辑信息 API返回值
// xiami.content.album.info.get
//
// 获取专辑信息
type XiamiContentAlbumInfoGetAPIResponse struct {
	model.CommonResponse
	XiamiContentAlbumInfoGetAPIResponseModel
}

// XiamiContentAlbumInfoGetAPIResponseModel is 获取专辑信息 成功返回结果
type XiamiContentAlbumInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_album_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 专辑信息
	AlbumList []AlbumDto `json:"album_list,omitempty" xml:"album_list>album_dto,omitempty"`
	// 返回结果
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
