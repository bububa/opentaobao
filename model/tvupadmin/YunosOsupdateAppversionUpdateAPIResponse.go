package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionUpdateAPIResponse 应用升级任务更新 API返回值
// yunos.osupdate.appversion.update
//
// 应用升级任务更新
type YunosOsupdateAppversionUpdateAPIResponse struct {
	model.CommonResponse
	YunosOsupdateAppversionUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *YunosOsupdateAppversionUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosOsupdateAppversionUpdateAPIResponseModel).Reset()
}

// YunosOsupdateAppversionUpdateAPIResponseModel is 应用升级任务更新 成功返回结果
type YunosOsupdateAppversionUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_appversion_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *YunosOsupdateAppversionUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolYunosOsupdateAppversionUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosOsupdateAppversionUpdateAPIResponse)
	},
}

// GetYunosOsupdateAppversionUpdateAPIResponse 从 sync.Pool 获取 YunosOsupdateAppversionUpdateAPIResponse
func GetYunosOsupdateAppversionUpdateAPIResponse() *YunosOsupdateAppversionUpdateAPIResponse {
	return poolYunosOsupdateAppversionUpdateAPIResponse.Get().(*YunosOsupdateAppversionUpdateAPIResponse)
}

// ReleaseYunosOsupdateAppversionUpdateAPIResponse 将 YunosOsupdateAppversionUpdateAPIResponse 保存到 sync.Pool
func ReleaseYunosOsupdateAppversionUpdateAPIResponse(v *YunosOsupdateAppversionUpdateAPIResponse) {
	v.Reset()
	poolYunosOsupdateAppversionUpdateAPIResponse.Put(v)
}
