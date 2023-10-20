package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTableauditQuerylauncherAPIResponse 运营位管控-查询联盟一体机运营位元数据列表 API返回值
// yunos.tvpubadmin.content.tableaudit.querylauncher
//
// 运营位管控-查询联盟一体机运营位元数据列表
type YunosTvpubadminContentTableauditQuerylauncherAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentTableauditQuerylauncherAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTableauditQuerylauncherAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentTableauditQuerylauncherAPIResponseModel).Reset()
}

// YunosTvpubadminContentTableauditQuerylauncherAPIResponseModel is 运营位管控-查询联盟一体机运营位元数据列表 成功返回结果
type YunosTvpubadminContentTableauditQuerylauncherAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_tableaudit_querylauncher_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询桌面坑位内容列表（用于联盟、一体机）
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTableauditQuerylauncherAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentTableauditQuerylauncherAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentTableauditQuerylauncherAPIResponse)
	},
}

// GetYunosTvpubadminContentTableauditQuerylauncherAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentTableauditQuerylauncherAPIResponse
func GetYunosTvpubadminContentTableauditQuerylauncherAPIResponse() *YunosTvpubadminContentTableauditQuerylauncherAPIResponse {
	return poolYunosTvpubadminContentTableauditQuerylauncherAPIResponse.Get().(*YunosTvpubadminContentTableauditQuerylauncherAPIResponse)
}

// ReleaseYunosTvpubadminContentTableauditQuerylauncherAPIResponse 将 YunosTvpubadminContentTableauditQuerylauncherAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentTableauditQuerylauncherAPIResponse(v *YunosTvpubadminContentTableauditQuerylauncherAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentTableauditQuerylauncherAPIResponse.Put(v)
}
