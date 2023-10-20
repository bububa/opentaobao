package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAdvertGettypesAPIResponse 获取广告位类型 API返回值
// yunos.tvpubadmin.content.advert.gettypes
//
// 获取广告位类型
type YunosTvpubadminContentAdvertGettypesAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentAdvertGettypesAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentAdvertGettypesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentAdvertGettypesAPIResponseModel).Reset()
}

// YunosTvpubadminContentAdvertGettypesAPIResponseModel is 获取广告位类型 成功返回结果
type YunosTvpubadminContentAdvertGettypesAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_advert_gettypes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Map&lt;Integer, String&gt; json格式
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentAdvertGettypesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentAdvertGettypesAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentAdvertGettypesAPIResponse)
	},
}

// GetYunosTvpubadminContentAdvertGettypesAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentAdvertGettypesAPIResponse
func GetYunosTvpubadminContentAdvertGettypesAPIResponse() *YunosTvpubadminContentAdvertGettypesAPIResponse {
	return poolYunosTvpubadminContentAdvertGettypesAPIResponse.Get().(*YunosTvpubadminContentAdvertGettypesAPIResponse)
}

// ReleaseYunosTvpubadminContentAdvertGettypesAPIResponse 将 YunosTvpubadminContentAdvertGettypesAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentAdvertGettypesAPIResponse(v *YunosTvpubadminContentAdvertGettypesAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentAdvertGettypesAPIResponse.Put(v)
}
