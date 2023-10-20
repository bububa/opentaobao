package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildRecoitemOfflineAPIResponse 下线少儿推荐内容接口 API返回值
// yunos.tvpubadmin.content.child.recoitem.offline
//
// 下线少儿推荐内容接口
type YunosTvpubadminContentChildRecoitemOfflineAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentChildRecoitemOfflineAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChildRecoitemOfflineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentChildRecoitemOfflineAPIResponseModel).Reset()
}

// YunosTvpubadminContentChildRecoitemOfflineAPIResponseModel is 下线少儿推荐内容接口 成功返回结果
type YunosTvpubadminContentChildRecoitemOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_recoitem_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下线操作结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChildRecoitemOfflineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolYunosTvpubadminContentChildRecoitemOfflineAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentChildRecoitemOfflineAPIResponse)
	},
}

// GetYunosTvpubadminContentChildRecoitemOfflineAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentChildRecoitemOfflineAPIResponse
func GetYunosTvpubadminContentChildRecoitemOfflineAPIResponse() *YunosTvpubadminContentChildRecoitemOfflineAPIResponse {
	return poolYunosTvpubadminContentChildRecoitemOfflineAPIResponse.Get().(*YunosTvpubadminContentChildRecoitemOfflineAPIResponse)
}

// ReleaseYunosTvpubadminContentChildRecoitemOfflineAPIResponse 将 YunosTvpubadminContentChildRecoitemOfflineAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentChildRecoitemOfflineAPIResponse(v *YunosTvpubadminContentChildRecoitemOfflineAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentChildRecoitemOfflineAPIResponse.Put(v)
}
