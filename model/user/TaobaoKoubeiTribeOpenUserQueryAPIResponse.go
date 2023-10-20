package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiTribeOpenUserQueryAPIResponse 获取用户openId API返回值
// taobao.koubei.tribe.open.user.query
//
// 口碑综合体通过手机号码获取加密后的用户openId
type TaobaoKoubeiTribeOpenUserQueryAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiTribeOpenUserQueryAPIResponseModel
}

// TaobaoKoubeiTribeOpenUserQueryAPIResponseModel is 获取用户openId 成功返回结果
type TaobaoKoubeiTribeOpenUserQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_tribe_open_user_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoKoubeiTribeOpenUserQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
