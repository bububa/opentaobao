package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDiccontroltaskAddAPIResponse 新增停开服任务 API返回值
// yunos.tvpubadmin.diccontroltask.add
//
// 新增停开服任务
type YunosTvpubadminDiccontroltaskAddAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDiccontroltaskAddAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDiccontroltaskAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDiccontroltaskAddAPIResponseModel).Reset()
}

// YunosTvpubadminDiccontroltaskAddAPIResponseModel is 新增停开服任务 成功返回结果
type YunosTvpubadminDiccontroltaskAddAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_diccontroltask_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDiccontroltaskAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminDiccontroltaskAddAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDiccontroltaskAddAPIResponse)
	},
}

// GetYunosTvpubadminDiccontroltaskAddAPIResponse 从 sync.Pool 获取 YunosTvpubadminDiccontroltaskAddAPIResponse
func GetYunosTvpubadminDiccontroltaskAddAPIResponse() *YunosTvpubadminDiccontroltaskAddAPIResponse {
	return poolYunosTvpubadminDiccontroltaskAddAPIResponse.Get().(*YunosTvpubadminDiccontroltaskAddAPIResponse)
}

// ReleaseYunosTvpubadminDiccontroltaskAddAPIResponse 将 YunosTvpubadminDiccontroltaskAddAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDiccontroltaskAddAPIResponse(v *YunosTvpubadminDiccontroltaskAddAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDiccontroltaskAddAPIResponse.Put(v)
}
