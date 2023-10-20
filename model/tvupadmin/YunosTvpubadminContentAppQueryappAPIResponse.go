package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentappqueryappAPIResponse 查询应用信息 API返回值
// yunos.tvpubadmin.content.app.queryapp
//
// 查询应用信息
type YunostvpubadmincontentappqueryappAPIResponse struct {
	model.CommonResponse
	YunostvpubadmincontentappqueryappAPIResponseModel
}

// YunostvpubadmincontentappqueryappAPIResponseModel is 查询应用信息 成功返回结果
type YunostvpubadmincontentappqueryappAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_app_queryapp_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Result&lt;AppInfo&gt;
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}
