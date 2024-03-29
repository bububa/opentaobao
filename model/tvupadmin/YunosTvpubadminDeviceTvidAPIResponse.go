package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceTvidAPIResponse 查询终端信息 API返回值
// yunos.tvpubadmin.device.tvid
//
// 通过UUID查询终端信息
type YunosTvpubadminDeviceTvidAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceTvidAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceTvidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceTvidAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceTvidAPIResponseModel is 查询终端信息 成功返回结果
type YunosTvpubadminDeviceTvidAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_tvid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *DeviceInfoDo `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceTvidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = nil
}

var poolYunosTvpubadminDeviceTvidAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceTvidAPIResponse)
	},
}

// GetYunosTvpubadminDeviceTvidAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceTvidAPIResponse
func GetYunosTvpubadminDeviceTvidAPIResponse() *YunosTvpubadminDeviceTvidAPIResponse {
	return poolYunosTvpubadminDeviceTvidAPIResponse.Get().(*YunosTvpubadminDeviceTvidAPIResponse)
}

// ReleaseYunosTvpubadminDeviceTvidAPIResponse 将 YunosTvpubadminDeviceTvidAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceTvidAPIResponse(v *YunosTvpubadminDeviceTvidAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceTvidAPIResponse.Put(v)
}
