package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户openId APIResponse
taobao.koubei.tribe.open.user.query

口碑综合体通过手机号码获取加密后的用户openId
*/
type TaobaoKoubeiTribeOpenUserQueryAPIResponse struct {
    model.CommonResponse
    TaobaoKoubeiTribeOpenUserQueryResponse
}

type TaobaoKoubeiTribeOpenUserQueryResponse struct {
    XMLName xml.Name `xml:"koubei_tribe_open_user_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoKoubeiTribeOpenUserQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
