package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
少儿大厅类目内容下线接口 APIResponse
yunos.tvpubadmin.content.child.nodeitem.offline

少儿大厅类目内容下线接口
*/
type YunosTvpubadminContentChildNodeitemOfflineAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentChildNodeitemOfflineResponse
}

type YunosTvpubadminContentChildNodeitemOfflineResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_nodeitem_offline_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 执行结果
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
