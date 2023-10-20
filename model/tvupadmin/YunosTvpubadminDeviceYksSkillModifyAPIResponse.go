package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksSkillModifyAPIResponse 修改技能 API返回值
// yunos.tvpubadmin.device.yks.skill.modify
//
// 修改技能
type YunosTvpubadminDeviceYksSkillModifyAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceYksSkillModifyAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceYksSkillModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceYksSkillModifyAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceYksSkillModifyAPIResponseModel is 修改技能 成功返回结果
type YunosTvpubadminDeviceYksSkillModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceYksSkillModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminDeviceYksSkillModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceYksSkillModifyAPIResponse)
	},
}

// GetYunosTvpubadminDeviceYksSkillModifyAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceYksSkillModifyAPIResponse
func GetYunosTvpubadminDeviceYksSkillModifyAPIResponse() *YunosTvpubadminDeviceYksSkillModifyAPIResponse {
	return poolYunosTvpubadminDeviceYksSkillModifyAPIResponse.Get().(*YunosTvpubadminDeviceYksSkillModifyAPIResponse)
}

// ReleaseYunosTvpubadminDeviceYksSkillModifyAPIResponse 将 YunosTvpubadminDeviceYksSkillModifyAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceYksSkillModifyAPIResponse(v *YunosTvpubadminDeviceYksSkillModifyAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceYksSkillModifyAPIResponse.Put(v)
}
