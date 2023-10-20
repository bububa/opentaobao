package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceUpdateosstatusAPIResponse 更新系统版本审核状态 API返回值
// yunos.tvpubadmin.device.updateosstatus
//
// 更新系统版本审核状态
type YunosTvpubadminDeviceUpdateosstatusAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceUpdateosstatusAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceUpdateosstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceUpdateosstatusAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceUpdateosstatusAPIResponseModel is 更新系统版本审核状态 成功返回结果
type YunosTvpubadminDeviceUpdateosstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_updateosstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceUpdateosstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminDeviceUpdateosstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceUpdateosstatusAPIResponse)
	},
}

// GetYunosTvpubadminDeviceUpdateosstatusAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceUpdateosstatusAPIResponse
func GetYunosTvpubadminDeviceUpdateosstatusAPIResponse() *YunosTvpubadminDeviceUpdateosstatusAPIResponse {
	return poolYunosTvpubadminDeviceUpdateosstatusAPIResponse.Get().(*YunosTvpubadminDeviceUpdateosstatusAPIResponse)
}

// ReleaseYunosTvpubadminDeviceUpdateosstatusAPIResponse 将 YunosTvpubadminDeviceUpdateosstatusAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceUpdateosstatusAPIResponse(v *YunosTvpubadminDeviceUpdateosstatusAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceUpdateosstatusAPIResponse.Put(v)
}
