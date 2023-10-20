package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildNodeitemOfflineAPIResponse 少儿大厅类目内容下线接口 API返回值
// yunos.tvpubadmin.content.child.nodeitem.offline
//
// 少儿大厅类目内容下线接口
type YunosTvpubadminContentChildNodeitemOfflineAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentChildNodeitemOfflineAPIResponseModel
}

// YunosTvpubadminContentChildNodeitemOfflineAPIResponseModel is 少儿大厅类目内容下线接口 成功返回结果
type YunosTvpubadminContentChildNodeitemOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_nodeitem_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 执行结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
