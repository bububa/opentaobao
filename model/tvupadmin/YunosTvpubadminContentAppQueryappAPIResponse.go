package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAppQueryappAPIResponse 查询应用信息 API返回值
// yunos.tvpubadmin.content.app.queryapp
//
// 查询应用信息
type YunosTvpubadminContentAppQueryappAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentAppQueryappAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentAppQueryappAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentAppQueryappAPIResponseModel).Reset()
}

// YunosTvpubadminContentAppQueryappAPIResponseModel is 查询应用信息 成功返回结果
type YunosTvpubadminContentAppQueryappAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_app_queryapp_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Result&lt;AppInfo&gt;
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentAppQueryappAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentAppQueryappAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentAppQueryappAPIResponse)
	},
}

// GetYunosTvpubadminContentAppQueryappAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentAppQueryappAPIResponse
func GetYunosTvpubadminContentAppQueryappAPIResponse() *YunosTvpubadminContentAppQueryappAPIResponse {
	return poolYunosTvpubadminContentAppQueryappAPIResponse.Get().(*YunosTvpubadminContentAppQueryappAPIResponse)
}

// ReleaseYunosTvpubadminContentAppQueryappAPIResponse 将 YunosTvpubadminContentAppQueryappAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentAppQueryappAPIResponse(v *YunosTvpubadminContentAppQueryappAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentAppQueryappAPIResponse.Put(v)
}
