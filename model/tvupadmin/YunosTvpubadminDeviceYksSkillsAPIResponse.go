package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksSkillsAPIResponse 根据设备id获取技能列表 API返回值
// yunos.tvpubadmin.device.yks.skills
//
// 根据设备id获取技能列表
type YunosTvpubadminDeviceYksSkillsAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceYksSkillsAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceYksSkillsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceYksSkillsAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceYksSkillsAPIResponseModel is 根据设备id获取技能列表 成功返回结果
type YunosTvpubadminDeviceYksSkillsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skills_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceYksSkillsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminDeviceYksSkillsAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceYksSkillsAPIResponse)
	},
}

// GetYunosTvpubadminDeviceYksSkillsAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceYksSkillsAPIResponse
func GetYunosTvpubadminDeviceYksSkillsAPIResponse() *YunosTvpubadminDeviceYksSkillsAPIResponse {
	return poolYunosTvpubadminDeviceYksSkillsAPIResponse.Get().(*YunosTvpubadminDeviceYksSkillsAPIResponse)
}

// ReleaseYunosTvpubadminDeviceYksSkillsAPIResponse 将 YunosTvpubadminDeviceYksSkillsAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceYksSkillsAPIResponse(v *YunosTvpubadminDeviceYksSkillsAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceYksSkillsAPIResponse.Put(v)
}
