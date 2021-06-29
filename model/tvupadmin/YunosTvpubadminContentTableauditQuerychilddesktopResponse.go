package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松查看小酷宝桌面坑位元数据列表 APIResponse
yunos.tvpubadmin.content.tableaudit.querychilddesktop

迎客松查看小酷宝桌面坑位元数据列表
*/
type YunosTvpubadminContentTableauditQuerychilddesktopAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentTableauditQuerychilddesktopResponse
}

type YunosTvpubadminContentTableauditQuerychilddesktopResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_tableaudit_querychilddesktop_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
