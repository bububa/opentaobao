package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAppQuerybylicenceAPIResponse 按牌照查询应用 API返回值
// yunos.tvpubadmin.content.app.querybylicence
//
// 按牌照查询应用
type YunosTvpubadminContentAppQuerybylicenceAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentAppQuerybylicenceAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentAppQuerybylicenceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentAppQuerybylicenceAPIResponseModel).Reset()
}

// YunosTvpubadminContentAppQuerybylicenceAPIResponseModel is 按牌照查询应用 成功返回结果
type YunosTvpubadminContentAppQuerybylicenceAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_app_querybylicence_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Result&lt;AppInfo&gt;
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentAppQuerybylicenceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentAppQuerybylicenceAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentAppQuerybylicenceAPIResponse)
	},
}

// GetYunosTvpubadminContentAppQuerybylicenceAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentAppQuerybylicenceAPIResponse
func GetYunosTvpubadminContentAppQuerybylicenceAPIResponse() *YunosTvpubadminContentAppQuerybylicenceAPIResponse {
	return poolYunosTvpubadminContentAppQuerybylicenceAPIResponse.Get().(*YunosTvpubadminContentAppQuerybylicenceAPIResponse)
}

// ReleaseYunosTvpubadminContentAppQuerybylicenceAPIResponse 将 YunosTvpubadminContentAppQuerybylicenceAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentAppQuerybylicenceAPIResponse(v *YunosTvpubadminContentAppQuerybylicenceAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentAppQuerybylicenceAPIResponse.Put(v)
}
