package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
运营位管控-查询联盟一体机运营位元数据列表 APIResponse
yunos.tvpubadmin.content.tableaudit.querylauncher

运营位管控-查询联盟一体机运营位元数据列表
*/
type YunosTvpubadminContentTableauditQuerylauncherAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentTableauditQuerylauncherResponse
}

type YunosTvpubadminContentTableauditQuerylauncherResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_tableaudit_querylauncher_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询桌面坑位内容列表（用于联盟、一体机）
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
