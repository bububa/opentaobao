package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceModelsAPIResponse 获取品牌下设备列表 API返回值
// yunos.tvpubadmin.device.models
//
// 获取品牌下设备列表
type YunosTvpubadminDeviceModelsAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceModelsAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceModelsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceModelsAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceModelsAPIResponseModel is 获取品牌下设备列表 成功返回结果
type YunosTvpubadminDeviceModelsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_models_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	ModelList []string `json:"model_list,omitempty" xml:"model_list>string,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceModelsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ModelList = m.ModelList[:0]
}

var poolYunosTvpubadminDeviceModelsAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceModelsAPIResponse)
	},
}

// GetYunosTvpubadminDeviceModelsAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceModelsAPIResponse
func GetYunosTvpubadminDeviceModelsAPIResponse() *YunosTvpubadminDeviceModelsAPIResponse {
	return poolYunosTvpubadminDeviceModelsAPIResponse.Get().(*YunosTvpubadminDeviceModelsAPIResponse)
}

// ReleaseYunosTvpubadminDeviceModelsAPIResponse 将 YunosTvpubadminDeviceModelsAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceModelsAPIResponse(v *YunosTvpubadminDeviceModelsAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceModelsAPIResponse.Put(v)
}
