package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceyksskillonlineAPIResponse 迎客松技能上架接口 API返回值
// yunos.tvpubadmin.device.yks.skill.online
//
// 迎客松技能上架接口
type YunostvpubadmindeviceyksskillonlineAPIResponse struct {
	model.CommonResponse
	YunostvpubadmindeviceyksskillonlineAPIResponseModel
}

// YunostvpubadmindeviceyksskillonlineAPIResponseModel is 迎客松技能上架接口 成功返回结果
type YunostvpubadmindeviceyksskillonlineAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_device_yks_skill_online_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
