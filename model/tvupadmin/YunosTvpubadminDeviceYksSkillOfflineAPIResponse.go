package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceYksSkillOfflineAPIResponse
技能下架 API返回值
yunos.tvpubadmin.device.yks.skill.offline

迎客松平台技能下架 */
type YunosTvpubadminDeviceYksSkillOfflineAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDeviceYksSkillOfflineAPIResponseModel
}

// YunosTvpubadminDeviceYksSkillOfflineAPIResponseModel is 技能下架 成功返回结果
type YunosTvpubadminDeviceYksSkillOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
