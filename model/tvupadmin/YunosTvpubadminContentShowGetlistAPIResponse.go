package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentShowGetlistAPIResponse
节目审核获取节目列表 API返回值
yunos.tvpubadmin.content.show.getlist

节目审核获取节目列表 */
type YunosTvpubadminContentShowGetlistAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentShowGetlistAPIResponseModel
}

// YunosTvpubadminContentShowGetlistAPIResponseModel is 节目审核获取节目列表 成功返回结果
type YunosTvpubadminContentShowGetlistAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_getlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}
