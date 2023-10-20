package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentapponoffappbylicenseAPIResponse 应用上下架操作 API返回值
// yunos.tvpubadmin.content.app.onoffappbylicense
//
// 应用上下架操作
type YunostvpubadmincontentapponoffappbylicenseAPIResponse struct {
	model.CommonResponse
	YunostvpubadmincontentapponoffappbylicenseAPIResponseModel
}

// YunostvpubadmincontentapponoffappbylicenseAPIResponseModel is 应用上下架操作 成功返回结果
type YunostvpubadmincontentapponoffappbylicenseAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_app_onoffappbylicense_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true/false
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}
