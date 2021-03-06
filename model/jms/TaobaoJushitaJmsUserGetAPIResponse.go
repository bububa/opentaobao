package jms

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsUserGetAPIResponse 查询某个用户是否同步消息 API返回值
// taobao.jushita.jms.user.get
//
// 查询某个用户是否同步消息，只支持单个查询
type TaobaoJushitaJmsUserGetAPIResponse struct {
	model.CommonResponse
	TaobaoJushitaJmsUserGetAPIResponseModel
}

// TaobaoJushitaJmsUserGetAPIResponseModel is 查询某个用户是否同步消息 成功返回结果
type TaobaoJushitaJmsUserGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jushita_jms_user_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步的用户信息
	OnsUser *TmcUser `json:"ons_user,omitempty" xml:"ons_user,omitempty"`
}
