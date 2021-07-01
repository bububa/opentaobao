package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentAppQueryappAPIResponse
查询应用信息 API返回值
yunos.tvpubadmin.content.app.queryapp

查询应用信息 */
type YunosTvpubadminContentAppQueryappAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentAppQueryappAPIResponseModel
}

// YunosTvpubadminContentAppQueryappAPIResponseModel is 查询应用信息 成功返回结果
type YunosTvpubadminContentAppQueryappAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_app_queryapp_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Result<AppInfo>
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}
