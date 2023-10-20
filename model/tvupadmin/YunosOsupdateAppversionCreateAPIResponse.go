package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionCreateAPIResponse 创建应用升级任务 API返回值
// yunos.osupdate.appversion.create
//
// 创建应用升级任务
type YunosOsupdateAppversionCreateAPIResponse struct {
	model.CommonResponse
	YunosOsupdateAppversionCreateAPIResponseModel
}

// Reset 清空结构体
func (m *YunosOsupdateAppversionCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosOsupdateAppversionCreateAPIResponseModel).Reset()
}

// YunosOsupdateAppversionCreateAPIResponseModel is 创建应用升级任务 成功返回结果
type YunosOsupdateAppversionCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_appversion_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *YunosOsupdateAppversionCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolYunosOsupdateAppversionCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosOsupdateAppversionCreateAPIResponse)
	},
}

// GetYunosOsupdateAppversionCreateAPIResponse 从 sync.Pool 获取 YunosOsupdateAppversionCreateAPIResponse
func GetYunosOsupdateAppversionCreateAPIResponse() *YunosOsupdateAppversionCreateAPIResponse {
	return poolYunosOsupdateAppversionCreateAPIResponse.Get().(*YunosOsupdateAppversionCreateAPIResponse)
}

// ReleaseYunosOsupdateAppversionCreateAPIResponse 将 YunosOsupdateAppversionCreateAPIResponse 保存到 sync.Pool
func ReleaseYunosOsupdateAppversionCreateAPIResponse(v *YunosOsupdateAppversionCreateAPIResponse) {
	v.Reset()
	poolYunosOsupdateAppversionCreateAPIResponse.Put(v)
}
