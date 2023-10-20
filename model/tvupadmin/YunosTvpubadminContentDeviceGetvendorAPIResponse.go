package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentDeviceGetvendorAPIResponse 查询设备Vendor信息 API返回值
// yunos.tvpubadmin.content.device.getvendor
//
// 查询设备Vendor信息
type YunosTvpubadminContentDeviceGetvendorAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentDeviceGetvendorAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentDeviceGetvendorAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentDeviceGetvendorAPIResponseModel).Reset()
}

// YunosTvpubadminContentDeviceGetvendorAPIResponseModel is 查询设备Vendor信息 成功返回结果
type YunosTvpubadminContentDeviceGetvendorAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_device_getvendor_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// list
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentDeviceGetvendorAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentDeviceGetvendorAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentDeviceGetvendorAPIResponse)
	},
}

// GetYunosTvpubadminContentDeviceGetvendorAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentDeviceGetvendorAPIResponse
func GetYunosTvpubadminContentDeviceGetvendorAPIResponse() *YunosTvpubadminContentDeviceGetvendorAPIResponse {
	return poolYunosTvpubadminContentDeviceGetvendorAPIResponse.Get().(*YunosTvpubadminContentDeviceGetvendorAPIResponse)
}

// ReleaseYunosTvpubadminContentDeviceGetvendorAPIResponse 将 YunosTvpubadminContentDeviceGetvendorAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentDeviceGetvendorAPIResponse(v *YunosTvpubadminContentDeviceGetvendorAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentDeviceGetvendorAPIResponse.Put(v)
}
