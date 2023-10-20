package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceQueryAPIResponse 获取设备列表 API返回值
// yunos.tvpubadmin.device.query
//
// 获取设备列表
type YunosTvpubadminDeviceQueryAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceQueryAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceQueryAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceQueryAPIResponseModel is 获取设备列表 成功返回结果
type YunosTvpubadminDeviceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *PaginationDo `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = nil
}

var poolYunosTvpubadminDeviceQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceQueryAPIResponse)
	},
}

// GetYunosTvpubadminDeviceQueryAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceQueryAPIResponse
func GetYunosTvpubadminDeviceQueryAPIResponse() *YunosTvpubadminDeviceQueryAPIResponse {
	return poolYunosTvpubadminDeviceQueryAPIResponse.Get().(*YunosTvpubadminDeviceQueryAPIResponse)
}

// ReleaseYunosTvpubadminDeviceQueryAPIResponse 将 YunosTvpubadminDeviceQueryAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceQueryAPIResponse(v *YunosTvpubadminDeviceQueryAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceQueryAPIResponse.Put(v)
}
