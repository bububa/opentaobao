package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksSkillOnlineAPIResponse 迎客松技能上架接口 API返回值
// yunos.tvpubadmin.device.yks.skill.online
//
// 迎客松技能上架接口
type YunosTvpubadminDeviceYksSkillOnlineAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceYksSkillOnlineAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceYksSkillOnlineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDeviceYksSkillOnlineAPIResponseModel).Reset()
}

// YunosTvpubadminDeviceYksSkillOnlineAPIResponseModel is 迎客松技能上架接口 成功返回结果
type YunosTvpubadminDeviceYksSkillOnlineAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_online_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDeviceYksSkillOnlineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminDeviceYksSkillOnlineAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDeviceYksSkillOnlineAPIResponse)
	},
}

// GetYunosTvpubadminDeviceYksSkillOnlineAPIResponse 从 sync.Pool 获取 YunosTvpubadminDeviceYksSkillOnlineAPIResponse
func GetYunosTvpubadminDeviceYksSkillOnlineAPIResponse() *YunosTvpubadminDeviceYksSkillOnlineAPIResponse {
	return poolYunosTvpubadminDeviceYksSkillOnlineAPIResponse.Get().(*YunosTvpubadminDeviceYksSkillOnlineAPIResponse)
}

// ReleaseYunosTvpubadminDeviceYksSkillOnlineAPIResponse 将 YunosTvpubadminDeviceYksSkillOnlineAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDeviceYksSkillOnlineAPIResponse(v *YunosTvpubadminDeviceYksSkillOnlineAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDeviceYksSkillOnlineAPIResponse.Put(v)
}
