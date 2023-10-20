package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceyksskilldeleteAPIResponse 技能删除 API返回值
// yunos.tvpubadmin.device.yks.skill.delete
//
// 删除技能
type YunostvpubadmindeviceyksskilldeleteAPIResponse struct {
	model.CommonResponse
	YunostvpubadmindeviceyksskilldeleteAPIResponseModel
}

// YunostvpubadmindeviceyksskilldeleteAPIResponseModel is 技能删除 成功返回结果
type YunostvpubadmindeviceyksskilldeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
