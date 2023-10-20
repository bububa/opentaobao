package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateVersionstatusUpdateAPIResponse 更新应用升级状态 API返回值
// yunos.osupdate.versionstatus.update
//
// 更新应用升级状态
type YunosOsupdateVersionstatusUpdateAPIResponse struct {
	model.CommonResponse
	YunosOsupdateVersionstatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *YunosOsupdateVersionstatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosOsupdateVersionstatusUpdateAPIResponseModel).Reset()
}

// YunosOsupdateVersionstatusUpdateAPIResponseModel is 更新应用升级状态 成功返回结果
type YunosOsupdateVersionstatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_versionstatus_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *YunosOsupdateVersionstatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolYunosOsupdateVersionstatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosOsupdateVersionstatusUpdateAPIResponse)
	},
}

// GetYunosOsupdateVersionstatusUpdateAPIResponse 从 sync.Pool 获取 YunosOsupdateVersionstatusUpdateAPIResponse
func GetYunosOsupdateVersionstatusUpdateAPIResponse() *YunosOsupdateVersionstatusUpdateAPIResponse {
	return poolYunosOsupdateVersionstatusUpdateAPIResponse.Get().(*YunosOsupdateVersionstatusUpdateAPIResponse)
}

// ReleaseYunosOsupdateVersionstatusUpdateAPIResponse 将 YunosOsupdateVersionstatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseYunosOsupdateVersionstatusUpdateAPIResponse(v *YunosOsupdateVersionstatusUpdateAPIResponse) {
	v.Reset()
	poolYunosOsupdateVersionstatusUpdateAPIResponse.Put(v)
}
