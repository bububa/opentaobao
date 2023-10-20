package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentSingleVideoGetlistAPIResponse 单视频审核获取视频列表 API返回值
// yunos.tvpubadmin.content.single.video.getlist
//
// 牌照方在审核单视频时获取单视频列表接口
type YunosTvpubadminContentSingleVideoGetlistAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentSingleVideoGetlistAPIResponseModel
}

// YunosTvpubadminContentSingleVideoGetlistAPIResponseModel is 单视频审核获取视频列表 成功返回结果
type YunosTvpubadminContentSingleVideoGetlistAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_single_video_getlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}
