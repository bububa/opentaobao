package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按牌照查询应用 API返回值 
yunos.tvpubadmin.content.app.querybylicence

按牌照查询应用
*/
type YunosTvpubadminContentAppQuerybylicenceAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentAppQuerybylicenceAPIResponseModel
}

// 按牌照查询应用 成功返回结果
type YunosTvpubadminContentAppQuerybylicenceAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_app_querybylicence_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // Result<AppInfo>
    Object   string `json:"object,omitempty" xml:"object,omitempty"`
}
