package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDiccontroltaskUpdateAPIResponse 停开服任务状态变更 API返回值
// yunos.tvpubadmin.diccontroltask.update
//
// 停开服任务状态变更
type YunosTvpubadminDiccontroltaskUpdateAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDiccontroltaskUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDiccontroltaskUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDiccontroltaskUpdateAPIResponseModel).Reset()
}

// YunosTvpubadminDiccontroltaskUpdateAPIResponseModel is 停开服任务状态变更 成功返回结果
type YunosTvpubadminDiccontroltaskUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_diccontroltask_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDiccontroltaskUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminDiccontroltaskUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDiccontroltaskUpdateAPIResponse)
	},
}

// GetYunosTvpubadminDiccontroltaskUpdateAPIResponse 从 sync.Pool 获取 YunosTvpubadminDiccontroltaskUpdateAPIResponse
func GetYunosTvpubadminDiccontroltaskUpdateAPIResponse() *YunosTvpubadminDiccontroltaskUpdateAPIResponse {
	return poolYunosTvpubadminDiccontroltaskUpdateAPIResponse.Get().(*YunosTvpubadminDiccontroltaskUpdateAPIResponse)
}

// ReleaseYunosTvpubadminDiccontroltaskUpdateAPIResponse 将 YunosTvpubadminDiccontroltaskUpdateAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDiccontroltaskUpdateAPIResponse(v *YunosTvpubadminDiccontroltaskUpdateAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDiccontroltaskUpdateAPIResponse.Put(v)
}
