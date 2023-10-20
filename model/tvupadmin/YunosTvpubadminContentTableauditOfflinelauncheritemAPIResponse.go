package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse 运营位管理-联盟一体机下线运营位内容 API返回值
// yunos.tvpubadmin.content.tableaudit.offlinelauncheritem
//
// 运营位管理-联盟一体机下线运营位内容
type YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponseModel).Reset()
}

// YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponseModel is 运营位管理-联盟一体机下线运营位内容 成功返回结果
type YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_tableaudit_offlinelauncheritem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下线桌面坑位内容（用于联盟、一体机）
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse)
	},
}

// GetYunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse
func GetYunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse() *YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse {
	return poolYunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse.Get().(*YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse)
}

// ReleaseYunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse 将 YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse(v *YunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentTableauditOfflinelauncheritemAPIResponse.Put(v)
}
