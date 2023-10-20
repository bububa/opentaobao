package tvupadmin

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *YunosTvpubadminContentChildNodeitemOfflineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentChildNodeitemOfflineAPIResponseModel).Reset()
}

// YunosTvpubadminContentChildNodeitemOfflineAPIResponseModel is 少儿大厅类目内容下线接口 成功返回结果
type YunosTvpubadminContentChildNodeitemOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_nodeitem_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 执行结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChildNodeitemOfflineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolYunosTvpubadminContentChildNodeitemOfflineAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentChildNodeitemOfflineAPIResponse)
	},
}

// GetYunosTvpubadminContentChildNodeitemOfflineAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentChildNodeitemOfflineAPIResponse
func GetYunosTvpubadminContentChildNodeitemOfflineAPIResponse() *YunosTvpubadminContentChildNodeitemOfflineAPIResponse {
	return poolYunosTvpubadminContentChildNodeitemOfflineAPIResponse.Get().(*YunosTvpubadminContentChildNodeitemOfflineAPIResponse)
}

// ReleaseYunosTvpubadminContentChildNodeitemOfflineAPIResponse 将 YunosTvpubadminContentChildNodeitemOfflineAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentChildNodeitemOfflineAPIResponse(v *YunosTvpubadminContentChildNodeitemOfflineAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentChildNodeitemOfflineAPIResponse.Put(v)
}
