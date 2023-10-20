package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTableauditQuerychilddesktopAPIResponse 迎客松查看小酷宝桌面坑位元数据列表 API返回值
// yunos.tvpubadmin.content.tableaudit.querychilddesktop
//
// 迎客松查看小酷宝桌面坑位元数据列表
type YunosTvpubadminContentTableauditQuerychilddesktopAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentTableauditQuerychilddesktopAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTableauditQuerychilddesktopAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentTableauditQuerychilddesktopAPIResponseModel).Reset()
}

// YunosTvpubadminContentTableauditQuerychilddesktopAPIResponseModel is 迎客松查看小酷宝桌面坑位元数据列表 成功返回结果
type YunosTvpubadminContentTableauditQuerychilddesktopAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_tableaudit_querychilddesktop_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTableauditQuerychilddesktopAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentTableauditQuerychilddesktopAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentTableauditQuerychilddesktopAPIResponse)
	},
}

// GetYunosTvpubadminContentTableauditQuerychilddesktopAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentTableauditQuerychilddesktopAPIResponse
func GetYunosTvpubadminContentTableauditQuerychilddesktopAPIResponse() *YunosTvpubadminContentTableauditQuerychilddesktopAPIResponse {
	return poolYunosTvpubadminContentTableauditQuerychilddesktopAPIResponse.Get().(*YunosTvpubadminContentTableauditQuerychilddesktopAPIResponse)
}

// ReleaseYunosTvpubadminContentTableauditQuerychilddesktopAPIResponse 将 YunosTvpubadminContentTableauditQuerychilddesktopAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentTableauditQuerychilddesktopAPIResponse(v *YunosTvpubadminContentTableauditQuerychilddesktopAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentTableauditQuerychilddesktopAPIResponse.Put(v)
}
