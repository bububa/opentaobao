package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowEditAPIResponse 媒资节目信息修改 API返回值
// yunos.tvpubadmin.content.show.edit
//
// 供迎客松修改媒资节目信息
type YunosTvpubadminContentShowEditAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentShowEditAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentShowEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentShowEditAPIResponseModel).Reset()
}

// YunosTvpubadminContentShowEditAPIResponseModel is 媒资节目信息修改 成功返回结果
type YunosTvpubadminContentShowEditAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentShowEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminContentShowEditAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentShowEditAPIResponse)
	},
}

// GetYunosTvpubadminContentShowEditAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentShowEditAPIResponse
func GetYunosTvpubadminContentShowEditAPIResponse() *YunosTvpubadminContentShowEditAPIResponse {
	return poolYunosTvpubadminContentShowEditAPIResponse.Get().(*YunosTvpubadminContentShowEditAPIResponse)
}

// ReleaseYunosTvpubadminContentShowEditAPIResponse 将 YunosTvpubadminContentShowEditAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentShowEditAPIResponse(v *YunosTvpubadminContentShowEditAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentShowEditAPIResponse.Put(v)
}
