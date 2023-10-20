package subuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubuserdutysgetAPIResponse 获取指定账户的所有职务信息列表 API返回值
// taobao.subuser.dutys.get
//
// 通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息）
type TaobaosubuserdutysgetAPIResponse struct {
	model.CommonResponse
	TaobaosubuserdutysgetAPIResponseModel
}

// TaobaosubuserdutysgetAPIResponseModel is 获取指定账户的所有职务信息列表 成功返回结果
type TaobaosubuserdutysgetAPIResponseModel struct {
	XMLName xml.Name `xml:"subuser_dutys_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 职务信息
	Dutys []Duty `json:"dutys,omitempty" xml:"dutys>duty,omitempty"`
}
