package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksSkillAddAPIResponse 添加技能 API返回值
// yunos.tvpubadmin.device.yks.skill.add
//
// 添加技能
type YunosTvpubadminDeviceYksSkillAddAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceYksSkillAddAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceYksSkillAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceYksSkillAddAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceYksSkillAddAPIResponseModel is 添加技能 成功返回结果
type YunosTvpubadminDeviceYksSkillAddAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceYksSkillAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminDeviceYksSkillAddAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceYksSkillAddAPIResponse)
	},
}

// GetYunosTvpubadminDeviceYksSkillAddAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceYksSkillAddAPIResponse
func GetYunosTvpubadminDeviceYksSkillAddAPIResponse() *YunosTvpubadminDeviceYksSkillAddAPIResponse {
	return poolYunosTvpubadminDeviceYksSkillAddAPIResponse.Get().(*YunosTvpubadminDeviceYksSkillAddAPIResponse)
}

// ReleaseYunosTvpubadminDeviceYksSkillAddAPIResponse 将 YunosTvpubadminDeviceYksSkillAddAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceYksSkillAddAPIResponse(v *YunosTvpubadminDeviceYksSkillAddAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceYksSkillAddAPIResponse.Put(v)
}
