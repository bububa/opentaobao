package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAppOnoffappbylicenseAPIResponse 应用上下架操作 API返回值
// yunos.tvpubadmin.content.app.onoffappbylicense
//
// 应用上下架操作
type YunosTvpubadminContentAppOnoffappbylicenseAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentAppOnoffappbylicenseAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentAppOnoffappbylicenseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentAppOnoffappbylicenseAPIResponseModel).Reset()
}

// YunosTvpubadminContentAppOnoffappbylicenseAPIResponseModel is 应用上下架操作 成功返回结果
type YunosTvpubadminContentAppOnoffappbylicenseAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_app_onoffappbylicense_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true/false
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentAppOnoffappbylicenseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminContentAppOnoffappbylicenseAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentAppOnoffappbylicenseAPIResponse)
	},
}

// GetYunosTvpubadminContentAppOnoffappbylicenseAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentAppOnoffappbylicenseAPIResponse
func GetYunosTvpubadminContentAppOnoffappbylicenseAPIResponse() *YunosTvpubadminContentAppOnoffappbylicenseAPIResponse {
	return poolYunosTvpubadminContentAppOnoffappbylicenseAPIResponse.Get().(*YunosTvpubadminContentAppOnoffappbylicenseAPIResponse)
}

// ReleaseYunosTvpubadminContentAppOnoffappbylicenseAPIResponse 将 YunosTvpubadminContentAppOnoffappbylicenseAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentAppOnoffappbylicenseAPIResponse(v *YunosTvpubadminContentAppOnoffappbylicenseAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentAppOnoffappbylicenseAPIResponse.Put(v)
}
