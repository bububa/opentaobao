package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
下线少儿推荐内容接口 API返回值 
yunos.tvpubadmin.content.child.recoitem.offline

下线少儿推荐内容接口
*/
type YunosTvpubadminContentChildRecoitemOfflineAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentChildRecoitemOfflineResponse
}

// 下线少儿推荐内容接口 成功返回结果
type YunosTvpubadminContentChildRecoitemOfflineResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_recoitem_offline_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 下线操作结果
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
}
